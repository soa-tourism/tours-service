package model

type Status int

const (
	Draft Status = iota
	Published
	Archived
)

var tourStatusStrings = map[Status]string{
	Draft:     "Draft",
	Published: "Published",
	Archived:  "Archived",
}

func (ts Status) String() string {
	return tourStatusStrings[ts]
}

func ParseStatus(statusStr string) Status {
	for status, statusString := range tourStatusStrings {
		if statusString == statusStr {
			return status
		}
	}
	return Draft
}
