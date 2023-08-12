

// Package controller xxx
package controller

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/queue"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bmsf-mesh/bmsf-mesos-adapter/controller/appnode"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bmsf-mesh/bmsf-mesos-adapter/controller/appsvc"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager, svcQ queue.Queue, nodeQ queue.Queue) error {
	if err := appsvc.Add(m, svcQ); err != nil {
		return err
	}
	if err := appnode.Add(m, nodeQ); err != nil {
		return err
	}
	return nil
}
