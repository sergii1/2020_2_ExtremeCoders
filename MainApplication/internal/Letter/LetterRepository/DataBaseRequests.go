package LetterRepository

import (
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"errors"
)

//go:generate mockgen -source=./DataBaseRequests.go -destination=../../../test/mock_LetterRepository/LetterRepositoryMock.go
var DbError = errors.New("Data Base error!")
var ReceiverNotFound = errors.New("Receiver not found!")
var SaveLetterError = errors.New("Save letter error!")

var ReceivedLetterError = errors.New("Could not get received letters!")
var SentLetterError = errors.New("Could not get sent letters!")
var WatchLetterError = errors.New("Could not watch letter!")
var DeleteLetterError = errors.New("Could not delete letter!")
var GetLetterByError = errors.New("Could not get letter by!")

type LetterDB interface {
	SaveMail(LetterModel.Letter) error
	GetReceivedLetters(string, uint64, uint64) (error, []LetterModel.Letter)
	GetSendedLetters(string) (error, []LetterModel.Letter)
	GetReceivedLettersDir(uint64) (error, []LetterModel.Letter)
	GetSendedLettersDir(uint64) (error, []LetterModel.Letter)
	WatchLetter(uint64) (error, LetterModel.Letter)

	DeleteLetter(uint64) error
	FindSimilar(string) string
	GetLetterBy(string, string) (error, []LetterModel.Letter)
}
