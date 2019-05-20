package zin

type HandlerFunc func(*Context)

type HandlerChain []HandlerFunc

type Engine struct {
	RouterGroup
}
