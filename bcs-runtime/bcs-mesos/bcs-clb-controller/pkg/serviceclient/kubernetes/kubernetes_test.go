

package kubernetes

import (
	"testing"

	k8scorev1 "k8s.io/api/core/v1"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetIndexFromStatefulSetName(t *testing.T) {
	tests := []struct {
		Name  string
		IsErr bool
		Index int
	}{
		{
			Name:  "pvpserver-0",
			IsErr: false,
			Index: 0,
		},
		{
			Name:  "pvpserver",
			IsErr: true,
			Index: -1,
		},
	}
	for _, test := range tests {
		outIndex, err := getIndexFromStatefulSetName(test.Name)
		if (test.IsErr && err == nil) || (!test.IsErr && err != nil) {
			t.Errorf("IsErr %v, but err %s", test.IsErr, err.Error())
		}
		if !test.IsErr {
			if outIndex != test.Index {
				t.Errorf("expect index %d, but get %d", test.Index, outIndex)
			}
		}
	}
}

func TestSortStatefulSetPod(t *testing.T) {
	tests := []struct {
		podsBefore []*k8scorev1.Pod
		podsAfter  []*k8scorev1.Pod
	}{
		{
			podsBefore: []*k8scorev1.Pod{
				{
					ObjectMeta: k8smetav1.ObjectMeta{
						Name: "pvp-3",
					},
				},
				{
					ObjectMeta: k8smetav1.ObjectMeta{
						Name: "pvp-2",
					},
				},
				{
					ObjectMeta: k8smetav1.ObjectMeta{
						Name: "pvp-9",
					},
				},
				{
					ObjectMeta: k8smetav1.ObjectMeta{
						Name: "pvp-110",
					},
				},
			},
			podsAfter: []*k8scorev1.Pod{
				{
					ObjectMeta: k8smetav1.ObjectMeta{
						Name: "pvp-2",
					},
				},
				{
					ObjectMeta: k8smetav1.ObjectMeta{
						Name: "pvp-3",
					},
				},
				{
					ObjectMeta: k8smetav1.ObjectMeta{
						Name: "pvp-9",
					},
				},
				{
					ObjectMeta: k8smetav1.ObjectMeta{
						Name: "pvp-110",
					},
				},
			},
		},
	}
	for _, test := range tests {
		sortStatefulSetPod(test.podsBefore)
		for index, pod := range test.podsBefore {
			if pod.Name != test.podsAfter[index].Name {
				t.Errorf("expect %v but get %v", test.podsAfter[index], pod)
			}
		}
	}
}
