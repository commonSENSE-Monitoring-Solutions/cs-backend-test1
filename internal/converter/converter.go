package converter

import (
	"cs-backend-test1/internal/model"
	"fmt"
	"time"
)

type UserConverter struct {
	result model.UserAccount
}

type Converter interface {
	Convert(val interface{}) (interface{}, error)
}

func (uc UserConverter) convertToTimeUnix(t *int64) *time.Time {
	if t == nil {
		return nil
	}

	timeUnix := time.Unix(*t, 0)

	return &timeUnix
}

func (uc UserConverter) convertActiveStatusToBool(status int) bool {
	if status == 1 {
		return true
	}

	return false
}

func (uc UserConverter) Convert(val interface{}) (interface{}, error) {
	user, ok := val.(model.User)
	if !ok {
		return nil, fmt.Errorf("unsupported user struct: %T", val)
	}

	result := model.UserAccount{
		Email:     user.EmailAddress,
		Password:  user.Password,
		Salt:      user.Salt,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		OldID:     int64(user.ID),
	}

	result.CreatedAt = uc.convertToTimeUnix(user.CreatedAt)
	result.UpdatedAt = uc.convertToTimeUnix(user.UpdatedAt)
	result.DeletedAt = uc.convertToTimeUnix(user.DeletedAt)

	result.IsActive = uc.convertActiveStatusToBool(user.IsActive)

	return result, nil
}
