package zin

import "net/http"

type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}

type RouterGroup struct {
	Handlers HandlerChain
	basePath string
	engine   *Engine
	root     bool
}

func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	group.Handlers = append(group.Handlers, middleware...)
	return group.returnObj()
}

func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.Handlers,
		basePath: group.calculateAbsolutePath(relativePath),
		engine:   group.engine,
	}
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlerChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)

}

func (group *RouterGroup) returnObj() IRoutes {
	if group.root {
		return group.engine
	}
	return group
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPath(group.basePath, relativePath)
}
