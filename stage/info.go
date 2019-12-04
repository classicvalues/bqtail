package stage

import (
	"fmt"
	"github.com/viant/toolbox"
	"path"
	"strings"
	"time"
)

const (
	//PathElementSeparator path separator
	PathElementSeparator = "--"
	//DispatchJob dispatch job name
	DispatchJob = "dispatch"
	//TailJob tail job name
	TailJob = "tail"

	nopAction = "nop"
)

//Info represents processing stage
type Info struct {
	SourceURI string  `json:",omitempty"`
	DestTable    string  `json:",omitempty"`
	EventID string  `json:",omitempty"`
	Action  string  `json:",omitempty"`
	Suffix  string  `json:",omitempty"`
	Step    int  `json:",omitempty"`
	Async   bool  `json:",omitempty"`
	RuleURL string  `json:",omitempty"`
}

//ID returns stage ID
func (i *Info) ID() string {
	return path.Join(i.DestTable, fmt.Sprintf("%v_%05d_%v", i.EventID, i.Step%99999, i.Action)+PathElementSeparator+i.Suffix)
}

//nopActions represents nop actions
var nopActions = map[string]bool{
	"move":   true,
	"delete": true,
	"notify": true,
}

func (i *Info) ChildInfo(action string, step int) *Info {
	step = (i.Step * 1000) + step
	suffix := i.Suffix
	if nopActions[action] {
		suffix = nopAction
	}
	return &Info{
		SourceURI:i.SourceURI,
		RuleURL:i.RuleURL,
		DestTable:    i.DestTable,
		Suffix:  suffix,
		EventID: i.EventID,
		Step:    step,
		Action:  action,
		Async: i.Async,
	}
}



func (i *Info) Wrap(action string) *Info {
	suffix := i.Suffix
	if nopActions[action] {
		suffix = nopAction
	}
	return &Info{
		SourceURI:i.SourceURI,
		RuleURL:i.RuleURL,
		DestTable:    i.DestTable,
		Suffix:  suffix,
		EventID: i.EventID,
		Step:    i.Step+1,
		Action:  action,
		Async: i.Async,
	}
}


//GetJobID returns  a job ID
func (i *Info) GetJobID() string {
	ID := i.ID()
	return Decode(ID)
}

//Decode decode path base ID to big query Job ID
func Decode(jobID string) string {
	if count := strings.Count(jobID, "/"); count > 0 {
		jobID = strings.Replace(jobID, "/", PathElementSeparator, count)
	}
	if count := strings.Count(jobID, "$"); count > 0 {
		jobID = strings.Replace(jobID, "$", "_", count)
	}
	if count := strings.Count(jobID, ":"); count > 0 {
		jobID = strings.Replace(jobID, ":", "_", count)
	}
	if count := strings.Count(jobID, "."); count > 0 {
		jobID = strings.Replace(jobID, ".", "_", count)
	}
	return jobID
}

//Parse parse encoded job ID
func Parse(encoded string) *Info {
	encoded = strings.Replace(encoded, PathElementSeparator, "/", strings.Count(encoded, PathElementSeparator))
	result := &Info{
		Suffix: TailJob,
		Action: nopAction,
	}

	if strings.HasSuffix(encoded, DispatchJob) {
		result.Suffix = DispatchJob
	}
	elements := strings.Split(encoded, "/")
	if len(elements) < 3 {
		result.DestTable = strings.Join(elements[:len(elements)-1], PathElementSeparator)
		result.EventID = fmt.Sprintf("%v", time.Now().Nanosecond()%10000)
	} else {
		eventOffset := len(elements) - 2
		eventElements := strings.Split(elements[eventOffset], "_")
		result.DestTable = strings.Join(elements[:eventOffset], PathElementSeparator)

		result.EventID = eventElements[0]
		if len(eventElements) > 2 {
			result.EventID = eventElements[0]
			result.Step = toolbox.AsInt(eventElements[1])
			result.Action = eventElements[2]
		}
	}
	result.Async = strings.HasSuffix(result.Suffix, DispatchJob)
	return result
}

//AsMap returns info map
func (i Info) AsMap() map[string]interface{} {
	aMap := map[string]interface{}{}
	_ = toolbox.DefaultConverter.AssignConverted(&aMap, i)
	return aMap
}

//New create a processing stage info
func New(source, dest, eventID, action, suffix string, async bool, step int, ruleURL string) *Info {
	return &Info{
		SourceURI:source,
		DestTable:    dest,
		EventID: eventID,
		Action:  action,
		Suffix:  suffix,
		Step:    step,
		Async:async,
		RuleURL:ruleURL,
	}
}