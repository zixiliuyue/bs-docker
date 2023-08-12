

package transfer

import (
	"context"
)

// Interface defines the interface of transfer webhook
type Interface interface {
	Transfer(context.Context, []byte) ([]byte, error)
}
