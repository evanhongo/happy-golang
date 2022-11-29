package service

type GooglePublicKey struct {
	Kid string `json:"kid"`
	N   string `json:"n"`
}

type IAuthService interface {
	GetGooglePublicKeyChain() ([]GooglePublicKey, error)
	VerifyIdToken(idToken string) error
}
