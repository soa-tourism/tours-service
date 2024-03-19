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

func ParseExecutionStatus(ExecutionStatusStr string) ExecutionStatus {
	for diff, diffStr := range ExecutionStatusStrings {
		if diffStr == ExecutionStatusStr {
			return diff
		}
	}
	return Completed
}
