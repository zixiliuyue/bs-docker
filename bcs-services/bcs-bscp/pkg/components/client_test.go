

package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefineCode(t *testing.T) {
	code, err := refineCode("00")
	assert.NoError(t, err)
	assert.True(t, code == 0)
}
