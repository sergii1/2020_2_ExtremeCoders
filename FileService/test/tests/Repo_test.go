package test

import (
	"FileService/internal/File/Repository/FileSystem"
	proto "FileService/proto"
	"testing"
)

func TestGetAvatarRepo(t *testing.T) {
	fs := FileSystem.New()
	_, err := fs.GetAvatar(&proto.User{})
	if err != nil {
		t.Log("TestGetAvatarRepo ", err.Error())
	}
}

func TestGetFilesRepo(t *testing.T) {
	fs := FileSystem.New()
	_, err := fs.GetFiles(&proto.LetterId{})
	if err != nil {
		t.Log("TestGetFilesRepo ", err.Error())
	}
}

func TestGetDefaultAvatarRepo(t *testing.T) {
	fs := FileSystem.New()
	_, err := fs.GetDefaultAvatar()
	if err != nil {
		t.Log("TestGetDefaultAvatarRepo ", err.Error())
	}
}

func TestSaveAvatarRepo(t *testing.T) {
	fs := FileSystem.New()
	err := fs.SaveAvatar(&proto.Avatar{})
	if err == nil {
		t.Log("TestGetDefaultAvatarRepo ")
	}

	err = fs.SaveAvatar(&proto.Avatar{Email: "selicium", FileName: "img.jpeg",Content: nil})
	if err == nil {
		t.Log("TestGetDefaultAvatarRepo ")
	}
}

func TestSaveFilesRepo(t *testing.T) {
	fs := FileSystem.New()
	err := fs.SaveFiles(&proto.Files{})
	if err == nil {
		t.Log("TestGetDefaultAvatarRepo ")
	}
}