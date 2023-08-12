

package types

// BcsTaskgroup xxx
// BcsPod pod for bcs
type BcsTaskgroup struct {
	TypeMeta      `json:",inline"`
	ObjectMeta    `json:"metadata"`
	PodSpec       `json:"spec"`
	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty"`
}
