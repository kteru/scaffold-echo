package request

// Login ...
type Login struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
