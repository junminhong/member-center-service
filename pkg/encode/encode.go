package encode

import (
	"encoding/base64"
	"github.com/junminhong/member-center-service/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

var zapLogger = logger.Setup()

func EmailToken(email string) string {
	pwd := []byte(email)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		zapLogger.Info(err.Error())
	}
	return base64.StdEncoding.EncodeToString([]byte(hash))
}
