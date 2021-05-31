package app

import (
	"encoding/json"
	"go-api-basic/domain"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	// TODO: エラーハンドリング
	if data["password"] != data["password_confirm"] {
		http.NotFound(w, r)
		return
	}

	user := domain.User{
		Name:  data["name"],
		Email: data["email"],
	}

	user.SetPassword(data["password"])

	DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	var user domain.User

	DB.Where("email = ?", data["email"]).First(&user)

	// TODO: エラーハンドリング
	if user.Id == 0 {
		http.NotFound(w, r)
		return
	}

	// TODO: エラーハンドリング
	if err := user.ComparePassword(data["password"]); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		// result.Status = http.StatusBadRequest
		// result.Message = "Content-Type entity header is need application/json"
		return
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		w.WriteHeader(400)
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

type Claims struct {
	jwt.StandardClaims
}

func User(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})
	if err != nil || !token.Valid {
		log.Println("Error while scaning customer" + err.Error())
	}

	claims := token.Claims.(*Claims)
	id := claims.Issuer

	var user domain.User
	DB.Where("id = ?", id).First(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
}
