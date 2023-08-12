

package util

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

// GetPodRevision returns revision hash of this pod.
func GetPodRevision(pod metav1.Object) string {
	return pod.GetLabels()[appsv1.ControllerRevisionHashLabelKey]
}

// GetPodsRevisions return revision hash set of these pods.
func GetPodsRevisions(pods []*v1.Pod) sets.String {
	revisions := sets.NewString()
	for _, p := range pods {
		revisions.Insert(GetPodRevision(p))
	}
	return revisions
}
