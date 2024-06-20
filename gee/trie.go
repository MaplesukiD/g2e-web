package gee

import "strings"

//前缀树trie

type node struct {
	//待匹配路由
	pattern string
	//路由路径一部分
	part string
	//子节点
	children []*node
	//是否精准匹配，part含有:或*时为true
	isWild bool
}

// matchChild 返回第一个匹配的子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 返回所有匹配的节点的列表
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 向树中插入新的路由模式
func (n *node) insert(pattern string, parts []string, height int) {
	//递归出口
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search 搜索和给定路由模式parts匹配的节点
func (n *node) search(parts []string, height int) *node {
	//递归出口
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
