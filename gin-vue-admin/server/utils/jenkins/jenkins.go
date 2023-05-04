package jenkins

import (
	"context"

	"github.com/bndr/gojenkins"
)

func GetJenkins() *gojenkins.Jenkins {
	jenkins, err := gojenkins.CreateJenkins(nil, "jenkinsurl", "admin", "password").Init(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return jenkins
}
