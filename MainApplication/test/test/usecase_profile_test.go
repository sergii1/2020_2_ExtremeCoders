package test

import (
	"Mailer/MainApplication/internal/User/UserModel"
	"Mailer/MainApplication/internal/User/UserRepository"
	"Mailer/MainApplication/internal/User/UserUseCase"
	mock "Mailer/MainApplication/test/mock_UserRepository"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().UpdateProfile(user, user.Email).Return(nil)
	uc := UserUseCase.New(mockLetter)
	_ = uc.Profile(user)
}

func TestProfileUpUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserDB(ctrl)
	mockLetter.EXPECT().UpdateProfile(user, user.Email).Return(UserRepository.CantUpdateUser)
	uc := UserUseCase.New(mockLetter)
	_ = uc.Profile(user)
}
