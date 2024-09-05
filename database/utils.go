package database

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/SevereCloud/vksdk/v3/events"
	"github.com/SevereCloud/vksdk/v3/object"
)

type DickInfo struct {
	Id        int
	Username  string
	DickSize  int64
	IssuedAt  *time.Time
	CreatedAt *time.Time
}

func GetUserDick(vk *api.VK, obj events.MessageNewObject) DickInfo {
	var dick DickInfo

	err := Pool.QueryRow(context.Background(), "SELECT id, username, dick_size, issued_at, created_at FROM dicks WHERE peer_id=$1 AND user_id=$2", obj.Message.PeerID, obj.Message.FromID).
		Scan(
			&dick.Id,
			&dick.Username,
			&dick.DickSize,
			&dick.IssuedAt,
			&dick.CreatedAt,
		)
	if err != nil {
		log.Println("Failed to get user, creating new: ", err)

		b := params.NewUsersGetBuilder()
		b.UserIDs([]string{strconv.Itoa(obj.Message.FromID)})
		b.Lang(object.LangRU)

		ui, err := vk.UsersGet(b.Params)
		if err != nil {
			log.Println("Failed to get user info: ", err)
			return dick
		}

		u := ui[0]

		err = Pool.QueryRow(
			context.Background(),
			"INSERT INTO dicks (username, user_id, peer_id) VALUES ($1, $2, $3) RETURNING id, username, dick_size, issued_at, created_at",
			u.LastName+" "+u.FirstName,
			obj.Message.FromID,
			obj.Message.PeerID,
		).
			Scan(
				&dick.Id,
				&dick.Username,
				&dick.DickSize,
				&dick.IssuedAt,
				&dick.CreatedAt,
			)
		if err != nil {
			log.Println("Failed to create user: ", err)
			return dick
		}
	}

	return dick
}

func UpdateUserDickSize(id int, dickSize int64) DickInfo {
	var dick DickInfo

	Pool.QueryRow(context.Background(), "UPDATE dicks SET dick_size=$1, issued_at=$2 WHERE id=$3 RETURNING id, username, dick_size, issued_at, created_at", dickSize, time.Now(), id).
		Scan(
			&dick.Id,
			&dick.Username,
			&dick.DickSize,
			&dick.IssuedAt,
			&dick.CreatedAt,
		)

	return dick
}
