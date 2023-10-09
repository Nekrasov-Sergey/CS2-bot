package app

import (
	"car_bot/internal/app/config"
	"car_bot/internal/app/ds"
	"car_bot/internal/app/dsn"
	"car_bot/internal/app/state"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"net/http"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/jmoiron/sqlx"
)

type App struct {
	ctx context.Context

	vk *api.VK
	lp *longpoll.LongPoll

	// db подключение к БД
	db *sqlx.DB
}

func NewApp(ctx context.Context) (*App, error) {
	cfg := config.FromContext(ctx)
	vk := api.NewVK(cfg.VKToken)
	//получаем всю инфу про группу
	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.WithError(err).Error("cant get groups by id")

		return nil, err
	}
	// БД
	db, err := sqlx.Connect("postgres", dsn.FromEnv())
	if err != nil {
		log.Println("nen", err)
		return nil, err
	}
	//starting long poll
	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Println("error on request")
		log.Error(err)
	}
	app := &App{
		ctx: ctx,
		vk:  vk,
		lp:  lp,
		db:  db,
	}
	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	var err error
	go func() error {
		if err = InitSysRoutes(ctx); err != nil {
			log.WithError(err).Error("can't InitSysRoute")
			return err
		}
		return nil
	}()

	var BotUser *ds.User
	var BotUsers []*ds.User
	// New message event
	a.lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)
		//смотрим, новый ли пользователь
		//query := "SELECT * FROM users WHERE vk_id = $1" //+ strconv.Itoa(obj.Message.PeerID)
		err := a.db.Select(&BotUsers, "SELECT * FROM users WHERE vk_id = $1", strconv.Itoa(obj.Message.PeerID))
		if err != nil {
			log.WithError(err).Error("Failed to scan user")
			return
		}

		//if the user writes for the first time, add to the database
		if len(BotUsers) == 0 {
			BotUser = &ds.User{}
			BotUser.VkID = obj.Message.PeerID
			BotUser.State = "StartState"
			_, err := a.db.ExecContext(a.ctx, "INSERT INTO users VALUES ($1, $2)", BotUser.VkID, BotUser.State)
			if err != nil {
				log.WithError(err).Error("Failed to set new user")
				return
			}
		} else {
			BotUser = BotUsers[0]
		}
		strInState := map[string]state.State{
			(&(state.StartState{})).Name():   &(state.StartState{}),
			(&(state.Question1{})).Name():    &(state.Question1{}),
			(&(state.Question2{})).Name():    &(state.Question2{}),
			(&(state.Question3{})).Name():    &(state.Question3{}),
			(&(state.Question4{})).Name():    &(state.Question4{}),
			(&(state.Question5{})).Name():    &(state.Question5{}),
			(&(state.Question6{})).Name():    &(state.Question6{}),
			(&(state.Question7{})).Name():    &(state.Question7{}),
			(&(state.Question8{})).Name():    &(state.Question8{}),
			(&(state.Question9{})).Name():    &(state.Question9{}),
			(&(state.Question10{})).Name():   &(state.Question10{}),
			(&(state.Question11{})).Name():   &(state.Question11{}),
			(&(state.Question12{})).Name():   &(state.Question12{}),
			(&(state.Question13{})).Name():   &(state.Question13{}),
			(&(state.Question14{})).Name():   &(state.Question14{}),
			(&(state.Question15{})).Name():   &(state.Question15{}),
			(&(state.Question16{})).Name():   &(state.Question16{}),
			(&(state.Question17{})).Name():   &(state.Question17{}),
			(&(state.Question18{})).Name():   &(state.Question18{}),
			(&(state.Question19{})).Name():   &(state.Question19{}),
			(&(state.Question20{})).Name():   &(state.Question20{}),
			(&(state.ResultsState{})).Name(): &(state.ResultsState{}),
		}
		ctc := state.ChatContext{
			User: BotUser,
			Vk:   a.vk,
			Db:   a.db,
			Ctx:  &ctx,
		}
		step := strInState[BotUser.State]
		nextStep := step.Process(ctc, obj.Message)
		BotUser.State = nextStep.Name()
		_, err = a.db.ExecContext(a.ctx, "UPDATE users SET State = $1 WHERE vk_id = $2", BotUser.State, BotUser.VkID)
		if err != nil {
			log.WithError(err).Error("cant set user")
			return
		}
	})
	// Run Bots Long Poll
	log.Println("Start Long Poll")
	if err := a.lp.Run(); err != nil {
		log.Fatal(err)
	}
	return nil
}

const (
	sysHTTPDefaultTimeout = 5 * time.Minute
)

func InitSysRoutes(ctx context.Context) error {

	mux := http.NewServeMux()
	{
		mux.HandleFunc("/ready", ReadyHandler)
		mux.HandleFunc("/live", LiveHandler)
	}

	port := "8080"

	s := &http.Server{
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: sysHTTPDefaultTimeout,
		ReadTimeout:  sysHTTPDefaultTimeout,
		IdleTimeout:  sysHTTPDefaultTimeout,
		Handler:      mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		fmt.Println(err)
	}
	return err
}

func ReadyHandler(w http.ResponseWriter, _ *http.Request) {
	httpStatus := http.StatusOK
	w.WriteHeader(httpStatus)
	enc := json.NewEncoder(w)
	_ = enc.Encode(map[string]bool{
		"ready": true,
	})
}

func LiveHandler(w http.ResponseWriter, _ *http.Request) {
	httpStatus := http.StatusOK
	w.WriteHeader(httpStatus)
	enc := json.NewEncoder(w)
	_ = enc.Encode(map[string]bool{
		"live": true,
	})
}
