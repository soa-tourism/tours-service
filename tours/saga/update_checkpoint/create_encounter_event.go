package update_checkpoint

type EncounterCreatedEvent struct {
	EncounterId          int64
	CheckpointId         string
	IsSecretPrerequisite bool
}

type CreateEncounterCommandType int8

const (
	UpdateCheckpoint CreateEncounterCommandType = iota
	UnknownCommand
)

type CreateEncounterCommand struct {
	Encounter EncounterCreatedEvent
	Type      CreateEncounterCommandType
}

type CreateEncounterReplyType int8

const (
	CheckpointUpdated CreateEncounterReplyType = iota
	CheckpointNotUpdated
	UnknownReply
)

type CreateEncounterReply struct {
	Encounter EncounterCreatedEvent
	Type      CreateEncounterReplyType
}
