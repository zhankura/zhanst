package zhanst

import (
	"errors"
	"strings"
)

const (
	part  = 1
	param = 2
)

type methodTree struct {
	method string
	root   *treeNode
}

type methodTrees map[string]methodTree

type treeNode struct {
	path     string
	children []*treeNode
	nodeType uint
	handlers HandlerChain
	indices  string
}

func min(first, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}

func (tree *methodTree) getValue(path string) (HandlerChain, Params) {
	params := make(Params, 0)
	node := tree.root
walk:
	for {
		if path == node.path {
			return node.handlers, params
		}
		maxLength := min(len(path), len(node.path))
		var end int
		for end = 0; end < maxLength; end++ {
			if path[end] != node.path[end] {
				break
			}
		}
		if end != len(node.path) {
			panic(errors.New("not found"))
		}
		path = path[end:]
		if id := strings.Index(node.indices, string(path[0])); id != -1 {
			node = node.children[id]
			continue walk
		} else {
			panic(errors.New("not found"))
		}
	}
}

func (tree *methodTree) addRoute(path string, handlers HandlerChain) {
	node := tree.root
walk:
	for {
		maxLength := min(len(path), len(node.path))
		var end int
		for end = 0; end < maxLength; end++ {
			if path[end] != node.path[end] {
				break
			}
		}
		if end != maxLength {
			newNode := new(treeNode)
			newNode.path = node.path[end:]
			newNode.handlers = node.handlers
			newNode.children = node.children
			newNode.indices = node.indices
			node.path = node.path[:end]
			node.handlers = nil
			insertNode := new(treeNode)
			insertNode.path = path[end:]
			insertNode.handlers = handlers
			node.children = []*treeNode{newNode, insertNode}
			node.indices = string([]uint8{newNode.path[0], insertNode.path[0]})
			return
		} else {
			var remainPath string
			if len(path) > len(node.path) {
				remainPath = path[end:]
				if id := strings.Index(node.indices, string(remainPath[0])); id != -1 {
					node = node.children[id]
					path = remainPath
					continue walk
				}
			} else if len(path) < len(node.path) {
				remainPath = node.path[end:]
			} else {
				panic(errors.New("url has been used"))
			}
			insertNode := new(treeNode)
			insertNode.path = remainPath
			insertNode.handlers = handlers
			node.children = append(node.children, insertNode)
			node.indices = node.indices + string(remainPath[0])
			return
		}
	}
}

func (tree *methodTree) AddRoute(path string, handlers HandlerChain) {
	tree.addRoute(path, handlers)
}

func (tree *methodTree) GetValue(path string) (HandlerChain, Params) {
	return tree.getValue(path)
}
