package middleware

import (
	"encoding/json"
	"fmt"
	. "github.com/davdwhyte87/travelfy/utils"
	"github.com/dgrijalva/jwt-go"
	"net/http/httputil"
	"net/http"
)

// AuthenticationMiddleware ... This middle ware validates a token for protected routes
func AuthenticationMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader :=r.Header.Get("Authorization")
		if authorizationHeader == "" {
			RespondWithError(w, http.StatusUnauthorized, "You are not authorized")
			return
		}
		token, err :=jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return  []byte(SecreteKey), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			type Person struct {
				Id int
				Name string
				Age int
			}

			var p = Person{Id:9, Name:"sjs", Age:2}
			b, _ := json.Marshal(p)
			httputil.DumpRequest(r,)
			r.Write(b)

			nextHandler.ServeHTTP(w, r)
		} else {
			fmt.Println(err)
			//RespondWithError(w, http.StatusUnauthorized, "An authorized error occurred")
		}
	})
}

