// package apis contains API definitions for the nine public API.
package apis

import (
	"strings"

	"github.com/gobuffalo/flect"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apps "github.com/ninech/apis/apps/v1alpha1"
	infrastructure "github.com/ninech/apis/infrastructure/v1alpha1"
	devtools "github.com/ninech/apis/devtools/v1alpha1"
	iam "github.com/ninech/apis/iam/v1alpha1"
	observability "github.com/ninech/apis/observability/v1alpha1"
	networking "github.com/ninech/apis/networking/v1alpha1"
	storage "github.com/ninech/apis/storage/v1alpha1"
	management "github.com/ninech/apis/management/v1alpha1"
	security "github.com/ninech/apis/security/v1alpha1"
)

func init() {
	AddToSchemes = append(AddToSchemes,
		apps.SchemeBuilder.AddToScheme,
		infrastructure.SchemeBuilder.AddToScheme,
		devtools.SchemeBuilder.AddToScheme,
		iam.SchemeBuilder.AddToScheme,
		observability.SchemeBuilder.AddToScheme,
		networking.SchemeBuilder.AddToScheme,
		storage.SchemeBuilder.AddToScheme,
		management.SchemeBuilder.AddToScheme,
		security.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}

// StaticRESTMapper returns a RESTMapper with all known types statically
// defined to avoid slow discovery.
func StaticRESTMapper(scheme *runtime.Scheme) *meta.DefaultRESTMapper {
	mapper := meta.NewDefaultRESTMapper(scheme.PrioritizedVersionsAllGroups())
	AddToMapper(mapper, apps.ApplicationGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, apps.BuildGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, apps.ProjectConfigGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, apps.ReleaseGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, infrastructure.CloudVirtualMachineGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, infrastructure.ClusterDataGroupVersionKind, meta.RESTScopeRoot)
	AddToMapper(mapper, infrastructure.KedaGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, infrastructure.KubernetesClusterGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, devtools.ArgoCDGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, iam.APIServiceAccountGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, iam.KubernetesClustersRoleBindingGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, iam.KubernetesServiceAccountGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, observability.AlertmanagerGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, observability.GrafanaGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, observability.LokiGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, observability.MetricsAgentGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, observability.PrometheusGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, observability.PromtailGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, observability.TempoGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, observability.TracingCollectorGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, networking.IngressNginxGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, networking.ServiceConnectionGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, networking.StaticEgressGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.BucketGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.BucketUserGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.BucketMigrationGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.DatabaseBackupGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.DatabaseBackupScheduleGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.KeyValueStoreGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.MySQLGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.MySQLDatabaseGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.ObjectsBucketGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.OpenSearchGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.PostgresGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.PostgresDatabaseGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, storage.RegistryGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, management.ProjectGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, security.ExternalSecretsGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, security.SealedSecretsGroupVersionKind, meta.RESTScopeNamespace)
	AddToMapper(mapper, security.SSHKeyGroupVersionKind, meta.RESTScopeNamespace)
	return mapper
}

// AddToMapper is an alternative implementation of DefaultRESTMapper.Add which
// uses flect.pluralize as it is more robust and also used by controller-gen.
func AddToMapper(mapper *meta.DefaultRESTMapper, gvk schema.GroupVersionKind, scope meta.RESTScope) {
	if len(gvk.Kind) == 0 {
		return
	}

	singularName := strings.ToLower(gvk.Kind)
	singular := gvk.GroupVersion().WithResource(singularName)
	plural := gvk.GroupVersion().WithResource(flect.Pluralize(singularName))
	mapper.AddSpecific(gvk, plural, singular, scope)
}
