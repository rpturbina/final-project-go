package crypto

import (
	"context"
	"os"

	"github.com/kataras/jwt"
	"github.com/rpturbina/final-project-go/pkg/domain/claim"
)

var (
	envSharedKey = os.Getenv("MY_GRAM_SECRET_JWT_SIGNATURE")
	sharedKey    = []byte(envSharedKey)
)

func CreateJWT(ctx context.Context, claim any) (string, error) {

	token, err := jwt.Sign(jwt.HS256, sharedKey, claim)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func VerifyJWT(ctx context.Context, token string) (claims claim.JWTToken, err error) {

	verifiedToken, err := jwt.Verify(jwt.HS256, sharedKey, []byte(token))
	if err != nil {
		return claims, err
	}

	err = verifiedToken.Claims(&claims)
	if err != nil {
		return claims, err
	}
	return claims, err
}
