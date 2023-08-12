

package template

import (
	rbacv1 "k8s.io/api/rbac/v1"
)

const (
	// ClusterRolePrefix xxx
	ClusterRolePrefix = "bke-"
)

// RoleTemplate xxx
type RoleTemplate struct {
	Name  string
	Rules []rbacv1.PolicyRule
}

// RoleTemplateStore xxx
var RoleTemplateStore map[string]*RoleTemplate

// InitRbacTemplates 初始化所有clusterrole角色，定义每个clusterrole的权限
func InitRbacTemplates() {
	RoleTemplateStore = make(map[string]*RoleTemplate)
	addRoleTemplate("cluster-manage", clusterManageRules)
	addRoleTemplate("cluster-readonly", clusterReadonlyRules)
	addRoleTemplate("services-view", servicesViewRules)
	addRoleTemplate("services-manage", servicesManageRules)
	addRoleTemplate("workloads-view", workloadsViewRules)
	addRoleTemplate("workloads-manage", workloadsManageRules)
}

func addRoleTemplate(roleName string, rules []rbacv1.PolicyRule) {
	roleTemplate := &RoleTemplate{
		Name:  ClusterRolePrefix + roleName,
		Rules: rules,
	}
	RoleTemplateStore[ClusterRolePrefix+roleName] = roleTemplate
}
