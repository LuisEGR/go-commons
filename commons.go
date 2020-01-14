package commons

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/tidwall/gjson"
)

// Hello function
func Hello() string {
	return "Hello, world."
}

// Base64Encode function
func Base64Encode(src string) string {
	return strings.
		TrimRight(base64.URLEncoding.
			EncodeToString([]byte(src)), "=")
}

// Base64Decode function
func Base64Decode(src string) (string, error) {
	if l := len(src) % 4; l > 0 {
		src += strings.Repeat("=", 4-l)
	}
	decoded, err := base64.URLEncoding.DecodeString(src)
	if err != nil {
		errMsg := fmt.Errorf("Decoding Error %s", err)
		return "", errMsg
	}
	return string(decoded), nil
}

// GetDataJWT function
func GetDataJWT(token string, key string) (string, error) {
	res1 := strings.Split(token, ".")
	if len(res1) < 3 {
		return "", errors.New("Invalid JWT")
	}
	res2, err := Base64Decode(res1[1])
	if err != nil {
		// return "", err
		return "", errors.New("Invalid JWT; " + err.Error())
	}
	info := gjson.Get(res2, key).String()
	return info, nil
}

// GetLogger function
func GetLogger() *logrus.Logger {
	logger := logrus.New()
	env := os.Getenv("ENVIRONMENT")

	if len(env) == 0 {
		env = "DEV"
	}
	// or command-line flag
	if env == "PRODUCTION" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetReportCaller(true)
		// The TextFormatter is default, you don't actually have to do this.
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	return logger
}
