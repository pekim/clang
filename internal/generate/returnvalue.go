package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/pekim/clang"
)

type returnValue struct {
	typ
	clangType clang.Type
}

func newReturnValue(typ clang.Type) returnValue {
	return returnValue{
		clangType: typ,
	}
}

func (rv *returnValue) enrich(gen *gen) {
	rv.typ = newTyp(rv.clangType, gen)
}

func (rv *returnValue) isSupported() bool {
	if rv.isVoid {
		return true
	}
	if rv.isEnum() {
		return true
	}
	if rv.isScalar() {
		return true
	}
	if rv.isPointerType() {
		return true
	}
	if rv.isString {
		return true
	}
	if rv.isStruct() {
		return true
	}
	if rv.isStructPointer() {
		return true
	}

	return false
}

func (rv returnValue) goDecl() jen.Code {
	if rv.isCXString() {
		return jen.String()
	}
	return rv.typ.goDecl()
}

func (rv returnValue) cVarToGoVar(g *jen.Group, cVar string, goVar string) {
	if rv.isString {
		// 	ret := libc.GoString(retC)
		g.Id(goVar).Op(":=").Qual("github.com/pekim/clang/internal/libc", "GoString").Call(
			jen.Id(cVar),
		)

	} else if rv.isStructPointer() {
		g.Id(goVar).Op(":=").Parens(jen.Op("*").Id(rv.structPointer.goName)).Parens(jen.Id(cVar))

	} else if rv.isCXString() {
		g.Id(goVar).Op(":=").Id(cVar).Dot("CString").Call()
		g.Id(cVar).Dot("DisposeString").Call().Op(";")

	} else {
		g.Id(goVar).Op(":=").Id(cVar)
	}
}
