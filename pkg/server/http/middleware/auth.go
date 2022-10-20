package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/final-project-go/pkg/usecase/crypto"
)

func CheckJWTAuth(ctx *gin.Context) {
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": 97,
			"type": "UNAUTHENTICATED",
		})
		return
	}

	stringToken := strings.Split(headerToken, " ")[1]

	claim := crypto.VerifyJWT(ctx, stringToken)

	// validate claim
	if claim.Issuer != "mygram.com" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": 97,
			"type": "UNAUTHENTICATED",
		})
		return
	}

	if claim.Audience != "user.mygram.com" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": 97,
			"type": "UNAUTHENTICATED",
		})
		return
	}

	if claim.Scope != "create update delete read" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": 97,
			"type": "UNAUTHENTICATED",
		})
		return
	}

	if !time.Unix(claim.NotValidBefore, 0).Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": 97,
			"type": "UNAUTHENTICATED",
		})
		return
	}

	if time.Unix(claim.ExpiredAt, 0).Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": 97,
			"type": "UNAUTHENTICATED",
		})
		return
	}

	ctx.Set("user", claim.Subject)
	ctx.Next()
}
