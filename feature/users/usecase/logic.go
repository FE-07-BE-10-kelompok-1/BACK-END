package usecase

import (
	"bookstore/domain"
	"bookstore/feature/users/data"
	"errors"
	"log"

	validator "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUsecase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(ud domain.UserData, v *validator.Validate) domain.UserUsecase {
	return &userUsecase{
		userData: ud,
		validate: v,
	}
}

func (ud *userUsecase) AddUser(newUser domain.User) (domain.User, error) {
	var cnv = data.FromModel(newUser)
	err := ud.validate.Struct(cnv)
	if err != nil {
		log.Println("Validation error : ", err.Error())
		return domain.User{}, err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrypt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	inserted, err := ud.userData.Insert(newUser)

	if err != nil {
		log.Println("error from usecase", err.Error())
		return domain.User{}, err
	}
	if inserted.ID == 0 {
		return domain.User{}, errors.New("cannot insert data")
	}
	return inserted, nil
}

func (ud *userUsecase) LoginUser(userLogin domain.User) (response int, data domain.User, err error) {
	response, data, err = ud.userData.Login(userLogin)

	return response, data, err
}

func (ud *userUsecase) UpdateUser(id int, updateProfile domain.User) (domain.User, error) {
	if id == -1 {
		return domain.User{}, errors.New("invalid user")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(updateProfile.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("error encrypt password", err)
		return domain.User{}, err
	}
	updateProfile.Password = string(hashed)
	result := ud.userData.Update(id, updateProfile)

	if result.ID == 0 {
		return domain.User{}, errors.New("error update user")
	}
	return result, nil
}

func (ud *userUsecase) GetProfile(id int) (domain.User, error) {
	data, err := ud.userData.GetSpecific(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}

func (ud *userUsecase) DeleteUser(id int) (row int, err error) {
	row, err = ud.userData.Delete(id)
	if err != nil {
		log.Println("delete usecase error", err.Error())
		if err == gorm.ErrRecordNotFound {
			return row, errors.New("data not found")
		} else {
			return row, errors.New("failed to delete user")
		}
	}
	return row, nil
}
