

package kubedriver

import (
	"reflect"

	restful "github.com/emicklei/go-restful"
)

// RouteByMethodName accepts a dynamic method name, this function uses reflect module
func RouteByMethodName(w *restful.WebService, methodName string, subPath string) *restful.RouteBuilder {
	methodValue := reflect.ValueOf(w).MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(subPath)})
	return methodValue[0].Interface().(*restful.RouteBuilder)
}
