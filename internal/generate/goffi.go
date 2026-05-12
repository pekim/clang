package generate

import (
	"slices"

	"github.com/dave/jennifer/jen"
)

func (ff functions) generateGoffiFunctionVars(file file) {
	for _, fn := range ff {
		if slices.Contains(manualFunctions, fn.cName) {
			continue
		}

		if _, supported := fn.isSupported(); supported {
			file.
				Var().Id("cif_"+fn.cName).
				Op("=").
				Op("&").Qual(typesPackage, "CallInterface").Block()
		}
	}

	file.Line()

	for _, fn := range ff {
		if slices.Contains(manualFunctions, fn.cName) {
			continue
		}

		if _, supported := fn.isSupported(); supported {
			file.Var().Id("ptr_"+fn.cName).Qual("unsafe", "Pointer")
		}
	}

	file.Line()
}

func (ff functions) generateGoffiInitFunction(file file) {
	file.Comment(`
Init initialises the library, and must be called before any other clang function is called.
	`)

	file.Func().
		Id("Init").
		Params(jen.Id("userPaths").Op("*").Id("LibraryPaths")).
		Error(). // return type
		BlockFunc(func(g *jen.Group) {
			g.Var().Id("err").Error()
			g.Line()

			g.
				List(jen.Id("library"), jen.Id("err")).
				Op(":=").
				Id("loadLibrary").Call(jen.Id("userPaths"))
			g.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Return().Id("err"),
			)
			g.Line()

			g.Id("err").Op("=").Id("initManual").Call(jen.Id("library"))
			g.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Return().Id("err"),
			)
			g.Line()

			for _, fn := range ff {
				if slices.Contains(manualFunctions, fn.cName) {
					continue
				}
				if _, supported := fn.isSupported(); !supported {
					continue
				}

				g.BlockFunc(func(g *jen.Group) {

					// get function pointer
					g.
						List(
							jen.Id("ptr_"+fn.cName),
							jen.Id("err"),
						).
						Op("=").
						Qual(ffiPackage, "GetSymbol").Call(
						jen.Id("library"),
						jen.Lit(fn.cName),
					)

					g.
						// Only prepare call interface if the function is available.
						If().Id("err").Op("==").Nil().BlockFunc(func(g *jen.Group) {

						// return type
						typeDescriptor, _ := fn.returnValue.goffiTypeDescriptor()
						g.
							Id("returnType").
							Op(":=").
							Add(typeDescriptor)

						// arg types
						g.
							Id("argTypes").
							Op(":=").
							Index().Op("*").Qual(typesPackage, "TypeDescriptor").
							ValuesFunc(func(g *jen.Group) {
								for _, param := range fn.params {
									typeDescriptor, _ := param.goffiTypeDescriptor()
									g.Line().Add(typeDescriptor)
								}
								g.Line()
							})

						// prepare call interface
						g.
							Id("err").
							Op("=").
							Qual(ffiPackage, "PrepareCallInterface").
							Call(
								jen.Id("cif_"+fn.cName),
								jen.Qual(typesPackage, "DefaultCall"),
								jen.Id("returnType"),
								jen.Id("argTypes"),
							)
						g.
							If().Id("err").Op("!=").Nil().BlockFunc(func(g *jen.Group) {
							g.Return().Qual("fmt", "Errorf").Call(
								jen.Lit("failed to prepare call interface for %s : %w"),
								jen.Lit(fn.cName),
								jen.Id("err"),
							)
						})
					})
				})

				g.Line()
			}

			g.Return().Nil()
		})
}
