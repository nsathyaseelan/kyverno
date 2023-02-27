package policy

import (
	"fmt"

	kyvernov1 "github.com/nsathyaseelan/kyverno/api/kyverno/v1"
	"github.com/nsathyaseelan/kyverno/pkg/clients/dclient"
	"github.com/nsathyaseelan/kyverno/pkg/logging"
	"github.com/nsathyaseelan/kyverno/pkg/policy/generate"
	"github.com/nsathyaseelan/kyverno/pkg/policy/mutate"
	"github.com/nsathyaseelan/kyverno/pkg/policy/validate"
	"golang.org/x/exp/slices"
)

// Validation provides methods to validate a rule
type Validation interface {
	Validate() (string, error)
}

// validateAction performs validation on the rule actions
// - Mutate
// - Validation
// - Generate
func validateActions(idx int, rule *kyvernov1.Rule, client dclient.Interface, mock bool) error {
	if rule == nil {
		return nil
	}

	var checker Validation

	// Mutate
	if rule.HasMutate() {
		checker = mutate.NewMutateFactory(rule.Mutation)
		if path, err := checker.Validate(); err != nil {
			return fmt.Errorf("path: spec.rules[%d].mutate.%s.: %v", idx, path, err)
		}
	}

	// Validate
	if rule.HasValidate() {
		checker = validate.NewValidateFactory(&rule.Validation)
		if path, err := checker.Validate(); err != nil {
			return fmt.Errorf("path: spec.rules[%d].validate.%s.: %v", idx, path, err)
		}
	}

	// Generate
	if rule.HasGenerate() {
		// TODO: this check is there to support offline validations
		// generate uses selfSubjectReviews to verify actions
		// this need to modified to use different implementation for online and offline mode
		if mock {
			checker = generate.NewFakeGenerate(rule.Generation)
			if path, err := checker.Validate(); err != nil {
				return fmt.Errorf("path: spec.rules[%d].generate.%s.: %v", idx, path, err)
			}
		} else {
			checker = generate.NewGenerateFactory(client, rule.Generation, logging.GlobalLogger())
			if path, err := checker.Validate(); err != nil {
				return fmt.Errorf("path: spec.rules[%d].generate.%s.: %v", idx, path, err)
			}
		}

		if slices.Contains(rule.MatchResources.Kinds, rule.Generation.Kind) {
			return fmt.Errorf("generation kind and match resource kind should not be the same")
		}
	}

	return nil
}
