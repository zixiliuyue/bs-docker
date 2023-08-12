

package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	bluekingio "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/apis/blueking.io"
)

// SchemeGroupVersion xxx
var SchemeGroupVersion = schema.GroupVersion{Group: bluekingio.GroupName, Version: "v1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
