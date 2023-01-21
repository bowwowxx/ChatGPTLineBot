package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	gpt "github.com/PullRequestInc/go-gpt3"
	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	if err != nil {
		panic("Get ENV Error")
	}

	log.Println(
		"ChannelSecret:", os.Getenv("ChannelSecret"),
		"ChannelAccessToken:", os.Getenv("ChannelAccessToken"),
		"OpenApiKey:", os.Getenv("OpenApiKey"),
	)

	http.HandleFunc("/callback", callbackHandler)
	port := "8080"
	addr := fmt.Sprintf(":%s", port)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic("http listen error!")
	}

}

func GetResponse(client gpt.Client, ctx context.Context, question string) string {
	resp, err := client.CompletionWithEngine(ctx, gpt.TextDavinci003Engine, gpt.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:        gpt.IntPtr(3000),
		Temperature:      gpt.Float32Ptr(0.9),
		TopP:             gpt.Float32Ptr(1),
		FrequencyPenalty: float32(0),
		PresencePenalty:  float32(0.6),
	})

	if err != nil {
		log.Println("Get Open AI Response Error: ", err)
	}

	answer := resp.Choices[0].Text
	answer = strings.TrimSpace(answer)
	return answer
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			http.Error(w, "Invalid signature", http.StatusBadRequest)
		} else {
			http.Error(w, "Error handling request", http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		if event.Type != linebot.EventTypeMessage {
			continue
		}
		message, ok := event.Message.(*linebot.TextMessage)
		if !ok {
			continue
		}
		prefix := "pikachu"
		if !strings.HasPrefix(message.Text, prefix) {
			log.Println("message: ", message.Text)
			continue
		}
		question := strings.TrimPrefix(message.Text, prefix)
		question = strings.TrimSpace(question)
		if question == "" {
			continue
		}
		log.Println("pikachu:", question)
		if strings.HasPrefix(question, "88") {
			_, err := bot.LeaveRoom(event.Source.RoomID).Do()
			if err != nil {
				log.Print(err)
			}
		}
		apiKey := os.Getenv("OpenApiKey")
		if apiKey == "" {
			panic("Missing API KEY")
		}
		log.Println("apiKey: ", apiKey)
		ctx := context.Background()
		client := gpt.NewClient(apiKey)
		answer := GetResponse(client, ctx, question)
		log.Println("pikachu say:", answer)
		if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answer)).Do(); err != nil {
			log.Print(err)
		}
	}

}
