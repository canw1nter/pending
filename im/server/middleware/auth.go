package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"pending/im/server/common"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Token")
		if tokenStr == "" {
			log.Printf("Recieve a unauthoried request from ip %s!\n", c.ClientIP())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenStr, &common.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return common.TokenKey, nil
		})
		if err != nil {
			log.Printf("Parse token from ip %s failed! err: %s\n %t", c.ClientIP(), err.Error(), err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			log.Printf("Token is invalid from ip %s!\n", c.ClientIP())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*common.TokenClaims)
		if !ok {
			log.Printf("Get claims failed from token, request ip %s!\n", c.ClientIP())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		expirationTime, err := claims.GetExpirationTime()
		if err != nil {
			log.Printf(
				"Can't get expiration time from token claims from ip %s! err: %s\n",
				c.ClientIP(), err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if time.Now().Add(time.Minute * 30).After(expirationTime.Time) {
			newToken, err := common.GenerateUserToken(claims.UUID, claims.Name)
			if err != nil {
				log.Printf("Generate new token failed! Request from ip %s!\n", c.ClientIP())
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Header("New-Token", newToken)
		}

		c.Set("uuid", claims.UUID)
		c.Set("name", claims.Name)
		c.Next()
	}
}
