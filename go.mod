module github.com/pekim/clang

go 1.26.2

// Address deprecation of clang_getDiagnosticCategoryName.
// Avoids a noisy message when generating api.
replace github.com/go-clang/clang-v15 => github.com/pekim/clang-v15 v0.0.0-20240830114552-c0d27ccce9ec

require github.com/go-clang/clang-v15 v0.0.0-00010101000000-000000000000
