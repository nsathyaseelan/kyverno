package autogen

import "github.com/nsathyaseelan/kyverno/pkg/logging"

var (
	logger = logging.WithName("autogen")
	debug  = logger.V(5)
)
