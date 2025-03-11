package v1alpha1

// NewMachineType returns an unparsed machine type that can be used to
// create/update objects. Parsing and validation is done server-side.
func NewMachineType(name string) MachineType {
	return MachineType{name: name}
}
