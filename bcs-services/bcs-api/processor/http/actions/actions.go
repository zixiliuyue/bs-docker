

// Package actions xxx
package actions

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/http/httpserver"
)

// Action restful action struct
type Action httpserver.Action

var apiActions = []*httpserver.Action{}

// RegisterAction register action to actions
func RegisterAction(action Action) {
	apiActions = append(apiActions, httpserver.NewAction(action.Verb, action.Path, action.Params, action.Handler))
}

// GetApiAction xxx
func GetApiAction() []*httpserver.Action {
	return apiActions
}
