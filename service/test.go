package service

import (
	"fmt"

	"github.com/spf13/viper"
)

var testService *TestService

type TestService struct{}

func GetTestService() *TestService {
	if testService == nil {
		testService = &TestService{}
	}
	return testService
}

func (t *TestService) Ping() string {
	return "Pong"
}

func (t *TestService) TestName(name string) string {
	configName := viper.GetString("testName")
	return fmt.Sprintf("request name is %s , config name is %s", name, configName)
}
