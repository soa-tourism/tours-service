package model

type ExecutionStatus int

const (
	Completed ExecutionStatus = iota
	Abandoned
	InProgress
)

var ExecutionStatusStrings = map[ExecutionStatus]string{
	Completed:  "Completed",
	Abandoned:  "Abandoned",
	InProgress: "InProgress",
}

func (d ExecutionStatus) String() string {
	return ExecutionStatusStrings[d]
}

func ParseExecutionStatus(status string) ExecutionStatus {
	var result ExecutionStatus
	if status == "2" || status == "InProgress" {
		result = InProgress
		return result
	}
	if status == "1" || status == "Abandoned" {
		result = Abandoned
		return result
	}
	if status == "0" || status == "Completed" {
		result = Completed
		return result
	}
	return InProgress
}
