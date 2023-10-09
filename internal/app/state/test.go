package state

import (
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type Question1 struct {
}

func (state Question1) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET long_range = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question2{}.PreviewProcess(ctc)
		return &Question2{}
	} else if messageText == "Нет" {
		Question2{}.PreviewProcess(ctc)
		return &Question2{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question1{}
	}
}

func (state Question1) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("1. Любите ли вы стрельбу на дальних дистанциях?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question1) Name() string {
	return "Question1"
}

// ////////////////////////////////////////////////////////
type Question2 struct {
}

func (state Question2) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET stealth_accuracy = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question3{}.PreviewProcess(ctc)
		return &Question3{}
	} else if messageText == "Нет" {
		Question3{}.PreviewProcess(ctc)
		return &Question3{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question2{}
	}
}

func (state Question2) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("2. Предпочитаете ли вы скрытность и точные выстрелы?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question2) Name() string {
	return "Question2"
}

// ////////////////////////////////////////////////////////
type Question3 struct {
}

func (state Question3) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET shoot_on_move = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question4{}.PreviewProcess(ctc)
		return &Question4{}
	} else if messageText == "Нет" {
		Question4{}.PreviewProcess(ctc)
		return &Question4{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question3{}
	}
}

func (state Question3) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("3. Приоритетны ли для вас быстрые движения и стрельба на ходу?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question3) Name() string {
	return "Question3"
}

// ////////////////////////////////////////////////////////
type Question4 struct {
}

func (state Question4) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET power = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question5{}.PreviewProcess(ctc)
		return &Question5{}
	} else if messageText == "Нет" {
		Question5{}.PreviewProcess(ctc)
		return &Question5{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question4{}
	}
}

func (state Question4) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("4. Важна ли вам мощь оружия над скоростью стрельбы?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question4) Name() string {
	return "Question4"
}

// ////////////////////////////////////////////////////////
type Question5 struct {
}

func (state Question5) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET shoot_in_head = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question6{}.PreviewProcess(ctc)
		return &Question6{}
	} else if messageText == "Нет" {
		Question6{}.PreviewProcess(ctc)
		return &Question6{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question5{}
	}
}

func (state Question5) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("5. Предпочитаете ли вы стрельбу в голову для быстрого убийства?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question5) Name() string {
	return "Question5"
}

// ////////////////////////////////////////////////////////
type Question6 struct {
}

func (state Question6) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET accuracy_damage = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question7{}.PreviewProcess(ctc)
		return &Question7{}
	} else if messageText == "Нет" {
		Question7{}.PreviewProcess(ctc)
		return &Question7{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question6{}
	}
}

func (state Question6) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("6. Любите ли вы оружие с высокой точностью и уроном на дальних дистанциях?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question6) Name() string {
	return "Question6"
}

// ////////////////////////////////////////////////////////
type Question7 struct {
}

func (state Question7) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET shoot_bursts = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question8{}.PreviewProcess(ctc)
		return &Question8{}
	} else if messageText == "Нет" {
		Question8{}.PreviewProcess(ctc)
		return &Question8{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question7{}
	}
}

func (state Question7) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("7. Предпочитаете ли вы оружие, способное стрелять очередями?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question7) Name() string {
	return "Question7"
}

// ////////////////////////////////////////////////////////
type Question8 struct {
}

func (state Question8) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET small_cost = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question9{}.PreviewProcess(ctc)
		return &Question9{}
	} else if messageText == "Нет" {
		Question9{}.PreviewProcess(ctc)
		return &Question9{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question8{}
	}
}

func (state Question8) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("8. Важен ли для вас бюджет при выборе оружия?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question8) Name() string {
	return "Question8"
}

// ////////////////////////////////////////////////////////
type Question9 struct {
}

func (state Question9) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET reward = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question10{}.PreviewProcess(ctc)
		return &Question10{}
	} else if messageText == "Нет" {
		Question10{}.PreviewProcess(ctc)
		return &Question10{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question9{}
	}
}

func (state Question9) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("9. Важна ли для вас награда за убийство?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question9) Name() string {
	return "Question9"
}

// ////////////////////////////////////////////////////////
type Question10 struct {
}

func (state Question10) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET close_range = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question11{}.PreviewProcess(ctc)
		return &Question11{}
	} else if messageText == "Нет" {
		Question11{}.PreviewProcess(ctc)
		return &Question11{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question10{}
	}
}

func (state Question10) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("10. Любите ли вы стрельбу на коротких дистанциях?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question10) Name() string {
	return "Question10"
}

// ////////////////////////////////////////////////////////
type Question11 struct {
}

