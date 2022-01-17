//go:build debugNflex
// +build debugNflex

package nflex

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
)

var counter int32 = 1000

func debug(args ...interface{}) {
	log.Println(args...)
}
func debugID() int {
	return int(atomic.AddInt32(&counter, 1))
}

func id(raw Source) string {
	switch s := raw.(type) {
	case offset:
		return fmt.Sprintf("O%d/%s", s.debugID, id(s.source))
	case prefixSource:
		return fmt.Sprintf("P%d/%s", s.debugID, id(s.source))
	case parsedYAML:
		return fmt.Sprintf("J%d/%v", s.debugID, s.pathToHere)
	case parsedJSON:
		return fmt.Sprintf("J%d/%v", s.debugID, s.pathToHere)
	case *MultiSource:
		ss := make([]string, len(s.sources))
		for i, source := range s.sources {
			ss[i] = id(source)
		}
		return fmt.Sprintf("M%d/%v<%s>", s.debugID, s.pathToHere, strings.Join(ss, "|"))
	default:
		return "?"
	}
}

var debugCombine = combine
