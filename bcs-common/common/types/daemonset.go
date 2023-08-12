

package types

// BcsDaemonset xxx
type BcsDaemonset struct {
	TypeMeta `json:",inline"`
	// AppMeta               `json:",inline"`
	ObjectMeta    `json:"metadata"`
	Spec          ReplicaControllerSpec `json:"spec"`
	RestartPolicy RestartPolicy         `json:"restartPolicy,omitempty"`
	KillPolicy    KillPolicy            `json:"killPolicy,omitempty"`
}
