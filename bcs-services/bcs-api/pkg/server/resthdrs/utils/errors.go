

package utils

// ClusterAreadyExistError xxx
type ClusterAreadyExistError struct {
	message string
}

// Error 用于错误处理
func (e *ClusterAreadyExistError) Error() string {
	return e.message
}

// NewClusterAreadyExistError xxx
func NewClusterAreadyExistError(message string) *ClusterAreadyExistError {
	return &ClusterAreadyExistError{
		message: message,
	}
}

// CannotCreateClusterError xxx
type CannotCreateClusterError struct {
	message string
}

// Error 用于错误处理
func (e *CannotCreateClusterError) Error() string {
	return e.message
}

// NewCannotCreateClusterError xxx
func NewCannotCreateClusterError(message string) *CannotCreateClusterError {
	return &CannotCreateClusterError{
		message: message,
	}
}

// ClusterInitFailedError xxx
type ClusterInitFailedError struct {
	message string
}

// Error 用于错误处理
func (e *ClusterInitFailedError) Error() string {
	return e.message
}

// NewClusterInitFailedError xxx
func NewClusterInitFailedError(message string) *ClusterInitFailedError {
	return &ClusterInitFailedError{
		message: message,
	}
}
