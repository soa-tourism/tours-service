package events

import (
	"fmt"
	saga "tours/saga/messaging"
	update_checkpoint "tours/saga/update_checkpoint"
	"tours/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCheckpointCommandHandler struct {
	checkpointService *service.CheckpointService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateCheckpointCommandHandler(checkpointService *service.CheckpointService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateCheckpointCommandHandler, error) {
	o := &UpdateCheckpointCommandHandler{
		checkpointService: checkpointService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	fmt.Println("YO")
	if err != nil {
		fmt.Println("YO")
		return nil, err
	}
	return o, nil
}

func (handler *UpdateCheckpointCommandHandler) handle(command *update_checkpoint.CreateEncounterCommand) {
	id, err := primitive.ObjectIDFromHex(command.Encounter.CheckpointId)
	if err != nil {
		return
	}
	reply := update_checkpoint.CreateEncounterReply{Encounter: command.Encounter}

	switch command.Type {
	case update_checkpoint.UpdateCheckpoint:
		checkpoint, _ := handler.checkpointService.FindCheckpoint(id)
		checkpoint.IsSecretPrerequisite = command.Encounter.IsSecretPrerequisite
		checkpoint.EncounterID = command.Encounter.EncounterId
		_, err := handler.checkpointService.UpdateCheckpoint(checkpoint)
		if err != nil {
			reply.Type = update_checkpoint.CheckpointNotUpdated
			fmt.Println("CH NOT UPDATE")
			return
		}
		reply.Type = update_checkpoint.CheckpointUpdated
		fmt.Println("CH UPDATE")
	default:
		reply.Type = update_checkpoint.UnknownReply
	}

	if reply.Type != update_checkpoint.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
