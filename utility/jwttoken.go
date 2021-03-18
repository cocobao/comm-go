package utility

import (
	"time"

	"git.qhfct.io/comm-go/log"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/models"
)

var (
	TokenGen *generates.JWTAccessGenerate
)

func SetTokenRandKey(rangeKey string) {
	TokenGen = generates.NewJWTAccessGenerate([]byte(rangeKey), jwt.SigningMethodHS512)
}

func GetAccessToken(userId string, exp time.Duration) string {
	m := models.NewToken()
	m.SetClientID(userId)
	m.SetUserID(userId)
	m.SetRedirectURI("")
	m.SetScope("user")
	m.SetAccessCreateAt(time.Now())
	m.SetAccessExpiresIn(exp)

	access, _, err := TokenGen.Token(&oauth2.GenerateBasic{
		Client: &models.Client{
			ID:     userId,
			Secret: "WS1ZI446GDHFrwwAOwYtMDD4hx1nTrBG",
		},
		UserID:    userId,
		TokenInfo: m,
	}, false)
	if err != nil {
		log.Error(err)
		return ""
	}
	return access
}
