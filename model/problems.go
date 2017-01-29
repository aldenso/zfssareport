package model

import (
	"fmt"
)

//Problems struct for all zfssa problems.
type Problems struct {
	List []Problem `json:"problems"`
}

//Problem struct for problem.
type Problem struct {
	Action      string `json:"action"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Diagnosed   string `json:"diagnosed"`
	//Diagnosed   time.Time  `json:"diagnosed"`
	HREF       string `json:"href"`
	Impact     string `json:"impact"`
	PhonedHome string `json:"phoned_home"`
	//PhonedHome time.Time `json:"phoned_home"`
	Response string `json:"response"`
	Severity string `json:"severity"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	UUID     string `json:"uuid"`
}

//PrintProblemInfo print some info about the problem.
func (problem *Problem) PrintProblemInfo() {
	line := fmt.Sprintln("---------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("Code: %-15s Type: %-10s Severity: %-10s\n"+
		"Diagnosed: %s\nDescription: %s\n"+
		"Impact: %s\nResponse: %s\n%s",
		problem.Code, problem.Type, problem.Severity,
		problem.Diagnosed, problem.Description,
		problem.Impact, problem.Response, line)
}
