package ginex

import (
	"strings"
	"time"

	"github.com/cocobao/comm-go/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/models"
)

var (
	//路由白名单
	WhiteList = []string{}

	//过滤中间路由层
	MidPath string

	//随机密钥
	rangeKey string

	TokenGen *generates.JWTAccessGenerate
)

func AddWhiteList(addrs []string) {
	WhiteList = append(WhiteList, addrs...)
}

func SetMinPath(p string) {
	MidPath = p
}

func SetRandKey(k string) {
	rangeKey = k
	TokenGen = generates.NewJWTAccessGenerate([]byte(rangeKey), jwt.SigningMethodHS512)
}

func GetAccessToken(uid, Secret string, tokenExpire time.Duration) string {
	m := models.NewToken()
	m.SetClientID(uid)
	m.SetUserID(uid)
	m.SetRedirectURI("")
	m.SetScope("user")
	m.SetAccessCreateAt(time.Now())
	m.SetAccessExpiresIn(tokenExpire)

	access, _, err := TokenGen.Token(&oauth2.GenerateBasic{
		Client: &models.Client{
			ID:     uid,
			Secret: Secret,
		},
		UserID:    uid,
		TokenInfo: m,
	}, false)
	if err != nil {
		log.Error(err)
		return ""
	}
	return access
}

func Xsfr(ctx *gin.Context) {
	path := ctx.Request.URL.Path

	//过滤中间路层
	if MidPath != "" {
		path = strings.ReplaceAll(path, MidPath, "")
	}

	//过滤白名单列表
	for _, uri := range WhiteList {
		if len(uri) <= len(path) && strings.Compare(uri, path[:len(uri)]) == 0 {
			ctx.Next()
			return
		}
	}

	//从http头部获取token信息
	Bearer := ctx.Request.Header.Get("Authorization")
	if strings.Index(Bearer, "Bearer") < 0 {
		log.Error("invalid Bearer:", Bearer)
		ctx.AbortWithStatus(405)
		return
	}

	auths := strings.Split(Bearer, " ")
	var accesstoken string
	if len(auths) > 1 {
		accesstoken = auths[1]
	} else {
		log.Error("invalid Bearer:", Bearer)
		ctx.AbortWithStatus(405)
		return
	}

	//jwt解密
	jti := &generates.JWTAccessClaims{}
	t, err := jwt.ParseWithClaims(accesstoken, jti, func(t *jwt.Token) (interface{}, error) {
		return []byte(rangeKey), nil
	})

	if err != nil {
		log.Info(err)
		ctx.AbortWithStatus(405)
		return
	}

	if !t.Valid || jti.Subject == "" {
		ctx.AbortWithStatus(405)
		return
	}

	//缓存用户id
	ctx.Set("uid", jti.Subject)

	ctx.Next()
}
