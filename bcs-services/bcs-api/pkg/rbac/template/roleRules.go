

package template

import (
	rbacv1 "k8s.io/api/rbac/v1"
)

// 集群管理员角色，可以管理所有资源
var clusterManageRules = []rbacv1.PolicyRule{
	{
		Verbs:     []string{"*"},
		APIGroups: []string{"*"},
		Resources: []string{"*"},
	},
	{
		NonResourceURLs: []string{"*"},
		Verbs:           []string{"*"},
	},
}

// 集群只读用户角色，能查看所有资源，但不能更改
var clusterReadonlyRules = []rbacv1.PolicyRule{
	{
		Verbs:     []string{"get", "list", "watch"},
		APIGroups: []string{"*"},
		Resources: []string{"*"},
	},
}

// services只读角色，能查看services， 但不能更改
var servicesViewRules = []rbacv1.PolicyRule{
	{
		Verbs:     []string{"get", "list", "watch"},
		APIGroups: []string{"*"},
		Resources: []string{"services", "endpoints"},
	},
}

// services管理角色，能管理所有services
var servicesManageRules = []rbacv1.PolicyRule{
	{
		Verbs:     []string{"*"},
		APIGroups: []string{"*"},
		Resources: []string{"services", "endpoints"},
	},
}

// workloads只读角色，能查看pods,deployments,statefulsets等资源
var workloadsViewRules = []rbacv1.PolicyRule{
	{
		Verbs:     []string{"get", "list", "watch"},
		APIGroups: []string{"*"},
		Resources: []string{"pods", "pods/attach", "pods/exec", "pods/portforward", "pods/proxy", "replicationcontrollers",
			"replicationcontrollers/scale", "daemonsets",
			"deployments", "deployments/rollback", "deployments/scale", "replicasets", "replicasets/scale", "statefulsets",
			"cronjobs", "jobs", "daemonsets", "deployments",
			"deployments/rollback", "deployments/scale", "replicasets", "replicasets/scale", "replicationcontrollers/scale",
			"horizontalpodautoscalers"},
	},
	{
		Verbs:     []string{"get", "list", "watch"},
		APIGroups: []string{"*"},
		Resources: []string{"limitranges", "pods/log", "pods/status", "replicationcontrollers/status", "resourcequotas",
			"resourcequotas/status", "bindings"},
	},
}

// workloads管理角色， 能查看和更改pods,deployments,statefulsets等资源
var workloadsManageRules = []rbacv1.PolicyRule{
	{
		Verbs:     []string{"*"},
		APIGroups: []string{"*"},
		Resources: []string{"pods", "pods/attach", "pods/exec", "pods/portforward", "pods/proxy", "replicationcontrollers",
			"replicationcontrollers/scale", "daemonsets",
			"deployments", "deployments/rollback", "deployments/scale", "replicasets", "replicasets/scale", "statefulsets",
			"cronjobs", "jobs", "daemonsets", "deployments",
			"deployments/rollback", "deployments/scale", "replicasets", "replicasets/scale", "replicationcontrollers/scale",
			"horizontalpodautoscalers"},
	},
	{
		Verbs:     []string{"get", "list", "watch"},
		APIGroups: []string{"*"},
		Resources: []string{"limitranges", "pods/log", "pods/status", "replicationcontrollers/status", "resourcequotas",
			"resourcequotas/status", "bindings"},
	},
}
