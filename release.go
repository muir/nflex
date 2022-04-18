//go:build !debugNflex
// +build !debugNflex

package nflex

func debugID() int                                 { return 0 }
func debug(args ...interface{})                    {}
func id(_ Source) string                           { return "" }
func debugCombine(_ []string, _ []string) []string { return nil }
func debugKeys(s Source) string                    { return "" }
