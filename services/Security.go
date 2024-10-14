package services

import (
	"argentina-tresury/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	Username  string `json:"username"`
	ChapterId uint   `json:"chapter_id"`
	jwt.StandardClaims
}

func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func CreateUser(username string, password string) error {
	hashedPassword := HashAndSalt([]byte(password))
	if err := model.DB.Create(&model.User{
		UserName: username,
		Password: hashedPassword,
	}).Error; err != nil {
		return err
	}
	return nil
}

func ValidateUser(username string, password string) *model.User {
	var user model.User
	if err := model.DB.Preload("Chapter").Model(model.User{UserName: username}).First(&user).
		Error; err != nil {
		return nil
	}
	if ComparePasswords(user.Password, []byte(password)) {
		return &user
	} else {
		return nil
	}
}

func GenerateToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username:  user.UserName,
		ChapterId: *user.ChapterID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func RefreshToken(claims *Claims) (string, error) {

	return GenerateToken(&model.User{
		UserName:  claims.Username,
		ChapterID: &claims.ChapterId})
}

func ValidateToken(tknStr string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}

func ChangePassword(username string, password string) error {
	hashedPassword := HashAndSalt([]byte(password))
	if err := model.DB.Model(&model.User{}).Where("user_name = ?", username).Update("password", hashedPassword).Error; err != nil {
		return err
	}
	return nil
}
