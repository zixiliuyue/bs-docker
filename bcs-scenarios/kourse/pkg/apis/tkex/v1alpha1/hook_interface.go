

package v1alpha1

// ----------------------------------------- GameDeployment ----------------------------------------

// GetPreDeleteHook returns predelete hook spec
func (g *GameDeployment) GetPreDeleteHook() *HookStep {
	return g.Spec.PreDeleteUpdateStrategy.Hook
}

// GetPreDeleteHookConditions returns predelete hook conditions
func (g *GameDeploymentStatus) GetPreDeleteHookConditions() []PreDeleteHookCondition {
	return g.PreDeleteHookConditions
}

// SetPreDeleteHookConditions sets predelete hook conditions
func (g *GameDeploymentStatus) SetPreDeleteHookConditions(newConditions []PreDeleteHookCondition) {
	g.PreDeleteHookConditions = newConditions
}

// GetPreInplaceHook returns preinplace hook spec
func (g *GameDeployment) GetPreInplaceHook() *HookStep {
	return g.Spec.PreInplaceUpdateStrategy.Hook
}

// GetPreInplaceHookConditions returns preinplace hook conditions
func (g *GameDeploymentStatus) GetPreInplaceHookConditions() []PreInplaceHookCondition {
	return g.PreInplaceHookConditions
}

// SetPreInplaceHookConditions sets preinplace hook conditions
func (g *GameDeploymentStatus) SetPreInplaceHookConditions(newConditions []PreInplaceHookCondition) {
	g.PreInplaceHookConditions = newConditions
}

// GetPostInplaceHook returns post inplace hook spec
func (g *GameDeployment) GetPostInplaceHook() *HookStep {
	return g.Spec.PostInplaceUpdateStrategy.Hook
}

// GetPostInplaceHookConditions returns post inplace hook conditions
func (g *GameDeploymentStatus) GetPostInplaceHookConditions() []PostInplaceHookCondition {
	return g.PostInplaceHookConditions
}

// SetPostInplaceHookConditions set post inplace hook conditions
func (g *GameDeploymentStatus) SetPostInplaceHookConditions(newConditions []PostInplaceHookCondition) {
	g.PostInplaceHookConditions = newConditions
}

// ----------------------------------------- GameStatefulSet ----------------------------------------

// GetPreDeleteHook returns predelete hook spec
func (g *GameStatefulSet) GetPreDeleteHook() *HookStep {
	return g.Spec.PreDeleteUpdateStrategy.Hook
}

// GetPreDeleteHookConditions returns predelete hook conditions
func (g *GameStatefulSetStatus) GetPreDeleteHookConditions() []PreDeleteHookCondition {
	return g.PreDeleteHookConditions
}

// SetPreDeleteHookConditions sets predelete hook conditions
func (g *GameStatefulSetStatus) SetPreDeleteHookConditions(newConditions []PreDeleteHookCondition) {
	g.PreDeleteHookConditions = newConditions
}

// GetPreInplaceHook returns preinplace hook spec
func (g *GameStatefulSet) GetPreInplaceHook() *HookStep {
	return g.Spec.PreInplaceUpdateStrategy.Hook
}

// GetPreInplaceHookConditions returns preinplace hook conditions
func (g *GameStatefulSetStatus) GetPreInplaceHookConditions() []PreInplaceHookCondition {
	return g.PreInplaceHookConditions
}

// SetPreInplaceHookConditions sets preinplace hook conditions
func (g *GameStatefulSetStatus) SetPreInplaceHookConditions(newConditions []PreInplaceHookCondition) {
	g.PreInplaceHookConditions = newConditions
}

// GetPostInplaceHook returns post inplace hook spec
func (g *GameStatefulSet) GetPostInplaceHook() *HookStep {
	return g.Spec.PostInplaceUpdateStrategy.Hook
}

// GetPostInplaceHookConditions returns post inplace hook conditions
func (g *GameStatefulSetStatus) GetPostInplaceHookConditions() []PostInplaceHookCondition {
	return g.PostInplaceHookConditions
}

// SetPostInplaceHookConditions set post inplace hook conditions
func (g *GameStatefulSetStatus) SetPostInplaceHookConditions(newConditions []PostInplaceHookCondition) {
	g.PostInplaceHookConditions = newConditions
}