func (state Question11) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET ammo_reserve = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question12{}.PreviewProcess(ctc)
		return &Question12{}
	} else if messageText == "Нет" {
		Question12{}.PreviewProcess(ctc)
		return &Question12{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question11{}
	}
}

func (state Question11) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("11. Предпочитаете ли вы оружие с большим запасом патронов?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question11) Name() string {
	return "Question11"
}

// ////////////////////////////////////////////////////////
type Question12 struct {
}

func (state Question12) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET light = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question13{}.PreviewProcess(ctc)
		return &Question13{}
	} else if messageText == "Нет" {
		Question13{}.PreviewProcess(ctc)
		return &Question13{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question12{}
	}
}

func (state Question12) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("12. Важно ли для вас, чтобы оружие было легким?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question12) Name() string {
	return "Question12"
}

// ////////////////////////////////////////////////////////
type Question13 struct {
}

func (state Question13) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET machine_gun = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question14{}.PreviewProcess(ctc)
		return &Question14{}
	} else if messageText == "Нет" {
		Question14{}.PreviewProcess(ctc)
		return &Question14{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question13{}
	}
}

func (state Question13) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("13. Любите ли вы стрельбу в пулеметном режиме?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question13) Name() string {
	return "Question13"
}

// ////////////////////////////////////////////////////////
type Question14 struct {
}

func (state Question14) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET fast_recharge = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question15{}.PreviewProcess(ctc)
		return &Question15{}
	} else if messageText == "Нет" {
		Question15{}.PreviewProcess(ctc)
		return &Question15{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question14{}
	}
}

func (state Question14) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("14. Предпочитаете ли вы оружие с высокой скоростью перезарядки?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question14) Name() string {
	return "Question14"
}

// ////////////////////////////////////////////////////////
type Question15 struct {
}

func (state Question15) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET low_recoil = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question16{}.PreviewProcess(ctc)
		return &Question16{}
	} else if messageText == "Нет" {
		Question16{}.PreviewProcess(ctc)
		return &Question16{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question15{}
	}
}

func (state Question15) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("15. Важно ли для вас оружие с низкой отдачей?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question15) Name() string {
	return "Question15"
}

// ////////////////////////////////////////////////////////
type Question16 struct {
}

func (state Question16) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET beautiful = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question17{}.PreviewProcess(ctc)
		return &Question17{}
	} else if messageText == "Нет" {
		Question17{}.PreviewProcess(ctc)
		return &Question17{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question16{}
	}
}

func (state Question16) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("16. Важна ли для вас красота оружия?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question16) Name() string {
	return "Question16"
}

// ////////////////////////////////////////////////////////
type Question17 struct {
}

func (state Question17) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET medium_range = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question18{}.PreviewProcess(ctc)
		return &Question18{}
	} else if messageText == "Нет" {
		Question18{}.PreviewProcess(ctc)
		return &Question18{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question17{}
	}
}

func (state Question17) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("17. Предпочитаете ли вы стрельбу на средних дистанциях?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question17) Name() string {
	return "Question17"
}

// ////////////////////////////////////////////////////////
type Question18 struct {
}

func (state Question18) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET rate_of_fire = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question19{}.PreviewProcess(ctc)
		return &Question19{}
	} else if messageText == "Нет" {
		Question19{}.PreviewProcess(ctc)
		return &Question19{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question18{}
	}
}

func (state Question18) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("18. Важно ли для вас оружие с высокой скоростью стрельбы?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question18) Name() string {
	return "Question18"
}

// ////////////////////////////////////////////////////////
type Question19 struct {
}

func (state Question19) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET aggressive = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		Question20{}.PreviewProcess(ctc)
		return &Question20{}
	} else if messageText == "Нет" {
		Question20{}.PreviewProcess(ctc)
		return &Question20{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question19{}
	}
}

func (state Question19) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("19. Любите ли вы агрессивный стиль игры и рашить противника?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question19) Name() string {
	return "Question19"
}

// ////////////////////////////////////////////////////////
type Question20 struct {
}

func (state Question20) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Да" {
		_, err := ctc.Db.Exec("UPDATE users SET accuracy_slow = $1 WHERE vk_id =$2", true, ctc.User.VkID)
		if err != nil {
			log.WithError(err).Error("Failed to update user")
		}
		ResultsState{}.PreviewProcess(ctc)
		return &ResultsState{}
	} else if messageText == "Нет" {
		ResultsState{}.PreviewProcess(ctc)
		return &ResultsState{}
	} else if messageText == "С начала" {
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	} else {
		state.PreviewProcess(ctc)
		return &Question20{}
	}
}

func (state Question20) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("20. Предпочитаете ли вы оружие с высокой точностью и медленной стрельбой?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "positive")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("С начала", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state Question20) Name() string {
	return "Question20"
}
