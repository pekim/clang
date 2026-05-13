package generate

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/pekim/clang"
)

// Functions that will be maintained manually, and should not be generated.
var manualFunctions = []string{
	"clang_visitChildren",
}

type function struct {
	cursor      clang.Cursor
	cName       string
	goName      string
	comment     string
	returnValue returnValue
	params      params
	isMethod    bool
}

type functions []function

func (gen *gen) addFunction(cursor clang.Cursor) {
	function := function{
		cursor:      cursor,
		cName:       cursor.Spelling(),
		goName:      exportedGoName(strings.TrimPrefix(cursor.Spelling(), "clang_")),
		comment:     commentText(cursor.ParsedComment()),
		returnValue: newReturnValue(cursor.ResultType()),
		params:      newParams(cursor),
	}

	gen.functions = append(gen.functions, function)
}

func (ff functions) enrich(gen *gen) {
	for i := range ff {
		(ff[i]).enrich(gen)
	}
}

func (function *function) enrich(gen *gen) {
	function.returnValue.enrich(gen)
	function.params.enrich(gen)

	if len(function.params) > 0 {
		// the receiver param
		param := function.params[0]

		if param.isEnum() ||
			param.isPointerType() ||
			param.isStruct() ||
			param.isStructPointer() {

			function.isMethod = true

			// Trim some get prefixes from name.
			typeName := ""
			if param.isEnum() {
				typeName = param.enum.goName
			}
			if param.isPointerType() {
				typeName = param.pointerType.goName
			}
			if param.isStruct() {
				typeName = param.struct_.goName
			}
			if param.isStringPointer {
				typeName = param.structPointer.goName
			}
			function.goName = strings.TrimPrefix(function.goName, typeName+"_get")
			function.goName = strings.TrimPrefix(function.goName, typeName)
			function.goName = strings.TrimPrefix(function.goName, "Get")
			function.goName = strings.TrimPrefix(function.goName, "_")
			function.goName = strings.Replace(function.goName, "_get", "", 1)
			if strings.HasPrefix(function.goName, typeName) && len(function.goName) > len(typeName) {
				if !slices.Contains([]string{
					"clang_getTranslationUnitCursor", // would result in a clash with clang_getCursor
					"clang_getCursorKind",            // would result in a clash with the Cursor.Kind field
				}, function.cName) {
					function.goName = strings.TrimPrefix(function.goName, typeName)
				}
			}
			function.goName = exportedGoName(function.goName)
		}
	}
}

func (function function) isSupported() (string, bool) {
	if slices.Contains(manualFunctions, function.cName) {
		return "", true
	}

	if !function.returnValue.isSupported() {
		return fmt.Sprintf("return value : %s", function.returnValue.typ), false
	}

	if reason, supported := function.params.isSupported(); !supported {
		return reason, false
	}

	return "", true
}

func (function function) generateImplementation(file file) {
	if slices.Contains(manualFunctions, function.cName) {
		return
	}

	defer file.Line()

	if reason, supported := function.isSupported(); !supported {
		file.Commentf("not supported : %s : %s", function.cName, reason)
		return
	}

	file.Comment(function.comment)
	file.Commentf("Wraps the C function %s.\n", function.cName)

	file.
		Func().

		// method receiver
		Do(func(s *jen.Statement) {
			if function.isMethod {
				s.Params(jen.Add(function.params[0].goDecl()))
			}
		}).
		Id(function.goName).

		// parameters
		ParamsFunc(func(g *jen.Group) {
			params := function.params
			if function.isMethod {
				params = params[1:]
			}
			for _, param := range params {
				g.Add(param.goDecl())
			}
		}).
		// return declaration
		ParamsFunc(func(g *jen.Group) {
			for _, param := range function.params {
				g.Add(param.goReturnDecl())
			}
			g.Add(function.returnValue.goDecl())
		}).

		// function body
		BlockFunc(func(g *jen.Group) {
			// C argument vars
			for _, param := range function.params {
				param.goVarToCVar(g)
			}

			if len(function.params) > 0 {
				g.Line()
			}

			// return var
			if !function.returnValue.isVoid {
				g.Var().Id("retC").Add(function.returnValue.cDecl())
			}

			// args var
			g.Id("args").Op(":=").Index().Qual("unsafe", "Pointer").ValuesFunc(func(g *jen.Group) {
				for i, param := range function.params {
					g.Line().Qual("unsafe", "Pointer").Parens(jen.Op("&").Do(func(s *jen.Statement) {
						if function.isMethod && i == 0 && param.isPointerType() {
							s.Id(param.cVar).Dot("ptr")
						} else {
							s.Id(param.cVar)
						}
					}))
				}
				g.Line()
			})
			g.Line()

			// call C function
			g.Id("err").Op(":=").Qual(ffiPackage, "CallFunction").
				// arguments
				CallFunc(func(g *jen.Group) {
					g.Line().Id("cif_" + function.cName)
					g.Line().Id("ptr_" + function.cName)

					if !function.returnValue.isVoid {
						g.Line().Qual("unsafe", "Pointer").Parens(jen.Op("&").Id("retC"))
					} else {
						g.Line().Nil()
					}

					g.Line().Id("args")
					g.Line()
				})

			g.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Panic(jen.Qual("fmt", "Sprintf").Call(
					jen.Lit("failed to call %s : %s"),
					jen.Lit(function.cName),
					jen.Id("err"),
				)),
			)

			if !function.returnValue.isVoid || function.params.someOut() {
				if !function.returnValue.isVoid {
					g.Line()
					function.returnValue.cVarToGoVar(g, "retC", "ret")
				}

				g.Return().ListFunc(func(g *jen.Group) {
					for _, param := range function.params {
						if param.isOut {
							g.Id(param.goName).Do(func(s *jen.Statement) {
								if param.isCXStringPointer() {
									s.Dot("CString").Call()
								}
							})
						}
					}
					if !function.returnValue.isVoid {
						g.Id("ret")
					}
				})
			}
		})
}

func (ff *functions) generate() {
	ff.generateGoffi()
	ff.generateImplementations()
}

func (ff *functions) generateGoffi() {
	file := newFile("clang", ".", "goffi")
	defer file.save()

	ff.generateGoffiFunctionVars(file)
	ff.generateGoffiInitFunction(file)
}

func (ff functions) generateImplementations() {
	file := newFile("clang", ".", "function")
	defer file.save()

	for _, function := range ff {
		function.generateImplementation(file)
	}
}

func (ff functions) stats() (int, int) {
	supportedCount := 0
	for _, fn := range ff {
		if _, supported := fn.isSupported(); supported {
			supportedCount++
		}
	}
	return supportedCount, len(ff)
}
