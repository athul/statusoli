
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
		"success":   "🌟✨",
	}
	texts:=map[string]string{
		"failure":   "എടാ മോനെ നീ പെട്ടു",
		"cancelled": "എന്തുപറ്റിയെടാ ഉവ്വേ ?",
		"success":   "എടാ മോനെ ആരോ ⭐️⭐️⭐️ ചെയ്തു  ",
	}
		
	var (
		// inputs provided by Github Actions runtime
		// we should define them in action.yml
		token  = os.Getenv("INPUT_TOKEN")
		chat   = os.Getenv("INPUT_CHAT")
		status = os.Getenv("INPUT_STATUS")
		//now=time.now()
		// github environment context
		//workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
		//commit   = os.Getenv("GITHUB_SHA")
		person	 =os.Getenv("GITHUB_ACTOR")
		//event	 =os.Getenv("GITHUB_EVENT_NAME")
	)

	// Create Telegram client using token
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	icon := icons[strings.ToLower(status)]
	text:=texts[strings.ToLower(status)]// which icon to use?
	//link := fmt.Sprintf("https://github.com/%s/commit/%s/checks", repo, commit)
	// Prepare message to send
	msg := fmt.Sprintf(`
	%s
	%s  
	Person: 	[%s](https://github.com/%s) 
	Repository 	*%s*
	`, icon, text,person,person,repo )

	// Send to chat using Markdown format
	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
