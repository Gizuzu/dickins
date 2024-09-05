package main

import (
	"context"
	"dickins/commands"
	"dickins/database"
	"log"
	"os"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/events"
	"github.com/SevereCloud/vksdk/v3/longpoll-bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	database.Connect()
	defer database.Disconnect()

	vk := api.NewVK(os.Getenv("VK_TOKEN"))

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal("Failed to get group ID: ", err)
	}

	lp, err := longpoll.NewLongPoll(vk, group.Groups[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		commands.HandleCommand(vk, obj)
	})

	log.Printf("Bot polling for group \"%s\" started\n", group.Groups[0].Name)
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
