

package pkg

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GeProjectList(t *testing.T) {

	resp, err := NewClientWithConfiguration(context.Background()).GetProject("7da12ea6af35464a8be39961a21e95d9")
	if err != nil {
		log.Fatal(err)
	}

	assert.NotNil(t, resp)
}
