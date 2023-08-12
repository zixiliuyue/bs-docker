

package azure

import (
	"log"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	networkextensionv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
)

// TestEnsureApplicationGatewayListener test ensure azure application gateway listener
// NOCC:tosa/fn_length(测试函数)
func TestEnsureApplicationGatewayListener(t *testing.T) {
	listener := &networkextensionv1.Listener{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "listener-3"},
		Spec: networkextensionv1.ListenerSpec{
			LoadbalancerID:    "appgw",
			Port:              80,
			EndPort:           0,
			Protocol:          "HTTP",
			ListenerAttribute: nil,
			Certificate:       nil,
			TargetGroup:       nil,
			Rules: []networkextensionv1.ListenerRule{
				{
					Domain: "",
					Path:   "/test",
					ListenerAttribute: &networkextensionv1.IngressListenerAttribute{HealthCheck: &networkextensionv1.
						ListenerHealthCheck{
						Enabled:       true,
						Timeout:       2,
						IntervalTime:  15,
						UnHealthNum:   1,
						HTTPCode:      31,
						HTTPCheckPath: "/test",
					}},
					// TargetGroup: &networkextensionv1.ListenerTargetGroup{
					// 	TargetGroupProtocol: "HTTP",
					// 	Backends: []networkextensionv1.ListenerBackend{
					// 		{
					// 			IP:   "10.224.0.6",
					// 			Port: 30284,
					// 		},
					// 	},
					// },
				},
			},
		},
		Status: networkextensionv1.ListenerStatus{},
	}

	// listener2 := &networkextensionv1.Listener{
	// 	TypeMeta:   metav1.TypeMeta{},
	// 	ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "listener-4"},
	// 	Spec: networkextensionv1.ListenerSpec{
	// 		LoadbalancerID:    "appgw",
	// 		Port:              80,
	// 		EndPort:           0,
	// 		Protocol:          "HTTP",
	// 		ListenerAttribute: nil,
	// 		Certificate:       nil,
	// 		TargetGroup:       nil,
	// 		Rules: []networkextensionv1.ListenerRule{
	// 			{
	// 				Domain: "",
	// 				Path:   "/test2",
	// 				TargetGroup: &networkextensionv1.ListenerTargetGroup{
	// 					TargetGroupProtocol: "HTTP",
	// 					Backends: []networkextensionv1.ListenerBackend{
	// 						{IP: "10.224.0.6",
	// 							Port: 30284},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	Status: networkextensionv1.ListenerStatus{},
	// }

	alb, err := NewAlb()
	if err != nil {
		log.Fatal(err)
	}

	err = alb.ensureApplicationGatewayListener("", []*networkextensionv1.Listener{listener})
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(res)
	//
	// res, err = alb.ensureApplicationGatewayListener("", listener2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(res)
}

// TestDeleteAppGwListener test delete azure application gateway listener
func TestDeleteAppGwListener(t *testing.T) {
	// listener := &networkextensionv1.Listener{
	// 	TypeMeta:   metav1.TypeMeta{},
	// 	ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "listener-4"},
	// 	Spec: networkextensionv1.ListenerSpec{
	// 		LoadbalancerID:    "appgw",
	// 		Port:              80,
	// 		EndPort:           0,
	// 		Protocol:          "HTTP",
	// 		ListenerAttribute: nil,
	// 		Certificate:       nil,
	// 		TargetGroup:       nil,
	// 		Rules: []networkextensionv1.ListenerRule{
	// 			{
	// 				Domain: "",
	// 				Path:   "/test2",
	// 				TargetGroup: &networkextensionv1.ListenerTargetGroup{
	// 					TargetGroupProtocol: "HTTP",
	// 					Backends: []networkextensionv1.ListenerBackend{
	// 						{IP: "10.224.0.6",
	// 							Port: 30284},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	Status: networkextensionv1.ListenerStatus{},
	// }
	// 	Status: networkextensionv1.ListenerStatus{},
	// }
	listener := &networkextensionv1.Listener{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "listener-3"},
		Spec: networkextensionv1.ListenerSpec{
			LoadbalancerID:    "appgw",
			Port:              80,
			EndPort:           0,
			Protocol:          "HTTP",
			ListenerAttribute: nil,
			Certificate:       nil,
			TargetGroup:       nil,
			Rules: []networkextensionv1.ListenerRule{
				{
					Domain: "",
					Path:   "/test",
					ListenerAttribute: &networkextensionv1.IngressListenerAttribute{HealthCheck: &networkextensionv1.
						ListenerHealthCheck{
						Enabled:       true,
						Timeout:       2,
						IntervalTime:  15,
						UnHealthNum:   1,
						HTTPCode:      31,
						HTTPCheckPath: "/test",
					}},
				},
			},
		},
		Status: networkextensionv1.ListenerStatus{},
	}

	alb, err := NewAlb()
	if err != nil {
		log.Fatal(err)
	}

	err = alb.deleteApplicationGatewayListener("", []*networkextensionv1.Listener{listener})
	if err != nil {
		log.Fatal(err)
	}

}

// TestEnsureLBListener test ensure load balancer listener
func TestEnsureLBListener(t *testing.T) {
	listener := &networkextensionv1.Listener{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "listener-3"},
		Spec: networkextensionv1.ListenerSpec{
			LoadbalancerID:    "lb",
			Port:              80,
			EndPort:           0,
			Protocol:          "TCP",
			ListenerAttribute: nil,
			Certificate:       nil,
			TargetGroup: &networkextensionv1.ListenerTargetGroup{
				TargetGroupProtocol: "TCP",
				Backends: []networkextensionv1.ListenerBackend{
					{IP: "10.224.0.5",
						Port: 32683},
				},
			},
		},
		Status: networkextensionv1.ListenerStatus{},
	}

	alb, err := NewAlb()
	if err != nil {
		log.Fatal(err)
	}
	res, err := alb.ensureLoadBalancerListener("", []*networkextensionv1.Listener{listener})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}

// TestDeleteLBListener test delete lb listener
func TestDeleteLBListener(t *testing.T) {
	listener := &networkextensionv1.Listener{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "listener-3"},
		Spec: networkextensionv1.ListenerSpec{
			LoadbalancerID:    "lb",
			Port:              80,
			EndPort:           0,
			Protocol:          "TCP",
			ListenerAttribute: nil,
			Certificate:       nil,
			TargetGroup: &networkextensionv1.ListenerTargetGroup{
				TargetGroupProtocol: "TCP",
				Backends: []networkextensionv1.ListenerBackend{
					{IP: "10.224.0.5",
						Port: 32683},
				},
			},
		},
		Status: networkextensionv1.ListenerStatus{},
	}

	alb, err := NewAlb()
	if err != nil {
		log.Fatal(err)
	}
	err = alb.deleteLoadBalancerListener("", []*networkextensionv1.Listener{listener})
	if err != nil {
		log.Fatal(err)
	}
}
