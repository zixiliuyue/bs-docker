

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	monitorv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/monitor/v1"
)

// ServiceMonitorReconciler reconciles a ServiceMonitor object
type ServiceMonitorReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=monitor.bkbcs.tencent.com,resources=servicemonitors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitor.bkbcs.tencent.com,resources=servicemonitors/status,verbs=get;update;patch

func (r *ServiceMonitorReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("servicemonitor", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *ServiceMonitorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitorv1.ServiceMonitor{}).
		Complete(r)
}
