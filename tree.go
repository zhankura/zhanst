package zin

type methodTree struct {
	method string
	root *node
}

type methodTrees map[string]methodTree

type node struct {
	path string
	child []*node
	handlers HandlerChain
}


func (tree methodTree) getValue(path string) (HandlerChain, Params){
	return nil, nil
}

func (tree methodTree) addRoute(path string, handlers HandlerChain) {

}

