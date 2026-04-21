module github.com/pekim/clang

go 1.26.2

// Address deprecation of clang_getDiagnosticCategoryName.
// Avoids a noisy message when generating api.
replace github.com/go-clang/clang-v15 => github.com/pekim/clang-v15 v0.0.0-20240830114552-c0d27ccce9ec

require github.com/go-clang/clang-v15 v0.0.0-00010101000000-000000000000

require (
	github.com/dave/jennifer v1.7.1
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
