package authManager

import (
	"errors"
	"github.com/kkomissarov/beggar/db"
	"github.com/kkomissarov/beggar/managers/jwtManager"
	"github.com/kkomissarov/beggar/models"
	"github.com/levenlabs/golib/timeutil"
)

func RevokeToken(token string) error {
	claims, err := jwtManager.ExtractClaims(token)
	if err != nil {
		return err
	}

	revokedToken := models.RevokedToken{
		Token:     token,
		ExpiredAt: timeutil.TimestampFromFloat64(claims["exp"].(float64)).Time,
	}

	db.DB.Create(&revokedToken)
	if revokedToken.ID == 0 {
		return errors.New("unable to logout")
	}

	return nil
}
