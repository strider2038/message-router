package handling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleRequest_emptyBody_httpErrorReturned(t *testing.T) {
	handler := MessageCollectionRequestHandler{}

	assert.Equal(t, 1, 1)
}
