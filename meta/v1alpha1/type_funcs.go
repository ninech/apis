package v1alpha1

func (t TypedReference) String() string {
	return fmt.Sprintf("%s/%s (%s)", t.Namespace, t.Name, t.GroupKind.String())
}
