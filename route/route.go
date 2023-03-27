package route

import (
	"sync"
)

// Route 路由算法
type Route struct {
	sync.RWMutex
	netTopologyChange bool
	netTopology       *NetworkTopology
}

// NewRoute initialization
func NewRoute() *Route {
	return &Route{
		netTopologyChange: false,
		netTopology:       newNetworkTopology(),
	}
}

// UpdateNetworkTopology Update Network Topology
func (r *Route) UpdateNetworkTopology(link *Link) bool {
	r.RLock()
	defer r.RUnlock()
	if r.netTopology.linkIsExist(link) {
		return false
	}

	if len(link.dstNodes) == 0 {
		if !r.netTopology.srcNodeIsExist(link.srcNode) {
			return false
		}
	}

	r.netTopology.deleteLink(link)

	if len(link.dstNodes) != 0 {
		r.netTopology.addLink(link)
	}

	r.netTopology.deleteNilLink()

	r.netTopologyChange = true
	return true
}

func (r *Route) GetRoutes(srcNode string, dstNode string) [][]string {
	r.RLock()
	defer r.RUnlock()

	var routes [][]string
	var visited []string
	visited = append(visited, srcNode)

	dfs(r.netTopology.getLinkList(), visited, &routes, dstNode)

	// todo 过滤重复的路径的
	return routes
}

// GetNetworkTopology get Network Topology
func (r *Route) GetNetworkTopology() []Link {
	r.RLock()
	defer r.RUnlock()
	return r.netTopology.getLinkList()
}

func dfs(nodes []Link, visited []string, routes *[][]string, dstNode string) {
	lastVisited := visited[len(visited)-1]

	if lastVisited == dstNode {
		// Found a route to the destination
		*routes = append(*routes, append([]string(nil), visited...))
		return
	}

	// Iterate over the neighbors of the last visited node
	for _, neighbor := range getNeighbors(nodes, lastVisited) {
		if !contains(visited, neighbor) {
			visited = append(visited, neighbor)
			dfs(nodes, visited, routes, dstNode)
			visited = visited[:len(visited)-1]
		}
	}
}

func getNeighbors(nodes []Link, srcNode string) []string {
	for _, node := range nodes {
		if node.srcNode == srcNode {
			return node.dstNodes
		}
	}

	return []string{}
}

func contains(arr []string, val string) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}
