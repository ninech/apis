// package apis contains API definitions for the nine public API.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"
	devtools "github.com/ninech/apis/devtools/v1alpha1"
	iam "github.com/ninech/apis/iam/v1alpha1"
	infrastructure "github.com/ninech/apis/infrastructure/v1alpha1"
	networking "github.com/ninech/apis/networking/v1alpha1"
	storage "github.com/ninech/apis/storage/v1alpha1"
	observability "github.com/ninech/apis/observability/v1alpha1"
	security "github.com/ninech/apis/security/v1alpha1"
)

func init() {
	AddToSchemes = append(AddToSchemes,
		devtools.SchemeBuilder.AddToScheme,
		iam.SchemeBuilder.AddToScheme,
		infrastructure.SchemeBuilder.AddToScheme,
		networking.SchemeBuilder.AddToScheme,
		storage.SchemeBuilder.AddToScheme,
		observability.SchemeBuilder.AddToScheme,
		security.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
