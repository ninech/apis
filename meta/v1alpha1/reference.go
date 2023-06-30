package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type nsGetter interface {
	GetNamespace() string
}

// String implements the Stringer interface for a reference.
func (r Reference) String() string {
	return r.NamespacedName().String()
}

// NamespacedName returns the namespaced name for this reference.
func (r Reference) NamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

// ObjectMeta returns an ObjectMeta of the referenced object.
// Does not populate anything else than Name and Namespace.
func (r Reference) ObjectMeta() metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

// NamespacedName returns the namespaced name of this local refenece with
// the namespace set to the given ns.
func (l LocalReference) NamespacedName(ns nsGetter) types.NamespacedName {
	return types.NamespacedName{
		Name:      l.Name,
		Namespace: ns.GetNamespace(),
	}
}

// InNamespace returns a Reference based on the LocalReference with the
// namespace set to the given ns.
func (l LocalReference) InNamespace(ns nsGetter) Reference {
	return Reference{
		Name:      l.Name,
		Namespace: ns.GetNamespace(),
	}
}
