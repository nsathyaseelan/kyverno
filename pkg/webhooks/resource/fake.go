package resource

import (
	"context"

	fakekyvernov1 "github.com/nsathyaseelan/kyverno/pkg/client/clientset/versioned/fake"
	kyvernoinformers "github.com/nsathyaseelan/kyverno/pkg/client/informers/externalversions"
	"github.com/nsathyaseelan/kyverno/pkg/clients/dclient"
	"github.com/nsathyaseelan/kyverno/pkg/config"
	"github.com/nsathyaseelan/kyverno/pkg/engine"
	engineapi "github.com/nsathyaseelan/kyverno/pkg/engine/api"
	"github.com/nsathyaseelan/kyverno/pkg/engine/context/resolvers"
	"github.com/nsathyaseelan/kyverno/pkg/event"
	"github.com/nsathyaseelan/kyverno/pkg/metrics"
	"github.com/nsathyaseelan/kyverno/pkg/openapi"
	"github.com/nsathyaseelan/kyverno/pkg/policycache"
	"github.com/nsathyaseelan/kyverno/pkg/registryclient"
	"github.com/nsathyaseelan/kyverno/pkg/webhooks"
	"github.com/nsathyaseelan/kyverno/pkg/webhooks/updaterequest"
	webhookutils "github.com/nsathyaseelan/kyverno/pkg/webhooks/utils"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
)

func NewFakeHandlers(ctx context.Context, policyCache policycache.Cache) webhooks.ResourceHandlers {
	client := fake.NewSimpleClientset()
	metricsConfig := metrics.NewFakeMetricsConfig()

	informers := kubeinformers.NewSharedInformerFactory(client, 0)
	informers.Start(ctx.Done())

	kyvernoclient := fakekyvernov1.NewSimpleClientset()
	kyvernoInformers := kyvernoinformers.NewSharedInformerFactory(kyvernoclient, 0)
	configMapResolver, _ := resolvers.NewClientBasedResolver(client)
	kyvernoInformers.Start(ctx.Done())

	dclient := dclient.NewEmptyFakeClient()
	configuration := config.NewDefaultConfiguration()
	rbLister := informers.Rbac().V1().RoleBindings().Lister()
	crbLister := informers.Rbac().V1().ClusterRoleBindings().Lister()
	urLister := kyvernoInformers.Kyverno().V1beta1().UpdateRequests().Lister().UpdateRequests(config.KyvernoNamespace())
	peLister := kyvernoInformers.Kyverno().V2alpha1().PolicyExceptions().Lister()
	rclient := registryclient.NewOrDie()

	return &handlers{
		client:         dclient,
		rclient:        rclient,
		configuration:  configuration,
		metricsConfig:  metricsConfig,
		pCache:         policyCache,
		nsLister:       informers.Core().V1().Namespaces().Lister(),
		rbLister:       rbLister,
		crbLister:      crbLister,
		urLister:       urLister,
		urGenerator:    updaterequest.NewFake(),
		eventGen:       event.NewFake(),
		openApiManager: openapi.NewFake(),
		pcBuilder:      webhookutils.NewPolicyContextBuilder(configuration, dclient, rbLister, crbLister),
		urUpdater:      webhookutils.NewUpdateRequestUpdater(kyvernoclient, urLister),
		engine: engine.NewEngine(
			configuration,
			dclient,
			rclient,
			engineapi.DefaultContextLoaderFactory(configMapResolver),
			peLister,
		),
	}
}
