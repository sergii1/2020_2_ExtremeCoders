package UserDelivery

import (
	"CleanArch/internal/User/UserModel"
	"CleanArch/internal/User/UserUseCase"
	"CleanArch/internal/errors"
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Delivery struct {
	Uc UserUseCase.UseCase
}

func GetStrFormValueSafety(r *http.Request, field string) string {
	return r.FormValue(field)
}

func (de *Delivery) Session(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		de.SignIn(w, r)
	}
	if r.Method == http.MethodDelete {
		de.Logout(w, r)
	}
}

func (de *Delivery) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	var user UserModel.User
	user.Name = GetStrFormValueSafety(r, "name")
	user.Surname = GetStrFormValueSafety(r, "surname")
	user.Email = GetStrFormValueSafety(r, "email")
	user.Password = GetStrFormValueSafety(r, "password1")
	de.LoadFile(&user, r)
	err, sid := de.Uc.Signup(user)
	var response []byte
	if err == nil {
		cookie := &http.Cookie{
			Name:    "session_id",
			Value:   sid,
			Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
		}
		cookie.Path = "/"
		http.SetCookie(w, cookie)
		response = SignUpError(err, cookie)
	} else {
		response = SignUpError(err, nil)
	}

	w.Write(response)
}

func (de *Delivery) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write(errors.GetErrorNotPostAns())
		return
	}
	var user UserModel.User

	user.Email = GetStrFormValueSafety(r, "email")
	user.Password = GetStrFormValueSafety(r, "password")
	err, sid := de.Uc.SignIn(user)
	var response []byte
	if err == nil {
		cookie := &http.Cookie{
			Name:    "session_id",
			Value:   sid,
			Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
		}
		cookie.Path = "/"
		http.SetCookie(w, cookie)
		response = SignInError(err, cookie)
	} else {
		response = SignInError(err, nil)
	}
	w.Write(response)
}

func (de *Delivery) GetUserByRequest(r *http.Request) (*UserModel.User, *http.Cookie, uint16) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return nil, nil, 401
	}
	uid, ok := de.Uc.Db.IsOkSession(session.Value)
	if ok != nil {
		return nil, nil, 402
	}
	user, err := de.Uc.Db.GetUserByUID(uid)
	if err != nil {
		return nil, nil, 402
	}
	return user, session, 200
}

func (de *Delivery) Profile(w http.ResponseWriter, r *http.Request) {
	user, session, err := de.GetUserByRequest(r)
	if err != 200 {
		w.Write(CookieError(err))
		return
	}
	if r.Method == http.MethodGet {
		w.Write(errors.GetOkAnsData(session.Value, *user))
		return
	} else if r.Method != http.MethodPut {
		w.Write(errors.GetErrorNotPostAns())
		return
	} else {
		var up UserModel.User
		up.Name = GetStrFormValueSafety(r, "profile_firstName")
		up.Surname = GetStrFormValueSafety(r, "profile_lastName")
		de.LoadFile(&up, r)
		err := de.Uc.Profile(up)
		w.Write(ProfileError(err, session))
		return
	}
	w.Write(errors.GetErrorUnexpectedAns())
}

func (de *Delivery) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Write(errors.GetErrorNotPostAns())
		return
	} else {
		_, session, err := de.GetUserByRequest(r)
		if err != 200 {
			w.Write(CookieError(err))
			return
		}

		e := de.Uc.Logout(session.Value)
		if e == nil {
			session.Expires = time.Now().AddDate(0, 0, -1)
			http.SetCookie(w, session)
		}
		w.Write(LogoutError(e))
		return
	}

	w.Write(errors.GetErrorUnexpectedAns())
}

func (de *Delivery) LoadFile(user *UserModel.User, r *http.Request) {
	file, fileHeader, err := r.FormFile("avatar")
	if file == nil {
		return
	}
	(*user).Img = fileHeader.Filename
	path := "./" + (*user).Email + "/" + fileHeader.Filename
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func (de *Delivery) GetAvatar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Write([]byte(""))
		return
	}
	if r.Method == http.MethodGet {
		user, _, Err := de.GetUserByRequest(r)
		if Err != 200 {
			CookieError(Err)
			return
		}
		if (*user).Img == "" {
			(*user).Img = "./default.jpeg"
		}

		file, err := os.Open((*user).Img)
		if err != nil {
			w.Write(errors.GetErrorUnexpectedAns())
			return
		}
		img, _, err := image.Decode(file)
		if err != nil {
			w.Write(errors.GetErrorUnexpectedAns())
			return
		}

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, img, nil); err != nil {
			w.Write(errors.GetErrorUnexpectedAns())
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
		if _, err := w.Write(buffer.Bytes()); err != nil {
			w.Write(errors.GetErrorUnexpectedAns())
			return
		}
		return
	}
}
