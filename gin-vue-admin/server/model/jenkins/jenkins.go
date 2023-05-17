package jenkins

type JobInfo struct {
	Name           string `json:"name"`
	BuildId        string `json:"buildId"`
	ParameterName  string `json:"paraName"`
	ParameterValue string `json:"paraValue"`
}
