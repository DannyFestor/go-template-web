package helpers

import (
	"net/http"

	"github.com/DannyFestor/go-template-web.git/cmd/config"
)

func SetRenderBlock(r *http.Request, value any) *http.Request {
	return SetRequestContext(r, config.RenderBlock, value)
}
