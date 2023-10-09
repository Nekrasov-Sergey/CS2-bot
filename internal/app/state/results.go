package state

import (
	"car_bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type ResultsState struct {
}

func (state ResultsState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	StartState{}.PreviewProcess(ctc)
	return &StartState{}

}

func (state ResultsState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Отлично! Вот ваш результат:")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Пройти тест еще раз", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
	output, err := repository.GetResults(ctc.Db, ctc.User.VkID)
	b.RandomID(0)
	b.Message(output)
	b.PeerID(ctc.User.VkID)
	_, err = ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get send results")
		log.Error(err)
	}
	err = repository.AddSession(ctc.Db, ctc.User.VkID, output)
	if err != nil {
		log.Println("Failed to add new session")
		log.Error(err)
	}
}

func (state ResultsState) Name() string {
	return "ResultsState"
}
