package v1alpha1

import (
	"fmt"
)

func (t TypedReference) String() string {
	return fmt.Sprintf("%s/%s (%s)", t.Namespace, t.Name, t.GroupKind.String())
}
