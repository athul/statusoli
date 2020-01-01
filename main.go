
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	//"time"
	"github.com/yanzay/tbot/v2"
)

func main() {
	// Just to do it a bit fancy
	icons := map[string]string{
		"failure":   "❗️",
		"cancelled": "❕",
		"success":   "✅",
	}
		
	var (
		// inputs provided by Github Actions runtime
		// we should define them in action.yml
		token  = os.Getenv("INPUT_TOKEN")
		chat   = os.Getenv("INPUT_CHAT")
		status = os.Getenv("INPUT_STATUS")
		//now=time.now()
		// github environment context
		workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
		//commit   = os.Getenv("GITHUB_SHA")
		person	 =os.Getenv("GITHUB_ACTOR")
		event	 =os.Getenv("GITHUB_EVENT_NAME")
	)

	// Create Telegram client using token
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	icon := icons[strings.ToLower(status)]
	//text:=texts[strings.ToLower(status)]// which icon to use?
	link := fmt.Sprintf("https://github.com/%s", repo,)
	// Prepare message to send
	msg := fmt.Sprintf(`
	Run %s
	
	-----New %s [ope/clo/reop/assi]----- 
	
	Person: 	[%s](https://github.com/%s) 
	
	Repository: 	[%s](%s)

	Workflow:	*%s*
	
	`, icon,event, person,person,repo,link,workflow )

	// Send to chat using Markdown format
	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
