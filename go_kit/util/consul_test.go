package util

import "testing"

func TestUserServiceRegister(t *testing.T) {
	ConsulKVTest()
}

func TestRegisterUserService(t *testing.T) {
	RegisterUserService()
}

func TestDeregisterUserService(t *testing.T) {
	DeregisterUserService()
}