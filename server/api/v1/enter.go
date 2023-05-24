package v1

import "code_memory/api/v1/code_memory"

type ApiGroup struct {
	ArticleApiGroup code_memory.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
