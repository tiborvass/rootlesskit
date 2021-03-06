package network

import (
	"github.com/rootless-containers/rootlesskit/pkg/common"
)

// ParentDriver is called from the parent namespace
type ParentDriver interface {
	// NetworkMode returns common.NetworkMode. Must not be HostNetwork.
	NetworkMode() common.NetworkMode
	// MTU returns MTU
	MTU() int
	// ConfigureNetwork sets up Slirp, updates msg, and returns destructor function.
	ConfigureNetwork(childPID int, stateDir string) (netmsg *common.NetworkMessage, cleanup func() error, err error)
}

// ChildDriver is called from the child namespace
type ChildDriver interface {
	ConfigureTap(netmsg common.NetworkMessage) (tap string, err error)
}
