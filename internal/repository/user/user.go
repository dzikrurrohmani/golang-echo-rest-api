package user

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"time"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/tracing"
	"gorm.io/gorm"
)

type userRepo struct {
	db          *gorm.DB
	gcm         cipher.AEAD
	time        uint32
	memory      uint32
	parallelism uint8
	keyLen      uint32
	secret      string
	signKey     *rsa.PrivateKey
	accessExp   time.Duration
}

func GetRepository(
	db *gorm.DB,
	secret string,
	time, memory,
	keyLen uint32,
	parallelism uint8,
	signKey *rsa.PrivateKey,
	accessExp time.Duration,
) (Repository, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &userRepo{
		db:          db,
		gcm:         gcm,
		time:        time,
		memory:      memory,
		parallelism: parallelism,
		keyLen:      keyLen,
		secret:      secret,
		signKey:     signKey,
		accessExp:   accessExp,
	}, nil
}

func (ur *userRepo) RegisterUser(ctx context.Context, userData model.User) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	if err := ur.db.WithContext(ctx).Create(&userData).Error; err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (ur *userRepo) CheckRegistered(ctx context.Context, username string) (bool, error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckRegistered")
	defer span.End()

	var userData model.User

	if err := ur.db.WithContext(ctx).Where(model.User{Username: username}).First(&userData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}

	return userData.ID != "", nil
}

func (ur *userRepo) GetUserData(ctx context.Context, username string) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetUserData")
	defer span.End()

	var userData model.User

	if err := ur.db.WithContext(ctx).Where(model.User{Username: username}).First(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}

func (ur *userRepo) VerifyLogin(ctx context.Context, username, password string, userData model.User) (bool, error) {
	ctx, span := tracing.CreateSpan(ctx, "VerifyLogin")
	defer span.End()

	if username != userData.Username {
		return false, nil
	}

	verified, err := ur.comparePassword(ctx, password, userData.Hash)
	if err != nil {
		return false, err
	}

	return verified, nil
}
