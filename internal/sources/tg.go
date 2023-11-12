package sources

import (
	"context"
	"fmt"
	"gonstructor/internal/domain"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TG struct {
	responser *TGResponder
	BotAPI    *tgbotapi.BotAPI
	reqeustCh chan domain.Request
}

type TGConfig struct {
	Token string `json:"token" bson:"token" yaml:"token" mapstructure:"token"`
}

type TGResponder struct {
	BotAPI *tgbotapi.BotAPI
}

const resourceCode = "tg"

func NewTGResource(conf *TGConfig, reqeustCh chan domain.Request) (TG, *TGResponder, error) {
	api, err := tgbotapi.NewBotAPI(conf.Token)
	responser := &TGResponder{
		BotAPI: api,
	}

	return TG{
		BotAPI:    api,
		reqeustCh: reqeustCh,
	}, responser, err
}

func (responder TGResponder) Send(ctx context.Context, s *state.State, response response.Response) error {
	userId, err := strconv.Atoi(s.UserID)

	if err != nil {
		return err
	}

	for _, message := range response.Messages {
		msg := tgbotapi.NewMessage(int64(userId), message.Text)

		if len(message.Buttons) > 0 {
			var btns [][]tgbotapi.InlineKeyboardButton
			for _, v := range message.Buttons {
				var line []tgbotapi.InlineKeyboardButton
				for _, b := range v {
					fmt.Println(b.Target.Resource)
					line = append(line, tgbotapi.NewInlineKeyboardButtonData(b.Display, b.Target.Resource))
				}
				btns = append(btns, line)
			}

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(btns...)
			msg.ParseMode = "Markdown"
		}

		_, err = responder.BotAPI.Send(msg)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

func (source *TG) Listen() {

	u := tgbotapi.NewUpdate(1)
	u.Timeout = 60

	updatesChan := source.BotAPI.GetUpdatesChan(u)

	func() {
		for update := range updatesChan {
			if update.Message != nil {
				initiatorID := strconv.FormatInt(update.FromChat().ID, 10)
				req := domain.Request{
					FromID:  initiatorID,
					Source:  resourceCode,
					Command: update.Message.Text,
				}

				source.reqeustCh <- req

				continue
			}

			if update.CallbackQuery != nil {
				initiatorID := strconv.FormatInt(update.FromChat().ID, 10)
				req := domain.Request{
					FromID:  initiatorID,
					Source:  resourceCode,
					Command: update.CallbackQuery.Data,
				}

				source.reqeustCh <- req
			}
		}
	}()
}
