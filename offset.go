package nflex

import (
	"strconv"

	"github.com/pkg/errors"
)

type offset struct {
	offsets []int
	source  Source
	debugID int
}

// WithOffset returns a modified source that offsets slice lookups
// such that it requires a larger key to do the lookup.
func WithOffset(source Source, offsets ...int) Source {
	if source == nil {
		return nil
	}
	return offset{
		offsets: offsets,
		source:  source,
		debugID: debugID(),
	}
}

func (o offset) transform(keys []string) ([]string, error) {
	n := make([]string, len(keys))
	for i, key := range keys {
		if o.offsets[i] == 0 {
			n[i] = key
			continue
		}
		k, err := strconv.Atoi(key)
		if err != nil {
			return nil, errors.Wrap(err, "expecting integer key")
		}
		n[i] = strconv.Itoa(k - o.offsets[i])
	}
	return n, nil
}

func (o offset) Exists(keys ...string) bool {
	tk, err := o.transform(keys)
	if err != nil {
		return false
	}
	return o.source.Exists(tk...)
}

func (o offset) GetBool(keys ...string) (bool, error) {
	tk, err := o.transform(keys)
	if err != nil {
		return false, err
	}
	return o.source.GetBool(tk...)
}

func (o offset) GetInt(keys ...string) (int64, error) {
	tk, err := o.transform(keys)
	if err != nil {
		return 0, err
	}
	return o.source.GetInt(tk...)
}

func (o offset) GetFloat(keys ...string) (float64, error) {
	tk, err := o.transform(keys)
	if err != nil {
		return 0, err
	}
	return o.source.GetFloat(tk...)
}

func (o offset) GetString(keys ...string) (string, error) {
	tk, err := o.transform(keys)
	if err != nil {
		return "", err
	}
	return o.source.GetString(tk...)
}

func (o offset) Recurse(keys ...string) Source {
	tk, err := o.transform(keys)
	if err != nil {
		debug("nflex/offset Recurse(", keys, ")", o.debugID, "-> nil")
		return nil
	}
	r := o.source.Recurse(tk...)
	if len(o.offsets) <= len(keys) {
		debug("nflex/offset Recurse(", keys, ")", o.debugID, "-> no offset")
		return r
	}
	n := WithOffset(r, o.offsets[len(keys):]...)
	debug("nflex/offset Recurse(", keys, ")", o.debugID, "->", n.(offset).debugID)
	return n
}

func (o offset) Keys(keys ...string) ([]string, error) {
	tk, err := o.transform(keys)
	if err != nil {
		return nil, err
	}
	return o.source.Keys(tk...)
}

func (o offset) Len(keys ...string) (int, error) {
	tk, err := o.transform(keys)
	if err != nil {
		return 0, err
	}
	length, err := o.source.Len(tk...)
	if len(o.offsets) <= len(keys) {
		return length, err
	}
	return length, err
}

func (o offset) Type(keys ...string) NodeType {
	tk, err := o.transform(keys)
	if err != nil {
		return Undefined
	}
	return o.source.Type(tk...)
}
