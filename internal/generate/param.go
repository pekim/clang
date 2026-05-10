package generate

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/pekim/clang"
)

type param struct {
	typ
	cName     string
	goName    string
	cVar      string
	comment   string
	clangType clang.Type
	isOut     bool

	arrayParam    *param // if non-nil, then it's the array param, and this param is the array length
	arrayLenParam *param // if non-nil, then it's the array length param, and this param is the array param
}

func newParam(cursor clang.Cursor, name string) param {
	return param{
		cName:     name,
		goName:    goName(name),
		cVar:      "c_" + goName(name),
		comment:   commentText(cursor.ParsedComment()),
		clangType: cursor.CursorType(),
	}
}

func (param *param) enrich(gen *gen, index int) {
	param.typ = newTyp(param.clangType, gen)
	param.isOut = index > 0 && !param.isArray() &&
		(param.isScalarPointer || param.isPointerTypePointer() || param.isStructPointer())
}

func (param param) isSupported() (string, bool) {
	if param.isEnum() {
		return "", true
	}
	if param.isEnumPointer() {
		return "", true
	}
	if param.isArray() {
		if param.isStringPointer || param.isStructPointer() {
			return "", true
		}
	}
	if param.isPointerType() {
		return "", true
	}
	if param.isPointerTypePointer() {
		return "", true
	}
	if param.isScalar() {
		return "", true
	}
	if param.isString {
		return "", true
	}
	if param.isStruct() {
		return "", true
	}
	if param.isStructPointer() {
		return "", true
	}
	if param.isPointer && param.isVoid {
		return "", true
	}

	return fmt.Sprintf("param %s : %s", param.cName, param.typ), false
}

func (param param) isArray() bool {
	return param.arrayLenParam != nil
}

func (param param) isArrayLen() bool {
	return param.arrayParam != nil
}

// var out_TU TranslationUnit
// c_out_TU:=&out_TU

func (param param) goDecl() jen.Code {
	if param.isArrayLen() {
		return jen.Null()
	}
	if param.isOut {
		return jen.Null()
	}

	if param.isArray() && param.isStringPointer {
		return jen.Id(param.goName).Index().String()
	}
	if param.isArray() && param.isStructPointer() {
		return jen.Id(param.goName).Index().Id(param.structPointer.goName)
	}

	return jen.Id(param.goName).Add(param.typ.goDecl())
}

func (param param) goReturnDecl() jen.Code {
	if !param.isOut {
		return jen.Null()
	}

	return param.goOutReturnDecl()
}

func (param param) outCVar(g *jen.Group) {
	g.Var().Id(param.goName).Add(param.goOutVarDecl())
	g.Id(param.cVar).Op(":=").Op("&").Id(param.goName)
}

func (param param) goVarToCVar(g *jen.Group) {
	if param.isOut {
		param.outCVar(g)
		return
	}

	if param.isArrayLen() {
		g.Id(param.cVar).Op(":=").Len(jen.Id(param.arrayParam.goName))
		return
	}
	if param.isArray() {
		if param.isStringPointer {
			freeVar := "free_" + param.cVar
			// c_someParam, free_c_someParam := libc.CStrings(someParam)
			g.
				List(jen.Id(param.cVar), jen.Id(freeVar)).
				Op(":=").
				Qual("github.com/pekim/clang/internal/libc", "CStrings").Call(jen.Id(param.goName))
			// defer free_c_someParam()
			g.Defer().Id(freeVar).Call()
		}
		if param.isStructPointer() {
			// var c_someParam unsafe.Pointer
			// if len(someParam) > 0 {
			// 	c_someParam = unsafe.Pointer(&unsaved_files[0])
			// }
			g.Var().Id(param.cVar).Qual("unsafe", "Pointer")
			g.If(jen.Len(jen.Id(param.goName)).Op(">").Lit(0)).Block(
				jen.Id(param.cVar).Op("=").Qual("unsafe", "Pointer").Parens(
					jen.Op("&").Id(param.goName).Index(jen.Lit(0)),
				),
			)
		}
		return
	}
	if param.isEnum() {
		g.Id(param.cVar).Op(":=").Id(param.goName)
		return
	}
	if param.isEnumPointer() {
		g.Id(param.cVar).Op(":=").Id(param.goName)
		return
	}
	if param.isPointerType() {
		g.Id(param.cVar).Op(":=").Id(param.goName)
		return
	}
	if param.isPointerTypePointer() {
		g.Id(param.cVar).Op(":=").Id(param.goName)
		return
	}
	if param.isScalar() {
		g.Id(param.cVar).Op(":=").Id(param.goName)
		return
	}
	if param.isString {
		freeVar := "free_" + param.cVar
		// c_someParam, free_c_someParam := libc.CString(someParam)
		g.
			List(jen.Id(param.cVar), jen.Id(freeVar)).
			Op(":=").
			Qual("github.com/pekim/clang/internal/libc", "CString").Call(jen.Id(param.goName))
		// defer free_c_someParam()
		g.Defer().Id(freeVar).Call()
		return
	}
	if param.isStruct() {
		g.Id(param.cVar).Op(":=").Id(param.goName)
		return
	}
	if param.isStructPointer() {
		g.Id(param.cVar).Op(":=").Id(param.goName)
		return
	}
	if param.isPointer && param.isVoid {
		g.Id(param.cVar).Op(":=").Id(param.goName)
		return
	}
}

type params []param

func newParams(cursor clang.Cursor) params {
	var params params

	p := 0
	cursor.VisitChildren(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
		if cursor.Kind == clang.Cursor_ParmDecl {
			name := cursor.CursorSpelling()
			if name == "" {
				// The parameter has no name, so use a generated name.
				name = fmt.Sprintf("p%d", p)
			}
			params = append(params, newParam(cursor, name))
			p++
		}

		return clang.ChildVisit_Continue
	})

	return params
}

func (pp params) enrich(gen *gen) {
	// Pair up array params with their array length params
	for i, param := range pp {
		if strings.HasPrefix(param.cName, "num_") {
			arrayParamIndex := pp.find(strings.TrimPrefix(param.cName, "num_"))
			if arrayParamIndex != -1 {
				arrayParam := pp[arrayParamIndex]
				pp[i].arrayParam = &arrayParam
				pp[arrayParamIndex].arrayLenParam = &pp[i]
			}
		}
	}

	for i := range pp {
		(pp[i]).enrich(gen, i)
	}
}

func (pp params) isSupported() (string, bool) {
	for _, param := range pp {
		if reason, supported := param.isSupported(); !supported {
			return reason, false
		}
	}

	return "", true
}

func (pp params) someOut() bool {
	for _, param := range pp {
		if param.isOut {
			return true
		}
	}
	return false
}

func (pp params) find(cName string) int {
	return slices.IndexFunc(pp, func(param param) bool { return param.cName == cName })
}
