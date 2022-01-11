package nflex

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOffset(t *testing.T) {
	var s Source
	s, err := UnmarshalFile("rdata1.yaml", WithFS(content))
	require.NoError(t, err, "open rdata1.yaml")
	s = s.Recurse("QQ")
	assert.Equal(t, "a", getString(t, s, "0"), "1st element, no offset")
	s = WithOffset(s, 2)
	assert.Equal(t, "b", getString(t, s, "3"), "2nd element, offset down")
}
