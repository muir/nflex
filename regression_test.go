package nflex

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegression01(t *testing.T) {
	var s Source
	s1, err := UnmarshalFile("rdata1.yaml", WithFS(content))
	require.NoError(t, err, "open rdata1.yaml")
	assert.Equal(t, Slice, s1.Recurse("QQ").Type(), "type of QQ")
	assert.Equal(t, 3, getLen(t, s1, "QQ"), "len, s1, non-recursed")
	assert.Equal(t, 3, getLen(t, s1.Recurse("QQ")), "len, s1, recursed")
	s = CombineSources(s, s1)
	s2, err := UnmarshalFile("rdata2.yaml", WithFS(content))
	assert.Equal(t, 3, getLen(t, s2, "QQ"), "len, s2, non-recursed")
	assert.Equal(t, 3, getLen(t, s2.Recurse("QQ")), "len, s2, recursed")
	require.NoError(t, err, "open rdata2.yaml")
	s = CombineSources(s, s2)
	s = s.Recurse("QQ")
	s = MultiSourceSetFirst(true).
		Combine(MultiSourceSetCombine(true)).
		Apply(s)
	if assert.Equal(t, 6, getLen(t, s), "length") {
		assert.Equal(t, "a", getString(t, s.Recurse("0")), "first element w/recurse")
		assert.Equal(t, "a", getString(t, s, "0"), "first element w/recurse")
		assert.Equal(t, "b", getString(t, s.Recurse("1")), "2nd element w/recurse")
		assert.Equal(t, "b", getString(t, s, "1"), "2nd element w/recurse")
		assert.Equal(t, "c", getString(t, s.Recurse("2")), "3rd element w/recurse")
		assert.Equal(t, "c", getString(t, s, "2"), "3rd element w/recurse")
		assert.Equal(t, "d", getString(t, s.Recurse("3")), "4th element w/recurse")
		assert.Equal(t, "d", getString(t, s, "3"), "4th element w/recurse")
		assert.Equal(t, "e", getString(t, s.Recurse("4")), "5th element w/recurse")
		assert.Equal(t, "e", getString(t, s, "4"), "5th element w/recurse")
		assert.Equal(t, "f", getString(t, s.Recurse("5")), "last element w/recurse")
		assert.Equal(t, "f", getString(t, s, "5"), "last element w/recurse")
	}
}

func TestRegression02(t *testing.T) {
	s1, err := UnmarshalFile("rdata1.yaml", WithFS(content))
	require.NoError(t, err, "unmarshal rdata1")
	m := NewMultiSource(s1)
	n := m.Recurse("foo")
	assert.True(t, n == nil, "recurse foo is nil")
}

func getLen(t *testing.T, s Source, args ...string) int {
	l, err := s.Len(args...)
	require.NoError(t, err, "getLen")
	return l
}

func getString(t *testing.T, s Source, args ...string) string {
	require.NotNil(t, s, "getstring s")
	v, err := s.GetString(args...)
	require.NoError(t, err, "getLen")
	return v
}
