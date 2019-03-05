package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

const helpMessage = "type in '@make-reminder <whats the weather in> <location>'"

/*
   CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
   initiating the socket connection and returning the client.
   DO NOT EDIT THIS FUNCTION. This is a fully complete implementation.
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

/*
   RespondToEvents waits for messages on the Slack client's incomingEvents channel,
   and sends a response when it detects the bot has been tagged in a message with @<botTag>.

   EDIT THIS FUNCTION IN THE SPACE INDICATED ONLY!
*/
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			message := strings.Replace(ev.Msg.Text, botTagString, "", -1)

			// TODO: Make your bot do more than respond to a help command. See notes below.
			// Make changes below this line and add additional funcs to support your bot's functionality.
			// sendHelp is provided as a simple example. Your team may want to call a free external API
			// in a function called sendResponse that you'd create below the definition of sendHelp,
			// and call in this context to ensure execution when the bot receives an event.

			// START SLACKBOT CUSTOM CODE
			// ===============================================================
			sendResponse(slackClient, message, ev.Channel)
			sendHelp(slackClient, message, ev.Channel)
			// ===============================================================
			// END SLACKBOT CUSTOM CODE

		}
	}
}

// sendHelp is a working help message, for reference.
func sendHelp(slackClient *slack.RTM, message, slackChannel string) {
	if strings.ToLower(message) != "help" {
		return
	}
	slackClient.SendMessage(slackClient.NewOutgoingMessage(helpMessage, slackChannel))
}

// sendResponse is NOT unimplemented --- write code in the function body to complete!

func sendResponse(slackClient *slack.RTM, message, slackChannel string) {
	args := strings.Split(message, " ")

	switch strings.ToLower(args[3]) {
	case "Chicago":
		slackClient.SendMessage(slackClient.NewOutgoingMessage(strings.Join(args[1:], " "), slackChannel))
		slackClient.SendMessage(slackClient.NewOutgoingMessage(getWeather("Chicago"), slackChannel))
		break
	case "San Francisco":
		slackClient.SendMessage(slackClient.NewOutgoingMessage("Lemme pull up this weather for you, gimme a second, i'll GO get, get it! Hah, what you thought I only told weather, nah I got jokes too, you know,for them cloudy days </3", slackChannel))
		slackClient.SendMessage(slackClient.NewOutgoingMessage(getWeather("San Francisc"), slackChannel))
		break
	default:
		slackClient.SendMessage(slackClient.NewOutgoingMessage("Lemme pull up this weather for you, gimme a second, i'll GO get, get it! Hah, what you thought I only told weather, nah I got jokes too, you know,for them cloudy days </3", slackChannel))
		slackClient.SendMessage(slackClient.NewOutgoingMessage(getWeather("San Francisco"), slackChannel))
		return
	}

	println("[RECEIVED] sendResponse:", args[0])

	// START SLACKBOT CUSTOM CODE
	// ===============================================================
	// TODO:
	//      1. Implement sendResponse for one or more of your custom Slackbot commands.
	//         You could call an external API here, or create your own string response. Anything goes!
	//      2. STRETCH: Write a goroutine that calls an external API based on the data received in this function.
	// ===============================================================
	// END SLACKBOT CUSTOM CODE
}

func getWeather(location string) string {

	// make request to get my weather
	res, _ := http.Get(fmt.Sprintf("api.openweathermap.org/data/2.5/weather?q=" + location + "&APPID=" + os.Getenv("API_KEY")))

	// read all into body
	body, _ := ioutil.ReadAll(res.Body)

	//
	// declare struct
	// type response struct {
	// 	Image string `json:"img"`
	// }

	type response struct {
		Message string `json:"message"`
	}
	data := response{}

	// marshal into struct
	json.Unmarshal(body, &data)

	// return image url
	return data.Message

}
