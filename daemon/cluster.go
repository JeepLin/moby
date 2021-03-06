package daemon // import "moby/daemon"

import (
	apitypes "moby/api/types"
	"moby/api/types/filters"
	lncluster "github.com/docker/libnetwork/cluster"
)

// Cluster is the interface for moby/daemon/cluster.(*Cluster).
type Cluster interface {
	ClusterStatus
	NetworkManager
	SendClusterEvent(event lncluster.ConfigEventType)
}

// ClusterStatus interface provides information about the Swarm status of the Cluster
type ClusterStatus interface {
	IsAgent() bool
	IsManager() bool
}

// NetworkManager provides methods to manage networks
type NetworkManager interface {
	GetNetwork(input string) (apitypes.NetworkResource, error)
	GetNetworks(filters.Args) ([]apitypes.NetworkResource, error)
	RemoveNetwork(input string) error
}
