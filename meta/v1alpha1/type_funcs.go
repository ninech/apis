package v1alpha1

import (
	"fmt"
)

func (t TypedReference) String() string {
	return fmt.Sprintf("%s/%s (%s)", t.Namespace, t.Name, t.GroupKind.String())
}

// String implements the stringer interface.
func (d DNSCheckType) String() string {
	return string(d)
}

// CheckTypeEntry returns the corresponding checkType entry for the given host. The
// second return parameter indicates if an entry was found. There should always
// just be one entry per checkType in a DNSVerificationStatusEntries.
func (vsl DNSVerificationStatusEntries) CheckTypeEntry(host string, checkType DNSCheckType) (DNSVerificationStatusEntry, bool) {
	for _, s := range vsl {
		if s.Name == host && s.CheckType == checkType {
			return s, true
		}
	}
	return DNSVerificationStatusEntry{}, false
}

// Verified returns if the host of the VerificationStatus was successfully verified.
func (vs DNSVerificationStatusEntry) Verified() bool {
	return (vs.CheckType == DNSCheckTXT || vs.CheckType == DNSCheckCNAME) && vs.LatestSuccess != nil
}
