package user

import (
	"context"
	"errors"
	"time"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/tracing"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func (ur *userRepo) CreateUserSession(ctx context.Context, userID string) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateUserSession")
	defer span.End()

	accessToken, err := ur.generateAccessToken(ctx, userID)
	if err != nil {
		return model.UserSession{}, err
	}

	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (ur *userRepo) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	_, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

	accessToken, err := jwt.ParseWithClaims(data.JWTToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return &ur.signKey.PublicKey, nil
	})
	if err != nil {
		return "", errors.New("access token expired/invalid")
	}
	accessTokenClaims, ok := accessToken.Claims.(*Claims)
	if !ok {
		return "", errors.New("unauthorized")
	}

	if accessToken.Valid {
		return accessTokenClaims.Subject, nil
	}

	return "", errors.New("unauthorized")
}

func (ur *userRepo) generateAccessToken(ctx context.Context, userID string) (string, error) {
	_, span := tracing.CreateSpan(ctx, "generateAccessToken")
	defer span.End()

	accessTokenExp := time.Now().Add(ur.accessExp).Unix()
	accessClaims := Claims{
		jwt.StandardClaims{
			ExpiresAt: accessTokenExp,
			Subject:   userID,
		},
	}

	accessJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaims)

	return accessJwt.SignedString(ur.signKey)
}
