package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "apps.nine.ch"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

var (
	ApplicationKind             = reflect.TypeOf(Application{}).Name()
	ApplicationGroupKind        = schema.GroupKind{Group: Group, Kind: ApplicationKind}.String()
	ApplicationKindAPIVersion   = ApplicationKind + "." + SchemeGroupVersion.String()
	ApplicationGroupVersionKind = SchemeGroupVersion.WithKind(ApplicationKind)
	BuildKind             = reflect.TypeOf(Build{}).Name()
	BuildGroupKind        = schema.GroupKind{Group: Group, Kind: BuildKind}.String()
	BuildKindAPIVersion   = BuildKind + "." + SchemeGroupVersion.String()
	BuildGroupVersionKind = SchemeGroupVersion.WithKind(BuildKind)
	ProjectConfigKind             = reflect.TypeOf(ProjectConfig{}).Name()
	ProjectConfigGroupKind        = schema.GroupKind{Group: Group, Kind: ProjectConfigKind}.String()
	ProjectConfigKindAPIVersion   = ProjectConfigKind + "." + SchemeGroupVersion.String()
	ProjectConfigGroupVersionKind = SchemeGroupVersion.WithKind(ProjectConfigKind)
	ReleaseKind             = reflect.TypeOf(Release{}).Name()
	ReleaseGroupKind        = schema.GroupKind{Group: Group, Kind: ReleaseKind}.String()
	ReleaseKindAPIVersion   = ReleaseKind + "." + SchemeGroupVersion.String()
	ReleaseGroupVersionKind = SchemeGroupVersion.WithKind(ReleaseKind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
	SchemeBuilder.Register(&Build{}, &BuildList{})
	SchemeBuilder.Register(&ProjectConfig{}, &ProjectConfigList{})
	SchemeBuilder.Register(&Release{}, &ReleaseList{})
}
