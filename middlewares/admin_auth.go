package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/davdwhyte87/travelfy/dao"
	"github.com/davdwhyte87/travelfy/utils"
	"github.com/dgrijalva/jwt-go"
)

// AdminAuthMiddleware ... This middle ware validates a token for protected routes
func AdminAuthMiddleware(nextHandler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "You are not authorized")
			return
		}
		token, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(SecreteKey), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var id string
			id, ok = claims["id"].(string)
			var userDao = dao.UserDAO{}
			user, userErr := userDao.FindById(id)
			if userErr != nil {
				utils.RespondWithError(w, http.StatusUnauthorized, "Error Authenticating ")
				return
			}
			if user.Type != 1 {
				utils.RespondWithError(w, http.StatusUnauthorized, "UnAunthorized")
				return
			}
			if !ok {
				utils.RespondWithError(w, http.StatusUnauthorized, "Error converting claim to string")
				return
			}
			ctx := context.WithValue(r.Context(), "user_id", id)
			nextHandler.ServeHTTP(w, r.WithContext(ctx))
		} else {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusUnauthorized, "An authorized error occurred")
		}
	})
}
