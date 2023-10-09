package state

import (
	"car_bot/internal/app/ds"
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type ChatContext struct {
	User *ds.User
	Vk   *api.VK
	Ctx  *context.Context
	Db   *sqlx.DB
}

type State interface {
	Name() string                                      //получаем название состояния в виде строки, чтобы в дальнейшем куда-то записать(БД)
	Process(ChatContext, object.MessagesMessage) State //нужно взять контекст, посмотреть на каком состоянии сейчас пользователь, метод должен вернуть состояние
	PreviewProcess(ctc ChatContext)
}

// ////////////////////////////////////////////////////////
type StartState struct {
}

func (state StartState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Начать тест" {
		_, err := ctc.Db.Exec("UPDATE users SET long_range = $1, stealth_accuracy = $2, shoot_on_move = $3, power = $4, shoot_in_head = $5, accuracy_damage = $6, shoot_bursts = $7, small_cost = $8, reward = $9, close_range = $10, ammo_reserve = $11, light = $12, machine_gun = $13, fast_recharge = $14, low_recoil = $15, beautiful = $16, medium_range = $17, rate_of_fire = $18, aggressive = $19, accuracy_slow = $20 WHERE vk_id =$21", false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to reset user fields")
		}
		Question1{}.PreviewProcess(ctc)
		return &Question1{}
	} else {
		state.PreviewProcess(ctc)
		return &StartState{}
	}
}

func (state StartState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Какое оружие из CS2 вам подходит больше?\n Для начала тестирования нажмите кнопку \"Начать тест\"")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Начать тест", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state StartState) Name() string {
	return "StartState"
}
