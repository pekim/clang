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
	if rv.isStruct() && rv.struct_.goName == "String" {
		return jen.String()
	}
	return rv.typ.goDecl()
}

func (rv returnValue) cVarToGoVar(cVar string, goVar string) jen.Code {
	if rv.isString {
		// 	ret := libc.GoString(retC)
		return jen.Id(goVar).Op(":=").Qual("github.com/pekim/clang/internal/libc", "GoString").Call(
			jen.Id(cVar),
		)
	}

	if rv.isStructPointer() {
		return jen.Id(goVar).Op(":=").Parens(jen.Op("*").Id(rv.structPointer.goName)).Parens(jen.Id(cVar))
	}

	if rv.isStruct() && rv.struct_.goName == "String" {
		return jen.Id(goVar).Op(":=").Id(cVar).Dot("CString").Call()
	}

	return jen.Id(goVar).Op(":=").Id(cVar)
}
