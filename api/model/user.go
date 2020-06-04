package model

// UserCreateRequest is a request for creating an user
type UserCreateRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// UserCreateResponse is a response for creating an user
type UserCreateResponse struct {
	Message string `json:"message"`
}

// UserGetByIDRequest is a request for getting an user which has same ID
type UserGetByIDRequest struct {
	ID string `json:"id"`
}

// UserGetByIDResponse is a response for getting an user which has same ID
type UserGetByIDResponse struct {
	Name  string `json:"name"`
	Posts []Game `json:"posts"`
}

// UserGetWithAuthenticationRequest is a request for getting an user which has same authentication
type UserGetWithAuthenticationRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// UserGetWithAuthenticationResponse is a response for getting an user which has same authentication
type UserGetWithAuthenticationResponse struct {
	Message string `json:"message"`
}
