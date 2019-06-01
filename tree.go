package zin

type methodTree struct {
	method string
	root   *treeNode
}

type methodTrees map[string]methodTree

type treeNode struct {
	path     string
	child    []*treeNode
	handlers HandlerChain
	indices  string
}

func (tree methodTree) getValue(path string) (HandlerChain, Params) {
	return nil, nil
}

func (tree methodTree) addRoute(path string, handlers HandlerChain) {
	node := tree.root
	if node == nil {
		node := new(treeNode)
		node.path = "/"
		node.child = make([]*treeNode, 0)
		node.handlers = make(HandlerChain, 0)
		tree.root = node
	}
walk:
	for {
		if len(path) < len(node.path) {
			var end int
			for end = 0; end < len(path); end++ {
				if path[end] != node.path[end] {
					break
				}
			}
			path = path[end:]

		}
		continue walk
	}
}
