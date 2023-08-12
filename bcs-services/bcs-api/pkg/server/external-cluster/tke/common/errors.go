

package common

const (
	// NoErr xxx
	NoErr = 0
	// NoErrCodeDesc xxx
	NoErrCodeDesc = "Success"

	// ErrQCloudGoClient xxx
	ErrQCloudGoClient = 9999
)

// LegacyAPIError xxx
type LegacyAPIError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

// Error 用于错误处理
func (lae LegacyAPIError) Error() string {
	return lae.Message
}

// VersionAPIError xxx
type VersionAPIError struct {
	Response struct {
		Error apiErrorResponse `json:"Error"`
	} `json:"Response"`
}

type apiErrorResponse struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

// Error 用于错误处理
func (vae VersionAPIError) Error() string {
	return vae.Response.Error.Message
}

// ClientError xxx
type ClientError struct {
	Message string
}

// Error 用于错误处理
func (ce ClientError) Error() string {
	return ce.Message
}

func makeClientError(err error) ClientError {
	return ClientError{
		Message: err.Error(),
	}
}
