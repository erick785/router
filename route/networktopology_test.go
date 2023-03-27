package route

import (
	"testing"
)

func TestVerifyNetWorkTopology(t *testing.T) {
	netTopology := &NetworkTopology{
		list: []*Link{
			{srcNode: "1", dstNodes: []string{"2", "3", "5"}},
			{srcNode: "2", dstNodes: []string{"1", "3", "7"}},
			{srcNode: "3", dstNodes: []string{"1", "2", "5", "6"}},
			{srcNode: "4", dstNodes: []string{"6", "7"}},
			{srcNode: "5", dstNodes: []string{"1", "3"}},
			{srcNode: "6", dstNodes: []string{"3", "4"}},
			{srcNode: "7", dstNodes: []string{"4", "2"}},
		},
	}
	t.Log(netTopology.getLinkList())

	n := netTopology.verifyNetWorkTopology("1")

	t.Log(n.getLinkList())

}

func TestAddLink(t *testing.T) {
	netTopology := newNetworkTopology()

	netTopology.addLink(&Link{srcNode: "1", dstNodes: []string{"2", "3", "5"}})
	t.Log(netTopology.getLinkList())
	netTopology.addLink(&Link{srcNode: "2", dstNodes: []string{"1", "3", "7"}})
	t.Log(netTopology.getLinkList())
	netTopology.addLink(&Link{srcNode: "3", dstNodes: []string{"2", "5", "6"}})
	t.Log(netTopology.getLinkList())
	netTopology.addLink(&Link{srcNode: "4", dstNodes: []string{"6", "7"}})
	t.Log(netTopology.getLinkList())
	netTopology.addLink(&Link{srcNode: "5", dstNodes: []string{"3"}})
	t.Log(netTopology.getLinkList())
	netTopology.addLink(&Link{srcNode: "6", dstNodes: []string{"3", "4"}})
	t.Log(netTopology.getLinkList())
	netTopology.addLink(&Link{srcNode: "7", dstNodes: []string{"4", "2"}})
	t.Log(netTopology.getLinkList())

	// [{1 [2 3 5]} {2 [1 3 7]} {3 [1 2 5 6]} {5 [1 3]} {7 [2 4]} {6 [3 4]} {4 [6 7]}]
}

func TestAddLink2(t *testing.T) {
	netTopology := newNetworkTopology()
	netTopology.addLink(&Link{srcNode: "1", dstNodes: []string{"2"}})
	t.Log(netTopology.getLinkList())

	netTopology.addLink(&Link{srcNode: "1", dstNodes: []string{"3"}})
	t.Log(netTopology.getLinkList())

	netTopology.addLink(&Link{srcNode: "1", dstNodes: []string{"5"}})
	t.Log(netTopology.getLinkList())
}

func TestAddLink3(t *testing.T) {
	netTopology := newNetworkTopology()
	netTopology.addLink(&Link{srcNode: "1", dstNodes: []string{"2"}})
	t.Log(netTopology.getLinkList())

	netTopology.addLink(&Link{srcNode: "1", dstNodes: []string{"3", "5"}})
	t.Log(netTopology.getLinkList())
}
