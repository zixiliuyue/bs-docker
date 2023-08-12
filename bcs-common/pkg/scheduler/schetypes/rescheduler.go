

package types

// ReschedMesosAddress xxx
// mesos address
type ReschedMesosAddress struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
}

// ReschedMesosInfo xxx
// mesos info
type ReschedMesosInfo struct {
	Address  ReschedMesosAddress `json:"address"`
	Hostname string              `json:"hostname"`
	Id       string              `json:"id"`
	IP       int                 `json:"ip"`
	Pid      string              `json:"pid"`
	Port     int                 `json:"port"`
	Version  string              `json:"version"`
}

// Framework xxx
// mesos framework info
type Framework struct {
	ID string `json:"id"`
}
