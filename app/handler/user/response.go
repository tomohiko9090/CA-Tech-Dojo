package handler

// UserCreateResponse 1.1 δ½ζ
type UserCreateResponse struct {
	Token string `json:"token"`
}

// UserGetResponse 1.2 εεΎ
type UserGetResponse struct {
	Name string `json:"name"`
}
