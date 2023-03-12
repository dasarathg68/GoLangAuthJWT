package middleware

import (
	"fmt"
	helper "github.com/dasarathg68/GoLangAuthJWT/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middleware entered")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Authorization is null in Header",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Format of Authorization is wrong",
			})
			c.Abort()
			return
		}
		// parts[0] is Bearer, parts is token.
		//mc, err := ParseToken(parts[1])
		//if err != nil {
		//	c.JSON(http.StatusUnauthorized, gin.H{
		//		"code": -1,
		//		"msg":  "Invalid Token.",
		//	})
		//	c.Abort()
		//	return
		//}
		claims, err := helper.ValidateToken(parts[1])
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		//clientToken := c.Request.Header.Get("authfunc")
		//fmt.Println(clientToken)
		//if clientToken == "" {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
		//	c.Abort()
		//	return
		//}
		//

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}
