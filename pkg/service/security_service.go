package service

import (
	"ftms-go/pkg/entity"
	"strings"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/ant0ine/go-json-rest/rest"
	"golang.org/x/crypto/bcrypt"
)

const secret = "QWERDSAFHXBVZF"

func CheckLogin(username, password string) (map[string]string, int) {
	resp := map[string]string{}
	user, err := GetUserByUsername(username)
	if err != nil {
		resp["error"] = "invalid username"
		return resp, 400
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		resp["error"] = "invalid login"
		return resp, 400
	}
	token, err := createToken(user)
	if err != nil {
		resp["error"] = "create token error"
		return resp, 400
	}
	resp["token"] = token
	resp["username"] = user.Username
	resp["roles"] = strings.Join(user.AppRoles, ",")
	return resp, 200
}

func createToken(user entity.AppUser) (token string, err error) {
	claims := jws.Claims{}
	claims.Set("username", user.Username)
	claims.Set("roles", strings.Join(user.AppRoles, ","))
	claims.SetIssuer("foamtecintl")
	now := time.Now()
	claims.SetIssuedAt(now)
	claims.SetExpiration(now.AddDate(0, 0, 3))
	tokenStruct := jws.NewJWT(claims, crypto.SigningMethodHS256)
	serialized, err := tokenStruct.Serialize([]byte(secret))
	if err != nil {
		return "", err
	}
	return string(serialized), nil
}

func CheckToken(r *rest.Request) (jwt.Claims, error) {
	cms, err := tokenValidator(strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1))
	if err != nil {
		return nil, err
	}
	return cms, nil
}

func tokenValidator(tokenString string) (jwt.Claims, error) {
	token, err := jws.ParseJWT([]byte(tokenString))
	if err != nil {
		return nil, err
	}

	validator := &jwt.Validator{}
	validator.SetIssuer("foamtecintl")

	err = token.Validate([]byte(secret), crypto.SigningMethodHS256, validator)
	return token.Claims(), err
}

func CheckRole(roleMaster, rolesValidate string) bool {
	if strings.Index(rolesValidate, roleMaster) >= 0 {
		return false
	}
	return true
}
