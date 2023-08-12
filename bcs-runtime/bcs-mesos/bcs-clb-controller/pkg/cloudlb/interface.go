

package cloudlb

import (
	cloudListenerType "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/network/v1"
)

const (
	// QCloudLB tencent cloud lb
	QCloudLB = "qcloudclb"
)

// Interface definition for cloud infrastructure
type Interface interface {
	LoadConfig() error                                                             // load all config item from Env (Subnet, security group)
	CreateLoadbalance() (*cloudListenerType.CloudLoadBalancer, error)              // create new loadbalancer if needed
	DescribeLoadbalance(name string) (*cloudListenerType.CloudLoadBalancer, error) // get loadbalancer by name, id or arn
	Update(old, cur *cloudListenerType.CloudListener) error                        // update event
	Add(ls *cloudListenerType.CloudListener) error                                 // new listener event
	Delete(ls *cloudListenerType.CloudListener) error                              // listener delete event
	ListListeners() ([]*cloudListenerType.CloudListener,
		error) // list all listener on clb instance controlled by this controller
}
