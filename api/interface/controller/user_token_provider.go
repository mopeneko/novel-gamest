package controller

// UserTokenProvider provides for generating and verifing tokens for users
type UserTokenProvider interface {
	Generate(id string) (string, error)
}
