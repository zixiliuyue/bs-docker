

package backend

import (
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

// SaveAdmissionWebhook xxx
// save admission webhook
func (b *backend) SaveAdmissionWebhook(admission *commtypes.AdmissionWebhookConfiguration) error {
	return b.store.SaveAdmissionWebhook(admission)
}

// UpdateAdmissionWebhook xxx
func (b *backend) UpdateAdmissionWebhook(admission *commtypes.AdmissionWebhookConfiguration) error {
	return b.store.SaveAdmissionWebhook(admission)
}

// FetchAdmissionWebhook xxx
func (b *backend) FetchAdmissionWebhook(ns, name string) (*commtypes.AdmissionWebhookConfiguration, error) {
	return b.store.FetchAdmissionWebhook(ns, name)
}

// DeleteAdmissionWebhook xxx
func (b *backend) DeleteAdmissionWebhook(ns string, name string) error {
	return b.store.DeleteAdmissionWebhook(ns, name)
}

// FetchAllAdmissionWebhooks xxx
func (b *backend) FetchAllAdmissionWebhooks() ([]*commtypes.AdmissionWebhookConfiguration, error) {
	return b.store.FetchAllAdmissionWebhooks()
}
