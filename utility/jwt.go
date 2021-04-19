package utility

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gitlab.com/firstkungz/log-go"
)

var key string

func JWTLoadKey() {
	key = viper.GetString("jwt.key1")
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

func CreateToken(userid string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", key) //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	// atClaims["user_macid"] = userid
	// atClaims["exp"] = time.Now().Add(0).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}
		// userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			// UserId:     userId,
		}, nil
	}
	return nil, err
}
func FetchAuth(authD *AccessDetails) (string, error) {
	userid := authD.AccessUuid
	// userID, _ := strconv.ParseUint(userid, 10, 64)
	return userid, nil
}
