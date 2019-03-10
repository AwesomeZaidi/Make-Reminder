package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

const helpMessage = "type in '@weather <whats the weather in> <location>'"

/*
   CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
   initiating the socket connection and returning the client.
   DO NOT EDIT THIS FUNCTION. This is a fully complete implementation.
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	fmt.Println("yo")
	api := slack.New(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	fmt.Println("yoooo")

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
	fmt.Println(strings.ToLower(args[0]))

	// slackClient.SendMessage(slackClient.NewOutgoingMessage("Lemme pull up this weather for you, gimme a second, i'll GO get, get it! Hah, what you thought I only told weather, nah I got jokes too, you know,for them cloudy days </3", slackChannel))
	dialogObj := getWeather(strings.ToLower(args[0]))
	slackClient.SendMessage(slackClient.NewOutgoingMessage(string(dialogObj.Weather[0].Description), slackChannel))

	// switch strings.ToLower(args[0]) {
	// case "chicago":
	// 	slackClient.SendMessage(slackClient.NewOutgoingMessage("Lemme pull up this weather for you, gimme a second, i'll GO get, get it! Hah, what you thought I only told weather, nah I got jokes too, you know,for them cloudy days </3", slackChannel))
	// 	// slackClient.SendMessage(slackClient.NewOutgoingMessage(strings.Join(args[1:], " "), slackChannel))
	// 	dialogObj := getWeather("chicago")
	// 	slackClient.SendMessage(slackClient.NewOutgoingMessage(string(dialogObj.Weather[0].Description), slackChannel))
	// 	break
	// // case "scan francisco":
	// // 	slackClient.SendMessage(slackClient.NewOutgoingMessage("Lemme pull up this weather for you, gimme a second, i'll GO get, get it! Hah, what you thought I only told weather, nah I got jokes too, you know,for them cloudy days </3", slackChannel))
	// // 	slackClient.SendMessage(slackClient.NewOutgoingMessage(getWeather("san francisco"), slackChannel))
	// // 	break
	// default:
	// 	slackClient.SendMessage(slackClient.NewOutgoingMessage("Lemme pull up this weather for you, gimme a second, i'll GO get, get it! Hah, what you thought I only told weather, nah I got jokes too, you know,for them cloudy days </3", slackChannel))
	// 	// slackClient.SendMessage(slackClient.NewOutgoingMessage(getWeather("san francisco"), slackChannel))
	// 	return
	// }

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

type WeatherDesc struct {
	Description string `json:"description"`
}

type Dialog struct {
	Weather []WeatherDesc `json:"weather"`
}

func getWeather(location string) Dialog {
	fmt.Println("hereee")
	// make request to get my weather

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + location + "&APPID=" + os.Getenv("API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	// read all into body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	// parse json obj
	// // var dialog map[string]interface{}
	// type Main struct {
	// 	humidity int32
	// }

	// type Dialog struct {
	// 	main Main
	// }
	// type AutoGenerated struct {
	// 	Coord struct {
	// 		Lon `json:"lon"`
	// 		Lat `json:"lat"`
	// 	} `json:"coord"`
	// 	Weather []struct {
	// 		ID          `json:"id"`
	// 		Main        `json:"main"`
	// 		Description `json:"description"`
	// 		Icon        `json:"icon"`
	// 	} `json:"weather"`
	// 	Base string `json:"base"`
	// 	Main struct {
	// 		Temp     `json:"temp"`
	// 		Pressure `json:"pressure"`
	// 		Humidity `json:"humidity"`
	// 		TempMin  `json:"temp_min"`
	// 		TempMax  `json:"temp_max"`
	// 	} `json:"main"`
	// 	Visibility int `json:"visibility"`
	// 	Wind       struct {
	// 		Speed `json:"speed"`
	// 		Deg   `json:"deg"`
	// 		Gust  `json:"gust"`
	// 	} `json:"wind"`
	// 	Snow struct {
	// 		OneH `json:"1h"`
	// 	} `json:"snow"`
	// 	Clouds struct {
	// 		All `json:"all"`
	// 	} `json:"clouds"`
	// 	Dt  int `json:"dt"`
	// 	Sys struct {
	// 		Type    `json:"type"`
	// 		ID      `json:"id"`
	// 		Message `json:"message"`
	// 		Country `json:"country"`
	// 		Sunrise `json:"sunrise"`
	// 		Sunset  `json:"sunset"`
	// 	} `json:"sys"`
	// 	ID   int    `json:"id"`
	// 	Name string `json:"name"`
	// 	Cod  int    `json:"cod"`
	// }

	var weatherObj Dialog

	err = json.Unmarshal(body, &weatherObj)
	fmt.Println(weatherObj)

	if err != nil {
		log.Fatalf("Dialog parsing error: %s", err)
		return Dialog{}
	}
	// fmt.Println(dialog.main.humidity)
	return weatherObj
	// type response struct {
	// 	Message string `json:"coord"`
	// }
	// data := response{}

	// // marshal into struct
	// json.Unmarshal(body, &data)

	// // return data
	// fmt.Println(data)
	// return data.Message
}
