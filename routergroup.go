package zhanst

import "net/http"

type IRoutes interface {
	Use(...HandlerFunc)

	Any(string, ...HandlerFunc)
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	OPTIONS(string, ...HandlerFunc)
	HEAD(string, ...HandlerFunc)

	StaticFile(string, string)
	Static(string, string)
	StaticFS(string, http.FileSystem)
}

type RouterGroup struct {
	Handlers HandlerChain
	basePath string
	engine   *Engine
	root     bool
}

func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
}

func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: append(group.Handlers, handlers...),
		basePath: group.calculateAbsolutePath(relativePath),
		engine:   group.engine,
	}
}

func (group *RouterGroup) GET(path string, handlers ...HandlerFunc) {
	group.handle("GET", path, handlers...)
}

func (group *RouterGroup) POST(path string, handlers ...HandlerFunc) {
	group.handle("POST", path, handlers...)
}

func (group *RouterGroup) PUT(path string, handlers ...HandlerFunc) {
	group.handle("PUT", path, handlers...)
}

func (group *RouterGroup) DELETE(path string, handlers ...HandlerFunc) {
	group.handle("DELETE", path, handlers...)
}

func (group *RouterGroup) PATCH(path string, handlers ...HandlerFunc) {
	group.handle("PATCH", path, handlers...)
}

func (group *RouterGroup) OPTIONS(path string, handlers ...HandlerFunc) {
	group.handle("OPTIONS", path, handlers...)
}

func (group *RouterGroup) HEAD(path string, handlers ...HandlerFunc) {
	group.handle("HEAD", path, handlers...)
}

func (group *RouterGroup) ANY(path string, handlers ...HandlerFunc) {
	group.handle("HEAD", path, handlers...)
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	absolutePath := group.calculateAbsolutePath(relativePath)
	combinedHandlers := group.combineHandlers(handlers...)
	group.engine.addRoute(httpMethod, absolutePath, combinedHandlers)
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPath(group.basePath, relativePath)
}

func (group *RouterGroup) combineHandlers(handlers ...HandlerFunc) HandlerChain {
	combinedHandlers := make(HandlerChain, len(group.Handlers)+len(handlers))
	copy(combinedHandlers, group.Handlers)
	copy(combinedHandlers[len(group.Handlers):], handlers)
	return combinedHandlers
}
