package route

import (
	"testing"
)

func makeRoute() *Route {
	return &Route{
		netTopology: &NetworkTopology{
			list: []*Link{
				&Link{srcNode: "1", dstNodes: []string{"2", "3", "5"}},
				&Link{srcNode: "2", dstNodes: []string{"1", "3"}},
				&Link{srcNode: "3", dstNodes: []string{"1", "2", "5", "6"}},
				&Link{srcNode: "4", dstNodes: []string{"6"}},
				&Link{srcNode: "5", dstNodes: []string{"1", "3"}},
				&Link{srcNode: "6", dstNodes: []string{"3", "4"}},
			},
		},
		netTopologyChange: false,
	}
}

func printNetworkTopologyList(list []*Link, t *testing.T) {
	for _, v := range list {
		t.Log(*v)
	}
}

func TestRoute(t *testing.T) {
	route := makeRoute()

	t.Log("test add Link1")
	if route.UpdateNetworkTopology(&Link{srcNode: "7", dstNodes: []string{"4"}}) {
		//printNetworkTopologyList(route.netTopology.list, t)
		t.Log(route.GetRoutes("1", "7"))

	} else {
		t.Log("not update network topology")
	}

	t.Log("test add Link2")
	if route.UpdateNetworkTopology(&Link{srcNode: "7", dstNodes: []string{"4", "2"}}) {
		//printNetworkTopologyList(route.netTopology.list, t)
		t.Log(route.GetRoutes("1", "7"))
	} else {
		t.Log("not update network topology")
	}

}
