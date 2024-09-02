package main

import (
	"dickins/database"
	"log"
	"os"

	"github.com/SevereCloud/vksdk/v3/api"
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
		log.Fatal(err)
	}

	lp, err := longpoll.NewLongPoll(vk, group.Groups[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bot polling for group \"%s\" started\n", group.Groups[0].Name)
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
