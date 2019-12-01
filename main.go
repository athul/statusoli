package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yanzay/tbot/v2"
)

func main() {
	icons := map[string]string{
		"failure":   "❗️❗️❗️Eda Mone Nee Pettu ",
		"cancelled": "⚠️⚠️⚠️Enthupatti Babymole?",
		"success":   "💯💯👍👍Adipoli Monuse",
	}
	var (
		token  = os.Getenv("TG_TOKEN")
		chat   = os.Getenv("CHAT_TOKEN")
		status = os.Getenv("INPUT_STATUS")

		workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
		commit   = os.Getenv("GITHUB_SHA")
	)
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")
	icon := icons[strings.ToLower(status)]
	link := fmt.Sprintf("https://github.com/%s/commit/%s/checks", repo, commit)

	msg := fmt.Sprintf(`%s*%s*: %s ([%s](%s))    **%s**`, icon, status, repo, workflow, link, texts)

	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
