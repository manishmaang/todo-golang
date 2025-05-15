package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var jwtSecret = []byte("your-secret-key")

// This setup of function let's us pass extra values to the gin functions without breaking the syntax and making the middleware more flexible
// ultimately the gin.Contezt function will be returned but the parameters which we pass to the main function will let gin function to use those parameters making this setup more flexible
func Authenticate_User() gin.HandlerFunc {
	//c is a gin context which represents the current request and response
	return func(c *gin.Context) {
		fmt.Printf("Inside the Authenticate User middleware in gin/golang")

		authHeader := c.GetHeader("Authorization") // get the token via autorization in headers
		fmt.Println("printing the auth header here : ", authHeader)

		if authHeader == "" {
			//gin.H = map[string]any and .JSON serialize the data being sent gin.H
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Authorization header is required"})
			c.Abort() //It stops further handlers or middleware from running for this request.
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println("Printing the tokenString here : ", tokenString)

		if tokenString == authHeader { // it didn't remove Beare so auth header didn't contain the bearer token
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "No Bearer token is provided in the request"})
			c.Abort()
			return
		}

		//jwt.Parse is decoding and validation the headers and body of the token, also verifies the token via callback function
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			// Ensure signing method is HMAC (where both parties have the same signature to encode the data)
			_, ok := token.Method.(*jwt.SigningMethodHMAC) //verifying wether the signing method is hmac or not
			if !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// token.Claims is a field in the parsed JWT token that holds all the claims (the key-value data inside the token).
		// jwt.MapClaims just make claims map[string]any which will be passed to 
		// 	jwt.MapClaims tells Go: “This is meant to be JWT claims.”
		// gin.H tells Go: “This is a JSON response payload.”
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"]) //request me user_id attach krdi hai
		}

		c.Next()

	}
}

// Also a way to define middleware but it is not reusable and plus doesn't provide flexibility to pass any other input in middleware
// Gin framework syntacx it can only handle the request when we pass the *gin.Context in the function and not anything else
// func MyMiddleware(c *gin.Context) {
// 	// do something
// 	c.Next()
// }
