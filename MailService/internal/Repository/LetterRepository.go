package Repository

import (
	"Mailer/MailService/internal/Model"
	"errors"
)

var DbError = errors.New("Data Base error!")
var ReceiverNotFound = errors.New("Receiver not found!")
var SaveLetterError = errors.New("Save letter error!")

var ReceivedLetterError = errors.New("Could not get received letters!")
var SentLetterError = errors.New("Could not get sent letters!")
var GetByLidError = errors.New("Could not get letter by lid!")
var SetLetterWatchedError = errors.New("Could not set letter watched!")
var DeleteLetterError = errors.New("Could not delete letter!")
var GetLetterByError = errors.New("Could not get letter by!")

//go:generate mockgen -source=./LetterRepository.go -destination=../../test/mock_LetterRepository/RepositoryMock.go
type LetterDB interface {
	SaveMail(Model.Letter) error
	GenerateLID() uint64
	SetLetterWatched(uint64) (error, Model.Letter)
	GetLetterByLid(uint64) (error, Model.Letter)

	GetLettersRecvDir(uint64, uint64, uint64) (error, []Model.Letter)
	GetLettersSentDir(uint64) (error, []Model.Letter)
	GetLettersRecv(string, uint64, uint64) (error, []Model.Letter)
	GetLettersSent(string) (error, []Model.Letter)
	GetLettersByFolder(uint64) (error, []Model.Letter)

	AddLetterToDir(uint64, uint64, bool) error
	RemoveLetterFromDir(uint64, uint64, bool) error
	RemoveDir(uint64, bool) error
	RemoveLetter(uint64) error

	FindSender(string) ([]string, error)
	FindReceiver(string) ([]string, error)
	FindTheme(string) ([]string, error)
	FindText(string) ([]string, error)

	GetLetterBy(string, string) (error, []Model.Letter)
}
