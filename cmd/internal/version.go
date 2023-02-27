package internal

import (
	"github.com/go-logr/logr"
	"github.com/nsathyaseelan/kyverno/pkg/version"
)

func ShowVersion(logger logr.Logger) {
	logger = logger.WithName("version")
	version.PrintVersionInfo(logger)
}
