package jin

import (
	"jin/log"
	"net/http"
)

// HandlerFunc is the handler of HTTP request
type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix string       // The prefix of the route
		parent *RouterGroup // support router group nesting
		engine *Engine      // all the groups share one engine instance
	}

	// Engine implements ServeHTTP
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// create a new group
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodGet, pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodPost, pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	log.Infof("Listening and serving HTTP on http://127.0.0.1%s/", addr)
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
