package generate

import "github.com/dave/jennifer/jen"

func (ss structs) generateTypeDescriptors() {
	file := newFile("clang", ".", "structtypedescriptor")
	defer file.save()

	for _, struct_ := range ss {
		struct_.generateTypeDescriptor(file)
	}

}

func (struct_ struct_) generateTypeDescriptor(file file) {
	file.
		Var().Id(struct_.typeDescriptorName).
		Op("=").
		Op("&").Qual(typesPackage, "TypeDescriptor").Values(jen.DictFunc(func(d jen.Dict) {
		d[jen.Id("Size")] = jen.Qual("unsafe", "Sizeof").Call(jen.Id(struct_.goName).Values())
		d[jen.Id("Alignment")] = jen.Lit(struct_.alignment)
		d[jen.Id("Kind")] = jen.Qual(typesPackage, "StructType")
		d[jen.Id("Members")] = jen.Index().Op("*").Qual(typesPackage, "TypeDescriptor").ValuesFunc(func(g *jen.Group) {
			for _, field := range struct_.fields {
				typeDescriptor, ok := field.goffiTypeDescriptor()
				if ok {
					g.Line().Add(typeDescriptor)
				} else {
					g.Line().Values(jen.DictFunc(func(d jen.Dict) {
						d[jen.Id("Size")] = jen.Lit(field.size)
						d[jen.Id("Alignment")] = jen.Lit(field.alignment)
						d[jen.Id("Kind")] = jen.Qual(typesPackage, "UInt8Type")
					}))
				}
			}
			g.Line()
		})
	}))

	file.Line()
}
