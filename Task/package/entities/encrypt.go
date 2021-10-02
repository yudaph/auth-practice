package entities

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"task/errs"
	"task/zapLog"
	"time"

	"github.com/golang-jwt/jwt"
)
const(
	jwtKey string = "MY-LONG-KEY-FOR-JWT"
)

type JwtClaims struct{
	UserDto
	jwt.StandardClaims
}

func EncryptPassword(password *string) string {
	hash := md5.Sum([]byte(*password + "salt (optional)"))
	return hex.EncodeToString(hash[:])
}

func CreateJWT(user *UserDto) (*string, *errs.AppError){
	claims := JwtClaims{
		*user,
		jwt.StandardClaims{
			ExpiresAt: int64(30*time.Minute),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(jwtKey))
	if err!=nil {
		zapLog.Error(err.Error())
		return nil, errs.NewAppError("internal server error", http.StatusInternalServerError)
	}

	return &ss, nil
}

func ValidateJWT(jwtstring *string) (*UserDto, *errs.AppError) {
	token, err := jwt.ParseWithClaims(*jwtstring, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return &claims.UserDto, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			zapLog.Info("Token not even token")
			return nil, errs.NewAppError("token not valid", http.StatusUnauthorized)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			zapLog.Info("Token Expired")
			return nil, errs.NewAppError("token expired", http.StatusUnauthorized)
		} else {
			zapLog.Error("Token couldn't be handle "+ err.Error())
			return nil, errs.NewAppError("token not valid", http.StatusUnauthorized)
		}
	} else {
		zapLog.Error("Token error couldn't be handle "+err.Error())
		return nil, errs.NewAppError("token not valid", http.StatusUnauthorized)
	}
}



