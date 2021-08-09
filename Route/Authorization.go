package Route

import (
	"../Config"
	"../Entity"
	"errors"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var authorizationMiddlewareFunc = func(env Entity.Env, filters []func(claims jwt.MapClaims) error) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {

			claims := token.Claims.(jwt.MapClaims)

			aud := env.App_Client_ID
			//自分のクライアントIDで署名されていない (必須)
			if claims.VerifyAudience(aud, true) {
				return token, errors.New("invalid audience")
			}

			//// Verify 'iss' claim
			////TODO
			////自分のドメインで署名されていない (必要?)
			//iss := env.App_Domain
			//if !token.Claims.(jwt.MapClaims).VerifyIssuer(iss, true) {
			//	return token, errors.New("invalid issuer")
			//}

			//期限切れのIDトークン(必須)
			if !claims.VerifyNotBefore(time.Now().Unix(), true) {
				return token, errors.New("invalid time")
			}

			var filterError error = nil
			for _, i2 := range filters {
				if tempErr := i2(claims); tempErr != nil {
					filterError = tempErr
					break
				}
			}

			if filterError != nil {
				return token, filterError
			}

			return []byte(env.App_Client_Secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

func AuthFilter(filters []func(claims jwt.MapClaims) error, success func()) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("authorization") == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "empty_id_token",
			})
			return
		}
		jwtMid := *authorizationMiddlewareFunc(Config.Env, filters)
		if err := jwtMid.CheckJWT(ctx.Writer, ctx.Request); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid_id_token",
			})
			return
		}
		success()
	}
}
