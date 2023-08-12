

package bcs

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProject(t *testing.T) {
	ctx := context.Background()

	projects, err := ListAuthorizedProjects(ctx, "")
	assert.NoError(t, err)
	fmt.Println(projects)
}
