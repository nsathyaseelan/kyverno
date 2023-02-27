package internal

import (
	"context"

	kyvernov1 "github.com/nsathyaseelan/kyverno/api/kyverno/v1"
	engineapi "github.com/nsathyaseelan/kyverno/pkg/engine/api"
)

func LoadContext(
	ctx context.Context,
	engine engineapi.Engine,
	pContext engineapi.PolicyContext,
	rule kyvernov1.Rule,
) error {
	loader := engine.ContextLoader(pContext.Policy(), rule)
	return loader(ctx, rule.Context, pContext.JSONContext())
}
