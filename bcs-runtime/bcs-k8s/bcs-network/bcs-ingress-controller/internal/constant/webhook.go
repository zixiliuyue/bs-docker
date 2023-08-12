

package constant

const (
	// ValidateMsgUnknownErr ingress validate get unknown error
	ValidateMsgUnknownErr = "unknown err: %+v"
	// ValidateMsgEmptySvc ingress validate get empty service
	ValidateMsgEmptySvc = "rule[%d] should have service"
	// ValidateMsgNotFoundSvc ingress validate can not get specified service
	ValidateMsgNotFoundSvc = "rule[%d] service '%s/%s' not found"
	// ValidateMsgInvalidWorkload ingress validate get invalid workload
	ValidateMsgInvalidWorkload = "port mapping[%d]'s workload have empty workload kind/namespace/name "
	// ValidateMsgEmptyWorkload ingress validate get empty workload
	ValidateMsgEmptyWorkload = "port mapping[%d]'s workload not found"

	// PortConflictMsg new create ingress/portpool has port conflict with existed
	PortConflictMsg = "conflict with kind[%s] namespace[%s] name[%s] on lbID[%s]"
)
