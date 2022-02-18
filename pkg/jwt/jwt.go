package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/junminhong/member-center-service/pkg/logger"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var zapLogger = logger.Setup()

func getLocalSecretKey(fileName string) []byte {
	nowWorkDir, err := os.Getwd()
	if err != nil {
		zapLogger.Error(err.Error())
	}
	secretKey, err := ioutil.ReadFile(nowWorkDir + "/" + fileName + ".pem")
	if err != nil {
		zapLogger.Error(err.Error())
	}
	return secretKey
}

func GenerateAtomicToken(timeLimit int) string {
	now := time.Now()
	jwtID := strconv.FormatInt(now.Unix(), 10)
	claims := &jwt.StandardClaims{
		ExpiresAt: now.Add(time.Duration(timeLimit) * time.Hour).Unix(),
		Id:        jwtID,
		IssuedAt:  now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(getLocalSecretKey("key"))
	if err != nil {
		zapLogger.Info(err.Error())
	}
	atomicToken, err := token.SignedString(privateKey)
	if err != nil {
		zapLogger.Info(err.Error())
	}
	return atomicToken
}

func VerifyAtomicToken(atomicToken string) bool {
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(getLocalSecretKey("pubkey"))
	tokenParts := strings.Split(atomicToken, ".")
	if len(tokenParts) != 3 {
		return false
	}
	err := jwt.SigningMethodRS256.Verify(strings.Join(tokenParts[0:2], "."), tokenParts[2], pubKey)
	if err != nil {
		zapLogger.Info(err.Error())
	}
	type MyCustomClaims struct {
		jwt.StandardClaims
	}
	token, err := jwt.ParseWithClaims(atomicToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})
	if err != nil {
		zapLogger.Info(err.Error())
	}
	return token.Valid
}
