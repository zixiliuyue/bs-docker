

package authcenter

// UsersQueryResponse all user information in AuthCenter relative to PaaS project
type UsersQueryResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}
