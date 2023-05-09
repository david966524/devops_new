package jenkins

type JobInfo struct {
	Name    string `json:"name"`
	BuildId string `json:"buildId"`
}
