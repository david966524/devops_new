package jenkins

import (
	"context"
	"fmt"

	"github.com/bndr/gojenkins"
)

func GetJenkins(url string, username string, password string) *gojenkins.Jenkins {

	jenkins, err := gojenkins.CreateJenkins(nil, url, username, password).Init(context.Background())
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	return jenkins
}
