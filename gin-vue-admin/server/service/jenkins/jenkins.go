package jenkins

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bndr/gojenkins"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/jenkins"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/jenkins"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/telegram"
)

//var jenkinsClient *gojenkins.Jenkins

func jenkinsClient() *gojenkins.Jenkins {
	//jenkinsClient := jenkins.GetJenkins("http://46.137.193.70:8888/", "admin", "P/IPDezTJGTeH9OLD+YU")
	//global.GVA_CONFIG.Jenkins.Url, global.GVA_CONFIG.Jenkins.Username, global.GVA_CONFIG.Jenkins.Password
	jenkinsClient := jenkins.GetJenkins(global.GVA_CONFIG.Jenkins.Url, global.GVA_CONFIG.Jenkins.Username, global.GVA_CONFIG.Jenkins.Password)
	return jenkinsClient
}

type JenkinsService struct{}

// func (js *JenkinsService) GetJobs() (*[]model.JobInfo, error) {
// 	jobs, err := jenkinsClient.GetAllJobs(context.Background())
// 	if err != nil {
// 		log.Println(err.Error())
// 		return nil, err
// 	}
// 	var jobsinfo []model.JobInfo
// 	for _, job := range jobs {
// 		var jobinfo model.JobInfo
// 		jobname := job.GetName()
// 		lastBuild, err := job.GetLastBuild(context.Background())
// 		if err != nil {
// 			continue
// 		}
// 		jobinfo.JobName = jobname
// 		jobinfo.LastBuildResult = lastBuild.Raw.Result
// 		jobinfo.LastBuildUrl = lastBuild.Raw.URL
// 		jobinfo.DisplayName = lastBuild.Raw.DisplayName
// 		// log.Println(jobname, "--------", lastBuild.Raw.Result, "----------", lastBuild.Raw.DisplayName, "-------", lastBuild.Raw.URL, "-----------", lastBuild.Raw.ChangeSet)
// 		jobsinfo = append(jobsinfo, jobinfo)
// 	}

// 	jobss, _ := jenkinsClient.GetAllJobNames(context.Background())
// 	fmt.Println(jobss)
// 	return &jobsinfo, nil
// }

func (js *JenkinsService) GetJobs() (*[]gojenkins.InnerJob, error) {

	jobs, err := jenkinsClient().GetAllJobNames(context.Background())
	if err != nil {
		return nil, err
	}
	return &jobs, nil
}

func (js *JenkinsService) GetJobBuildHistory(jobName string) ([]gojenkins.JobBuild, error) {
	// jobBuilds, err := jenkinsClient.GetAllBuildIds(context.Background(), jobName)
	// if err != nil {
	// 	return nil, err
	// }
	// for _, v := range jobBuilds {
	// 	fmt.Println(v)
	// }
	contxt := context.Background()
	job, err := jenkinsClient().GetJob(contxt, jobName)
	if err != nil {
		return nil, err
	}
	buildId, err := job.GetAllBuildIds(contxt)
	return buildId, err
}

func (js *JenkinsService) GetBuildInfo(jobinfo model.JobInfo) (*gojenkins.BuildResponse, error) {
	id, _ := strconv.Atoi(jobinfo.BuildId)
	contxt := context.Background()
	info, err := jenkinsClient().GetBuild(contxt, jobinfo.Name, int64(id))
	if err != nil {
		return nil, err
	}
	return info.Raw, nil
}

func (js *JenkinsService) GetBuildConsoleOut(jobinfo model.JobInfo) (string, error) {
	id, _ := strconv.Atoi(jobinfo.BuildId)
	contxt := context.Background()
	info, err := jenkinsClient().GetBuild(contxt, jobinfo.Name, int64(id))
	if err != nil {
		return "nil", err
	}
	return info.GetConsoleOutput(contxt), nil
}

func (js *JenkinsService) GetBuildParameters(jobinfo model.JobInfo) ([]gojenkins.ParameterDefinition, error) {
	contxt := context.Background()
	job, err := jenkinsClient().GetJob(contxt, jobinfo.Name)
	if err != nil {
		return nil, err
	}
	params, err := job.GetParameters(contxt)

	return params, err
}

func (js *JenkinsService) Buildjob(jobinfo *model.JobInfo) (int64, error) {
	contxt := context.Background()
	job, _ := jenkinsClient().GetJob(contxt, jobinfo.Name)
	buildId, err := job.Jenkins.BuildJob(contxt, jobinfo.Name, nil)
	if err != nil {
		return buildId, err
	}
	msg := fmt.Sprintf(`
	jenkins message
		%v 已经触发构建 
		url: %v
	`, jobinfo.Name, job.Raw.URL)
	go func() {
		telegram.SendTgMsg(msg)
	}()
	return buildId, err
}

func (js *JenkinsService) BuildJobParameter(jobinfo *model.JobInfo) (int64, error) {
	contxt := context.Background()
	paramsMap := make(map[string]string)
	paramsMap[jobinfo.ParameterName] = jobinfo.ParameterValue
	job, _ := jenkinsClient().GetJob(contxt, jobinfo.Name)
	buildId, err := job.Jenkins.BuildJob(contxt, jobinfo.Name, paramsMap)
	if err != nil {
		return buildId, err
	}
	msg := fmt.Sprintf(`
	jenkins message
		%v 已经触发构建 
		url: %v
	`, jobinfo.Name, job.Raw.URL)
	go func() {
		telegram.SendTgMsg(msg)
	}()
	return buildId, err
}
