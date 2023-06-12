package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/sirupsen/logrus"
)

func GetSessionData(r *http.Request) (model.UserSession, error) {
	authString := r.Header.Get("Authorization")
	splitString := strings.Split(authString, " ")
	if len(splitString) != 2 {
		logrus.WithFields(logrus.Fields{
			"auth_string": authString,
		}).Warn("[delivery][rest][utils][GetSessionData] auth string mismatch")

		return model.UserSession{}, errors.New("unauthorized")
	}
	accessString := splitString[1]

	return model.UserSession{
		JWTToken: accessString,
	}, nil
}
