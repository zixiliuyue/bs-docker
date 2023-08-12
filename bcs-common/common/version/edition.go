

package version

const (
	// PlatformName xxx
	PlatformName = "bcs"
)

var (
	// BcsEdition will be injected during build time.
	BcsEdition = InnerEdition
)

const (
	// InnerEdition xxx
	InnerEdition = "inner_edition"
	// CommunicationEdition xxx
	CommunicationEdition = "communication_edition"
	// EnterpriseEdition xxx
	EnterpriseEdition = "enterprise_edition"
)

// GetEdition xxx
func GetEdition() string {
	return BcsEdition
}
