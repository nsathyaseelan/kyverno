package utils

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	kyvernov1 "github.com/nsathyaseelan/kyverno/api/kyverno/v1"
	engineapi "github.com/nsathyaseelan/kyverno/pkg/engine/api"
	"github.com/nsathyaseelan/kyverno/pkg/metrics"
	policyExecutionDuration "github.com/nsathyaseelan/kyverno/pkg/metrics/policyexecutionduration"
	policyResults "github.com/nsathyaseelan/kyverno/pkg/metrics/policyresults"
)

type reporterFunc func(metrics.ResourceRequestOperation) error

func registerMetric(logger logr.Logger, m string, requestOperation string, r reporterFunc) {
	if op, err := metrics.ParseResourceRequestOperation(requestOperation); err != nil {
		logger.Error(err, fmt.Sprintf("error occurred while registering %s metrics", m))
	} else {
		if err := r(op); err != nil {
			logger.Error(err, fmt.Sprintf("error occurred while registering %s metrics", m))
		}
	}
}

// POLICY RESULTS

func RegisterPolicyResultsMetricMutation(ctx context.Context, logger logr.Logger, metricsConfig metrics.MetricsConfigManager, requestOperation string, policy kyvernov1.PolicyInterface, engineResponse engineapi.EngineResponse) {
	registerMetric(logger, "kyverno_policy_results_total", requestOperation, func(op metrics.ResourceRequestOperation) error {
		return policyResults.ProcessEngineResponse(ctx, metricsConfig, policy, engineResponse, metrics.AdmissionRequest, op)
	})
}

func RegisterPolicyResultsMetricValidation(ctx context.Context, logger logr.Logger, metricsConfig metrics.MetricsConfigManager, requestOperation string, policy kyvernov1.PolicyInterface, engineResponse engineapi.EngineResponse) {
	registerMetric(logger, "kyverno_policy_results_total", requestOperation, func(op metrics.ResourceRequestOperation) error {
		return policyResults.ProcessEngineResponse(ctx, metricsConfig, policy, engineResponse, metrics.AdmissionRequest, op)
	})
}

func RegisterPolicyResultsMetricGeneration(ctx context.Context, logger logr.Logger, metricsConfig metrics.MetricsConfigManager, requestOperation string, policy kyvernov1.PolicyInterface, engineResponse engineapi.EngineResponse) {
	registerMetric(logger, "kyverno_policy_results_total", requestOperation, func(op metrics.ResourceRequestOperation) error {
		return policyResults.ProcessEngineResponse(ctx, metricsConfig, policy, engineResponse, metrics.AdmissionRequest, op)
	})
}

// POLICY EXECUTION

func RegisterPolicyExecutionDurationMetricMutate(ctx context.Context, logger logr.Logger, metricsConfig metrics.MetricsConfigManager, requestOperation string, policy kyvernov1.PolicyInterface, engineResponse engineapi.EngineResponse) {
	registerMetric(logger, "kyverno_policy_execution_duration_seconds", requestOperation, func(op metrics.ResourceRequestOperation) error {
		return policyExecutionDuration.ProcessEngineResponse(ctx, metricsConfig, policy, engineResponse, metrics.AdmissionRequest, op)
	})
}

func RegisterPolicyExecutionDurationMetricValidate(ctx context.Context, logger logr.Logger, metricsConfig metrics.MetricsConfigManager, requestOperation string, policy kyvernov1.PolicyInterface, engineResponse engineapi.EngineResponse) {
	registerMetric(logger, "kyverno_policy_execution_duration_seconds", requestOperation, func(op metrics.ResourceRequestOperation) error {
		return policyExecutionDuration.ProcessEngineResponse(ctx, metricsConfig, policy, engineResponse, metrics.AdmissionRequest, op)
	})
}

func RegisterPolicyExecutionDurationMetricGenerate(ctx context.Context, logger logr.Logger, metricsConfig metrics.MetricsConfigManager, requestOperation string, policy kyvernov1.PolicyInterface, engineResponse engineapi.EngineResponse) {
	registerMetric(logger, "kyverno_policy_execution_duration_seconds", requestOperation, func(op metrics.ResourceRequestOperation) error {
		return policyExecutionDuration.ProcessEngineResponse(ctx, metricsConfig, policy, engineResponse, metrics.AdmissionRequest, op)
	})
}
