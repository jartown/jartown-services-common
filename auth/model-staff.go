package auth

import (
	common "github.com/singkorn/jartown-services-common"
	"golang.org/x/crypto/bcrypt"
)

const RoleAdmin = 256
const RoleShopOwner = 16
const RoleShopStaff = 4

type UserStaff struct {
	ID        string `bson:"-" json:"id"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"password,omitempty"`
	ShopID    string `bson:"shop_id" json:"shop_id"`
	RoleLevel int64  `bson:"role_level" json:"role_level"`
}

func (u UserStaff) StripPassword() UserStaff {
	u.Password = ""
	return u
}

func (u UserStaff) HashPassword() (UserStaff, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return UserStaff{}, err
	}
	u.Password = string(hashed)
	return u, nil
}

func (u UserStaff) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return common.ErrPasswordMismatch
		}
		return err
	}
	return nil
}
