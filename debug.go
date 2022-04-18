//go:build debugNflex
// +build debugNflex

package nflex

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync/atomic"
)

var counter int32 = 1000

func debug(args ...interface{}) {
	c := make([]interface{}, len(args))
	for i, arg := range args {
		if f, ok := arg.(func() string); ok {
			c[i] = f()
		} else {
			c[i] = arg
		}
	}
	log.Println(c...)
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

func debugKeys(s Source) string {
	switch s.Type() {
	case Undefined:
		return "=UNDEFINED"
	case Nil:
		return "=NIL"
	case String:
		str, err := s.GetString()
		if err != nil {
			return "=s!" + err.Error() + "!"
		}
		return "='" + str + "'"
	case Float:
		f, err := s.GetFloat()
		if err != nil {
			return "=f!" + err.Error() + "!"
		}
		return strconv.FormatFloat(f, 'E', -1, 64)
	case Int:
		i, err := s.GetInt()
		if err != nil {
			return "=i!" + err.Error() + "!"
		}
		return strconv.FormatInt(i, 10)
	case Bool:
		b, err := s.GetBool()
		if err != nil {
			return "=b!" + err.Error() + "!"
		}
		return strconv.FormatBool(b)
	case Slice:
		l, err := s.Len()
		if err != nil {
			return "slice:!" + err.Error() + "!"
		}
		if l == 0 {
			return "slice:[]"
		}
		return "slice:[0.." + strconv.Itoa(l-1) + "]"
	case Map:
		k, err := s.Keys()
		if err != nil {
			return "keys:!" + err.Error() + "!"
		}
		if len(k) == 0 {
			return "keys:{}"
		}
		return `{"` + strings.Join(k, `", "`) + `"}`
	default:
		return "=???"
	}
}
