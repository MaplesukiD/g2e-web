package gee

import (
	"fmt"
	"strings"
	"testing"
)

// matchChild和matchChildren方法测试
func TestTrieMatch(t *testing.T) {
	root := &node{}

	// 插入示例节点
	root.children = append(root.children, &node{part: "bar"})
	root.children = append(root.children, &node{part: "baz"})
	root.children = append(root.children, &node{part: ":wild", isWild: true})

	// 使用 matchChild 方法
	matchedChild := root.matchChild("bar")
	if matchedChild != nil {
		fmt.Printf("matchChild('bar') returned: %s\n", matchedChild.part)
	} else {
		t.Error("matchChild('bar') returned nil, expected 'bar'")
	}

	matchedChildWild := root.matchChild("qux")
	if matchedChildWild != nil {
		fmt.Printf("matchChild('qux') returned: %s\n", matchedChildWild.part)
	} else {
		t.Error("matchChild('qux') returned nil, expected 'qux'")
	}

	// 使用 matchChildren 方法
	matchedChildren := root.matchChildren("bar")
	fmt.Printf("matchChildren('bar') returned: ")
	for _, child := range matchedChildren {
		fmt.Printf("%s ", child.part)
	}
	fmt.Println()

	matchedChildrenWild := root.matchChildren("qux")
	fmt.Printf("matchChildren('qux') returned: ")
	for _, child := range matchedChildrenWild {
		fmt.Printf("%s ", child.part)
	}
	fmt.Println()
}

// search方法测试
func TestTrieSearch(t *testing.T) {
	root := &node{}

	// 插入一些路由模式
	root.insert("/users/:id", []string{"users", ":id"}, 0)
	root.insert("/articles/*filepath", []string{"articles", "*filepath"}, 0)
	root.insert("/static/*filepath", []string{"static", "*filepath"}, 0)

	// 测试搜索功能
	tests := []struct {
		path     string
		expected string
	}{
		{"/users/123", "/users/:id"},
		{"/articles/some-article", "/articles/*filepath"},
		{"/static/js/app.js", "/static/*filepath"},
		{"/unknown/path", ""},
	}

	for _, test := range tests {
		node := root.search(strings.Split(test.path, "/")[1:], 0)
		if node == nil || node.pattern != test.expected {
			t.Errorf("Path %s: Expected %s, but got %s", test.path, test.expected, node.pattern)
		} else {
			t.Log("ok")
		}
	}
}
