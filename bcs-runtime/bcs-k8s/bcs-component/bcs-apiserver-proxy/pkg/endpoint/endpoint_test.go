

package endpoint

import (
	"testing"
)

func getEndpointsClient() ClusterEndpointsIP {
	k8sConfig := getK8sConfig()
	if k8sConfig == nil {
		return nil
	}

	clusterEndpointsClient, err := NewEndpointsClient(WithK8sConfig(*k8sConfig), WithDebug(true))
	if err != nil {
		return nil
	}

	return clusterEndpointsClient
}

func TestEndpoints_GetClusterEndpoints(t *testing.T) {

	client := getEndpointsClient()
	if client == nil {
		t.Fatalf("getEndpointsClient failed")
		return
	}

	t.Logf("%+v", client)

	endpoints, err := client.GetClusterEndpoints()
	if err != nil {
		t.Fatalf("GetClusterEndpoints failed: %v", err)
		return
	}
	t.Logf("GetClusterEndpoints %+v", endpoints)
}
