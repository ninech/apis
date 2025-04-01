package v1alpha1

import (
	"encoding/json"
)

// NewMachineType returns an unparsed machine type that can be used to
// create/update objects. Parsing and validation is done server-side.
func NewMachineType(name string) MachineType {
	return MachineType{name: name}
}

func (mt MachineType) String() string {
	return mt.name
}

// MarshalJSON implements the json.Marshaler interface.
func (mt MachineType) MarshalJSON() ([]byte, error) {
	if mt.name == "" {
		// encode unset name as JSON's "null". This is very important to make
		// defaulting work when creating resources from Go.
		return []byte("null"), nil
	}
	return json.Marshal(mt.name)
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (mt *MachineType) UnmarshalJSON(b []byte) error {
	var name string
	if err := json.Unmarshal(b, &name); err != nil {
		return err
	}
	mt.name = name
	return nil
}

func (mt MachineType) Equal(other MachineType) bool {
	return mt.name == other.name
}
