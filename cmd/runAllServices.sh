#!/bin/bash
#chmod ugo+x runAllServices.sh
#cd ..
#pwd
echo this is deploy in path
pwd
go run ./FileService/cmd/main.go &>FileService.txt &
go run ./MailService/cmd/main.go &>MailService.txt &
go run ./UserService/cmd/main.go &>UserService.txt &
go run ./MainApplication/cmd/main.go &>MainApplication.txt &
