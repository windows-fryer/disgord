package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	discordgo "disgord.dev/disgord"
)

func main() {
	godotenv.Load(".env")

	application := discordgo.BuildDiscordApplicationToken(os.Getenv("BOT_TOKEN"))

	user, err := application.GetUser()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(user)
}
