

package backend

import (
	"fmt"
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

// RegisterCustomResource xxx
// custom resource register
func (b *backend) RegisterCustomResource(crr *commtypes.Crr) error {
	crrs, err := b.store.ListCustomResourceRegister()
	if err != nil {
		return err
	}

	for _, c := range crrs {
		if c.Spec.Names.Kind == crr.Spec.Names.Kind {
			return nil
		}
	}

	return b.store.SaveCustomResourceRegister(crr)
}

// UnregisterCustomResource xxx
func (b *backend) UnregisterCustomResource(string) error {
	// xxx

	return nil
}

// CreateCustomResource xxx
func (b *backend) CreateCustomResource(crd *commtypes.Crd) error {
	crrs, err := b.store.ListCustomResourceRegister()
	if err != nil {
		return err
	}

	var kindOk bool
	for _, c := range crrs {
		if c.Spec.Names.Kind == string(crd.Kind) {
			kindOk = true
		}
	}

	if !kindOk {
		return fmt.Errorf("custom resource kind %s is invalid", crd.Kind)
	}

	return b.store.SaveCustomResourceDefinition(crd)
}

// UpdateCustomResource xxx
func (b *backend) UpdateCustomResource(crd *commtypes.Crd) error {
	crrs, err := b.store.ListCustomResourceRegister()
	if err != nil {
		return err
	}

	var kindOk bool
	for _, c := range crrs {
		if c.Spec.Names.Kind == string(crd.Kind) {
			kindOk = true
		}
	}

	if !kindOk {
		return fmt.Errorf("custom resource kind %s is invalid", crd.Kind)
	}

	return b.store.SaveCustomResourceDefinition(crd)
}

// DeleteCustomResource xxx
func (b *backend) DeleteCustomResource(kind, ns, name string) error {
	return b.store.DeleteCustomResourceDefinition(kind, ns, name)
}

// ListCustomResourceDefinition xxx
func (b *backend) ListCustomResourceDefinition(kind, ns string) ([]*commtypes.Crd, error) {
	return b.store.ListCustomResourceDefinition(kind, ns)
}

// ListAllCrds xxx
func (b *backend) ListAllCrds(kind string) ([]*commtypes.Crd, error) {
	return b.store.ListAllCrds(kind)
}

// FetchCustomResourceDefinition xxx
func (b *backend) FetchCustomResourceDefinition(kind, ns, name string) (*commtypes.Crd, error) {
	return b.store.FetchCustomResourceDefinition(kind, ns, name)
}
