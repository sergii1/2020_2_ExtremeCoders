package UseCase

import (
	"MailService/internal/Model"
	"MailService/internal/Repository"
)

type UseCase struct{
	re Repository.LetterDB
}

type Interface interface {
	GetLettersSendDir(dir uint64) (error, []Model.Letter)
	GetLettersRecvDir(dir uint64) (error, []Model.Letter)
	GetLettersSend(email string) (error, []Model.Letter)
	GetLettersRecv(email string) (error, []Model.Letter)
	SaveLetter(letter Model.Letter) error
	WatchLetter(lid uint64) (error, Model.Letter)
}

func New(repo Repository.LetterDB) Interface {
	return UseCase{re: repo}
}

func (uc UseCase) GetLettersRecvDir(dir uint64) (error, []Model.Letter){
	err, letters:=uc.re.GetLettersRecvDir(dir)
	return err, letters
}

func (uc UseCase) GetLettersSendDir(dir uint64) (error, []Model.Letter){
	err, letters:=uc.re.GetLettersSentDir(dir)
	return err, letters
}

func (uc UseCase) SaveLetter(letter Model.Letter) error{
	err:=uc.re.SaveMail(letter)
	return err
}

func (uc UseCase) WatchLetter(lid uint64) (error, Model.Letter){
	return uc.re.SetLetterWatched(lid)
}

func (uc UseCase)GetLettersSend(email string) (error, []Model.Letter){
	err, letters:=uc.re.GetLettersSent(email)
	return err, letters
}

func (uc UseCase)GetLettersRecv(email string) (error, []Model.Letter){
	err, letters:=uc.re.GetLettersRecv(email)
	return err, letters
}

