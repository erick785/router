package route

import "strings"

// Link Node Link
type Link struct {
	srcNode  string
	dstNodes []string
}

// NewNodeLink initialization
func NewNodeLink(srcNode string, dstNodes []string) *Link {
	return &Link{srcNode: srcNode, dstNodes: dstNodes}
}

// GetSrcNode get srcNode
func (l *Link) GetSrcNode() string {
	return l.srcNode
}

// GetDstNodes get dstNodes
func (l *Link) GetDstNodes() []string {
	return l.dstNodes
}

// DstNodeIsExist dstNode is exist
func (l *Link) DstNodeIsExist(dstNode string) bool {
	for _, tmpDstNode := range l.dstNodes {
		if strings.EqualFold(tmpDstNode, dstNode) {
			return true
		}
	}
	return false
}
