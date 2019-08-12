package handlers

import (
	"net/http"
	"os"

	appCtx "github.com/HichuYamichu/go-webserver-reference/app/context"
	jwt "github.com/dgrijalva/jwt-go"
)

var secret = os.Getenv("SECRET")

// Authenticate : Sets cookie granting api access
func Authenticate(reqCtx *appCtx.Context, w http.ResponseWriter, r *http.Request) *AppError {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})
	return nil
}
