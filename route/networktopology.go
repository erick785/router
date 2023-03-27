package route

import (
	"strings"

	"github.com/erick785/router/util"
)

// NetworkTopology topological graph
type NetworkTopology struct {
	list []*Link
}

func newNetworkTopology() *NetworkTopology {
	return &NetworkTopology{}
}

func (n *NetworkTopology) linkIsExist(link *Link) bool {
	for _, tmpLink := range n.list {
		if strings.EqualFold(tmpLink.srcNode, link.srcNode) {
			if len(tmpLink.dstNodes) == len(link.dstNodes) {
				for _, dstNode := range tmpLink.dstNodes {
					if util.IsStrExist(dstNode, link.dstNodes) {
						return true
					}
				}
			}
		}
	}
	return false
}

func (n *NetworkTopology) srcNodeIsExist(srcNode string) bool {
	for _, tmpLink := range n.list {
		if strings.EqualFold(tmpLink.srcNode, srcNode) {
			return true
		}
	}
	return false
}

func (n *NetworkTopology) deleteNilLink() {
	for _, tmpLink := range n.list {
		if len(tmpLink.dstNodes) == 0 {
			for k, tLink := range n.list {
				if tLink == tmpLink {
					n.list = append(n.list[:k], n.list[k+1:]...)
				}
			}
		}
	}
}

func (n *NetworkTopology) deleteLink(link *Link) {
	for k, tmpLink := range n.list {
		if strings.EqualFold(tmpLink.srcNode, link.srcNode) {
			n.list = append(n.list[:k], n.list[k+1:]...)
		}
	}
	//delete srcnode from dstNodes
	for k, tmpLink := range n.list {
		for key, dstNode := range tmpLink.dstNodes {
			if strings.EqualFold(dstNode, link.srcNode) {
				n.list[k].dstNodes = append(n.list[k].dstNodes[:key], n.list[k].dstNodes[key+1:]...)
			}
		}
	}

}

func (n *NetworkTopology) addLink(link *Link) {
	if !n.linkIsExist(link) {
		if !n.srcNodeIsExist(link.srcNode) {
			n.list = append(n.list, link)
		} else {
			for k, tmpLink := range n.list {
				if strings.EqualFold(tmpLink.srcNode, link.srcNode) {
					for _, dstNode := range link.dstNodes {
						if !n.list[k].DstNodeIsExist(dstNode) {
							n.list[k].dstNodes = append(n.list[k].dstNodes, dstNode)
						}
					}
				}
			}
		}
	}
	for k, tmpLink := range n.list {
		for _, tmpDstNode := range link.dstNodes {
			if strings.EqualFold(tmpDstNode, tmpLink.srcNode) {
				if !n.list[k].DstNodeIsExist(link.srcNode) {
					n.list[k].dstNodes = append(n.list[k].dstNodes, link.srcNode)
				}
			}
		}
	}

	for _, dstNode := range link.dstNodes {
		if !n.srcNodeIsExist(dstNode) {
			n.list = append(n.list, &Link{srcNode: dstNode, dstNodes: []string{link.srcNode}})
		}
	}
}

func (n *NetworkTopology) getLinkList() []Link {
	tmplist := []Link{}
	for _, v := range n.list {
		tmplist = append(tmplist, *v)
	}
	return tmplist
}

func (n *NetworkTopology) getLink(srcNode string) *Link {
	for _, tmpLink := range n.list {
		if strings.EqualFold(tmpLink.srcNode, srcNode) {
			return tmpLink
		}
	}
	return nil
}

func (n *NetworkTopology) verifyNetWorkTopology(localNode string) *NetworkTopology {
	tmpNetworktopology := &NetworkTopology{}
	link := n.getLink(localNode)
	cache := []string{}
	cache = append(cache, localNode)
	cache = append(cache, link.dstNodes...)
	return n.checkLink(cache, tmpNetworktopology, link, localNode)
}

// 该函数的作用是检查网络拓扑结构中的链路是否存在，并返回一个更新后的网络拓扑结构。
// 如果该链路已经存在于给定的网络拓扑结构中，则返回原始网络拓扑结构，
// 否则将该链路添加到给定的网络拓扑结构中。
func (n *NetworkTopology) checkLink(cache []string, nt *NetworkTopology, link *Link, localNode string) *NetworkTopology {
	if nt.linkIsExist(link) {
		return nt
	}
	nt.list = append(nt.list, link)

	for _, dstNode := range link.dstNodes {
		if localNode == link.srcNode {
			n.checkLink(cache, nt, n.getLink(dstNode), localNode)
		}
		if !util.IsStrExist(dstNode, cache) {
			cache = append(cache, dstNode)
			n.checkLink(cache, nt, n.getLink(dstNode), localNode)

		}
	}
	return nt
}
