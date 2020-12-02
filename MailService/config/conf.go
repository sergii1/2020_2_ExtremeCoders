package config

import "time"

const (
	DbUser     = "postgres"
	DbPassword = "987654321"
	//DbPassword   = "1538"
	DbDB         = "maila"
	Port         = ":8080"
	ReadTimeout  = 10 * time.Second
	WriteTimeout = 10 * time.Second
)

var AllowedOriginsCORS = []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://95.163.209.195:3000", "http://95.163.209.195:80"}
var AllowedHeadersCORS = []string{"Version", "Authorization", "Content-Type", "csrf_token"}
var AllowedMethodsCORS = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

const (
	SizeSID  = 32
	CsrfSize = 32
)

var SidRunes = "1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
var CsrfRunes = "1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
