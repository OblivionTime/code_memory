package router

import "code_memory/router/code_memory"

type RouterGroup struct {
	System code_memory.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
