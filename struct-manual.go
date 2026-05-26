package clang

import "unsafe"

// GetStrings returns the StringSet's strings
func (ss StringSet) GetStrings() []string {
	strings := make([]string, ss.Count)

	for i, s := range unsafe.Slice(ss.Strings, ss.Count) {
		strings[i] = s.CString()
	}
	return strings
}
