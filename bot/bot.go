package bot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/shreya-hegdee/shreya-discordbot/config"
)

var BotID string
var jokes = []string{
	"Why don't scientists trust atoms? Because they make up everything!",
	"What do you call fake spaghetti? An impasta!",
	"I told my wife she was drawing her eyebrows too high. She looked surprised.",
}

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func sendRandomResponse(s *discordgo.Session, channelID string, responses []string) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(responses))
	response := responses[randomIndex]

	_, _ = s.ChannelMessageSend(channelID, response)
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	switch m.Content {
	case "Hi":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hellooo!")
	case "tell me time":
		currentTime := time.Now().Format("15:04:05")
		_, _ = s.ChannelMessageSend(m.ChannelID, "The current time is: "+currentTime)
	case "tell me date":
		currentDate := time.Now().Format("2006-01-02")
		_, _ = s.ChannelMessageSend(m.ChannelID, "The current date is: "+currentDate)
	case "what is your name":
		_, _ = s.ChannelMessageSend(m.ChannelID, "My name is Ping-ping. Nice to meet you!")
	case "who is your creator":
		_, _ = s.ChannelMessageSend(m.ChannelID, "I was created by Shreya and Apeksha.")
	case "tell me a joke":
		sendRandomResponse(s, m.ChannelID, jokes)
	}
}
