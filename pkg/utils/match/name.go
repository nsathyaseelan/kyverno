package match

import (
	"github.com/nsathyaseelan/kyverno/pkg/utils/wildcard"
)

func CheckName(expected, actual string) bool {
	return wildcard.Match(expected, actual)
}
