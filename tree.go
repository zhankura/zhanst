package zhanst

import (
	"errors"
	"strings"
)

const (
	part  = 0
	param = 1
)

type methodTree struct {
	method string
	root   *treeNode
}

type methodTrees map[string]methodTree

type treeNode struct {
	path     string
	children []*treeNode
	nodeType int
	handlers HandlerChain
	indices  string
}

type nodeMid struct {
	path string
	node *treeNode
}

func min(first, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}

func (tree *methodTree) getValue(path string) (HandlerChain, Params) {
	params := make(Params)
	node := tree.root
	nodes := make([]nodeMid, 0)
walk:
	for {
		if path == node.path && node.nodeType == part {
			return node.handlers, params
		}
		if id := strings.Index(path, "/"); id == -1 && node.nodeType == param {
			if _, ok := params[node.path]; !ok {
				params[node.path] = path
			} else {
				panic(errors.New("param name can not be same"))
			}
			return node.handlers, params
		}
		var end int
		if node.nodeType == part {
			maxLength := min(len(path), len(node.path))
			for end = 0; end < maxLength; end++ {
				if path[end] != node.path[end] {
					break
				}
			}
			if end != len(node.path) {
				panic(errors.New("not found"))
			}
		} else {
			for end = 0; end < len(path); end++ {
				if path[end] == '/' {
					break
				}
			}
			if _, ok := params[node.path]; !ok {
				params[node.path] = path
			}
		}
		path = path[end:]
		if strings.Contains(node.indices, ":") && strings.Contains(node.indices, string(path[0])) {
			nodes = append(nodes, nodeMid{
				path: path,
				node: node,
			})
		}
		if id := strings.Index(node.indices, ":"); id != -1 {
			node = node.children[id]
		}
		if id := strings.Index(node.indices, string(path[0])); id != -1 {
			node = node.children[id]
		} else {
			if len(nodes) == 0 {
				panic(errors.New("not found"))
			} else {
				mid := nodes[len(nodes)-1]
				node = mid.node
				path = mid.path
			}
		}
		continue walk
	}
}

func insertChildNode(path string, handlers HandlerChain, children []*treeNode, indices string, newChildTree bool) *treeNode {
	newRoot := new(treeNode)
	if !newChildTree {
		newRoot.path = path
		newRoot.handlers = handlers
		newRoot.children = children
		newRoot.indices = indices
		return newRoot
	}
	fatherParam := false
	var node *treeNode = nil
walk:
	for {
		var begin = 0
		var end = 0
		var position = 0
		var nodeType int
		for position = 0; position < len(path); position++ {
			if path[position] == ':' {
				break
			}
		}
		if position == 0 {
			if fatherParam == true {
				panic(errors.New("param can not be close to"))
			}
			nodeType = param
			fatherParam = true
			begin = position + 1
			for end = begin; end < len(path); end++ {
				if path[end] == ':' {
					panic(errors.New("param error"))
				}
				if path[end] == '/' {
					break
				}
			}
		} else {
			nodeType = part
			end = position
			if path[position-1] != '/' && position != len(path) {
				panic(errors.New(":param must after /"))
			}
			if path[begin:end] != "/" {
				fatherParam = false
			}
		}
		var nowNode *treeNode
		if node != nil {
			newNode := new(treeNode)
			nowNode = newNode
		} else {
			nowNode = newRoot
		}
		nowNode.path = path[begin:end]
		nowNode.nodeType = nodeType
		if node != nil {
			if nodeType == param {
				node.indices = node.indices + ":"
			} else {
				node.indices = node.indices + string(nowNode.path[0])
			}
			node.children = append(node.children, nowNode)
		}
		if end == len(path) {
			nowNode.children = children
			nowNode.handlers = handlers
			return newRoot
		} else {
			nowNode.children = make([]*treeNode, 0)
			nowNode.handlers = make(HandlerChain, 0)
			node = nowNode
			path = path[end:]
		}
		continue walk
	}

}

func getParam(path string) string {
	var begin int
	var end int
	if strings.HasPrefix(path, ":") {
		begin = 1
	}
	for end = 0; end < len(path); end++ {
		if path[end] == '/' {
			break
		}
	}
	return path[begin:end]
}

func (tree *methodTree) addRoute(path string, handlers HandlerChain) {
	node := tree.root
walk:
	for {
		maxMateLength := min(len(path), len(node.path))
		var end int
		for end = 0; end < maxMateLength; end++ {
			if path[end] != node.path[end] {
				break
			}
		}
		if end != maxMateLength {
			newNode := insertChildNode(node.path[end:], node.handlers, node.children, node.indices, false)
			node.path = node.path[:end]
			node.handlers = nil
			remainPath := path[end:]
			insertNode := insertChildNode(remainPath, handlers, make([]*treeNode, 0), "", true)
			node.children = []*treeNode{newNode, insertNode}
			node.indices = string([]uint8{newNode.path[0], remainPath[0]})
			return
		} else {
			var remainPath string
			if len(path) > len(node.path) {
				remainPath = path[end:]
				if id := strings.Index(node.indices, string(remainPath[0])); id != -1 {
					node = node.children[id]
					path = remainPath
					if node.nodeType == param {
						if node.path != getParam(remainPath) {
							panic(errors.New("param has been wrong"))
						}
					}
					continue walk
				}
				insertNode := insertChildNode(remainPath, handlers, make([]*treeNode, 0), "", true)
				node.children = append(node.children, insertNode)
				node.indices = node.indices + string(remainPath[0])
			} else if len(path) < len(node.path) {
				remainPath = node.path[end:]
				insertNode := insertChildNode(remainPath, node.handlers, make([]*treeNode, 0), "", true)
				insertNode.children = node.children
				insertNode.indices = node.indices
				node.handlers = handlers
				node.children = []*treeNode{insertNode}
				node.indices = string(remainPath[0])
			} else {
				panic(errors.New("url has been used"))
			}
			node.path = node.path[:end]
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
