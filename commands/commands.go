package commands

import (
	"strings"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/events"
)

type commandhandler func(*api.VK, events.MessageNewObject)

type commandinfo struct {
	name    string
	handler commandhandler
}

var commandsList = map[string]commandinfo{
	"писюн2": {
		name:    "писюн",
		handler: dickHandler,
	},
}

var prefix = "/"

func HandleCommand(vk *api.VK, obj events.MessageNewObject) {
	if !strings.HasPrefix(obj.Message.Text, prefix) {
		return
	}

	cn := strings.Split(strings.TrimPrefix(obj.Message.Text, prefix), " ")[0]
	cmd, ok := commandsList[cn]
	if !ok {
		return
	}

	go cmd.handler(vk, obj)
}
