

package argocd

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/pkg/proxy"
)

// WebhookPlugin defines the webhook plugin
type WebhookPlugin struct {
	*mux.Router
	middleware MiddlewareInterface
}

// Init initialize webhook plugin
func (plugin *WebhookPlugin) Init() error {
	plugin.Path("").Methods("POST").
		Handler(plugin.middleware.HttpWrapper(plugin.executeWebhook))

	blog.Infof("argocd webhook plugin init successfully")
	return nil
}

func (plugin *WebhookPlugin) executeWebhook(ctx context.Context, r *http.Request) *httpResponse {
	user := ctx.Value(ctxKeyUser).(*proxy.UserInfo)
	blog.Infof("user %s request webhook", user.GetUser())
	return nil
}
