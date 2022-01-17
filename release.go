//go:build !debugNflex

package nflex

var debugging = false

func debugID() int                                 { return 0 }
func debugf(fmt string, args ...interface{})       {}
func debug(args ...interface{})                    {}
func callers(levels int) []string                  { return nil }
func id(_ Source) string                           { return "" }
func debugCombine(_ []string, _ []string) []string { return nil }
