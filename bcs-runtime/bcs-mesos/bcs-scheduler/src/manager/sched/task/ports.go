

package task

import "github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/mesosproto/mesos"

func getOfferPorts(offer *mesos.Offer) (ports []uint64) {
	for _, resource := range offer.Resources {
		if resource.GetName() == "ports" {
			for _, rang := range resource.GetRanges().GetRange() {
				for i := rang.GetBegin(); i <= rang.GetEnd(); i++ {
					ports = append(ports, i)
				}
			}
		}
	}
	return ports
}
