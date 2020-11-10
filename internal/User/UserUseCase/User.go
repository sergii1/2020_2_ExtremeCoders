package UserUseCase

import (
	"CleanArch/internal/User/UserModel"
	"CleanArch/internal/User/UserRepository"
	"net/http"
	"time"
)

type UseCase struct{
	Db UserRepository.UserDB
}

func (uc *UseCase)Signup(user UserModel.User) (error, *http.Cookie) {
	if uc.Db.IsEmailExists(user.Email){
		//return 401, nil
		return UserRepository.EmailAlreadyExists, nil
	}

	user.Id = uc.Db.GenerateUID()
	sid := string(uc.Db.GenerateSID())
	uc.Db.AddUser(&user)
	if uc.Db.AddSession(sid, user.Id, &user) != nil {
		//return 401, nil
		return UserRepository.DbError, nil
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
	}
	cookie.Path = "/"

	//return 200, cookie
	return nil, cookie
}

func (uc *UseCase)SignIn(user UserModel.User) (error, *http.Cookie) {
	userEx, erro := uc.Db.GetUserByEmail(user.Email)
	if !erro {
		//return 404, nil
	}
	if userEx.Password != user.Password {
		//return 401, nil

	}
	sid := string(uc.Db.GenerateSID())
	uc.Db.RemoveSessionByUID(userEx.Id)
	if uc.Db.AddSession(sid, userEx.Id, &user) != nil {
		//return 401, nil
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
	}
	cookie.Path = "/"
	//return 200, cookie
	return nil, cookie

}
