

package gsepipeline

import (
	"time"

	"github.com/docker/engine-api/types"
)

// Storage save data to gse interface
type Storage interface {
	AddStats(msg LogMsg) error
}

type eventActor struct {
	ID         string            `json:"ID"`
	Attributes map[string]string `json:"Attributes"`
}

// EventJSON event json
type EventJSON struct {
	Status   string     `json:"status"`
	ID       string     `json:"id"`
	From     string     `json:"from"`
	Type     string     `json:"Type"`
	Action   string     `json:"Action"`
	Actor    eventActor `json:"Actor"`
	Time     uint64     `json:"time"`
	TimeNano uint64     `json:"timeNano"`

	line string
}

// ContainerMount definition for container mount
type ContainerMount struct {
	Source      string
	Destination string
}

// ContainerReference Container reference contains enough information to uniquely identify a container
type ContainerReference struct {
	// The container id
	ID string `json:"id,omitempty"`

	CreateTime time.Time `json:"created,omitempty"`

	// The absolute name of the container. This is unique on the machine.
	Name string `json:"name,omitempty"`

	Image     string `json:"image,omitempty"`
	IPAddress string `json:"host,omitempty"`

	ContainerRootDirector string `json:"container_root,omitempty"`

	// Other names by which the container is known within a certain namespace.
	// This is unique within that namespace.
	Aliases []string `json:"aliases,omitempty"`

	// Namespace under which the aliases of a container are unique.
	// An example of a namespace is "docker" for Docker containers.
	Namespace string `json:"namespace,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`

	ContainerInfo types.ContainerJSON `json:"-"`
}

// LogMsg log object
type LogMsg struct {
	FileName     string              `json:"filename,omitempty"`
	Timestamp    time.Time           `json:"timestamp"`
	Stream       string              `json:"stream"`
	Log          interface{}         `json:"log"`
	DataID       uint64              `json:"dataid"`
	ContainerRef *ContainerReference `json:"container_info,omitempty"`
}
