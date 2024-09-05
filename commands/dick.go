package commands

import (
	"dickins/database"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/SevereCloud/vksdk/v3/events"
)

func dickHandler(vk *api.VK, obj events.MessageNewObject) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.PeerID(obj.Message.PeerID)

	dick := database.GetUserDick(vk, obj)

	if dick.Id == 0 {
		b.Message("Произошла ошибка, попробуй позже :(")

		_, err := vk.MessagesSend(b.Params)
		if err != nil {
			log.Println("Failed to send message to peer: ", err)
		}

		return
	}

	if (dick.IssuedAt != nil) && (time.Now().UTC().Add(3*time.Hour).UnixMilli() >= dick.IssuedAt.UTC().Add(3*time.Hour).UnixMilli()) {
		b.Message("Сегодня ты уже мерил свою писюнчик, давай завтра")

		_, err := vk.MessagesSend(b.Params)
		if err != nil {
			log.Println("Failed to send message to peer: ", err)
		}

		return
	}

	dickAppender := rand.IntN(16)

	b.Message(fmt.Sprintf("Ух, твой писюн увеличился на %d см!\nТеперь у тебя писюн %d см.", dickAppender, dick.DickSize+int64(dickAppender)))
	database.UpdateUserDickSize(dick.Id, dick.DickSize+int64(dickAppender))

	_, err := vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to send message to peer: ", err)
	}
}
