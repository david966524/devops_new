package jenkins

import (
	"context"

	"github.com/bndr/gojenkins"
)

func GetJenkins(url string, username string, password string) *gojenkins.Jenkins {

	jenkins, err := gojenkins.CreateJenkins(nil, url, username, password).Init(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return jenkins
}
