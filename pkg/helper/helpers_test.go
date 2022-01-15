package helper

import (
	"os"
	"strings"
	"testing"
)

func Test_ReadConfig(t *testing.T) {

	cfg := ReadConfig(ReadEnv())

	if cfg == nil {
		t.Errorf("Config is nil")
	}

}

func Test_ReadEnv(t *testing.T) {

	mockEnv := "prod"
	os.Setenv("GO_ENV", mockEnv)

	env := ReadEnv()

	if !strings.EqualFold(mockEnv, env) {
		t.Errorf("Wrong environment. Actual %s", env)
	}

}

func Test_ReadPort(t *testing.T) {

	mockPort := "1923"

	os.Setenv("PORT", mockPort)

	port := readPort()

	if !strings.EqualFold(mockPort, port) {
		t.Errorf("Wrong port value. Actual %s", port)
	}

}
