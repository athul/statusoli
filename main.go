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
	// Just to do it a bit fancy
	icons := map[string]string{
		"failure":   "❗️",
		"cancelled": "❕",
		"success":   "✅",
	}
	texts := map[string]string{
		"issues":        "‼️‼️‼️",
		"pull_request":  "🔃🔀⤴️🔃",
		"issue_comment": "🗣❗️🗣❗️🗣❗️🗣❗️",
		"push":          "⬆️⬆️⬆️⬆️",
		"watch":         "⭐️⭐️⭐️⭐️",
		"schedule":      "⏰⏰⏰⏰",
	}

	var (
		// inputs provided by Github Actions runtime
		// we should define them in action.yml
		token    = os.Getenv("INPUT_TOKEN")
		chat     = os.Getenv("INPUT_CHAT")
		status   = os.Getenv("INPUT_STATUS")
		stars    = os.Getenv("INPUT_STARGAZERS")
		forks    = os.Getenv("INPUT_FORKERS")
		ititle   = os.Getenv("INPUT_IU_TITLE")
		inum     = os.Getenv("INPUT_IU_NUM")
		ibody    = os.Getenv("INPUT_IU_BODY")
		icomment = os.Getenv("INPUT_IU_COM")
		prstate  = os.Getenv("INPUT_PR_STATE")
		prnum    = os.Getenv("INPUT_PR_NUM")
		prtitle  = os.Getenv("INPUT_PR_TITLE")
		prbody   = os.Getenv("INPUT_PR_BODY")

		// github environment context
		workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
		commit   = os.Getenv("GITHUB_SHA")
		person   = os.Getenv("GITHUB_ACTOR")
		event    = os.Getenv("GITHUB_EVENT_NAME")
	)

	// Create Telegram client using token
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	icon := icons[strings.ToLower(status)]
	text := texts[strings.ToLower(event)] // which icon to use?
	link := fmt.Sprintf("https://github.com/%s/commit/%s/checks", repo, commit)
	var msg string
	// Prepare message to send
	if event == "issues" {
		msg = fmt.Sprintf(`
		%s | %s 
		
		Status: 	*%s*

		Repository:  	 %s 

		Issue Number:  %s	| [%s]

		Issue Title: 	%s

		Issue Body:		*%s*

		Link:		[%s](%s)

		Triggered by:   *%s* 
		
		Event:		 *%s*
		
		`, icon, text, status, repo, inum, prstate, ititle, ibody, workflow, link, person, event)
	}
	if event == "schedule" {
		msg = fmt.Sprintf(`
		%s | %s 
		
		Status: 	*%s*

		Repository:  	 %s 

		*This was run on Schedule*

		Link:		[%s](%s)

		Triggered by:   *%s* 
		
		Event:		 *%s*
		
		`, icon, text, status, repo, workflow, link, person, event)
	}
	if event == "issue_comment" {
		msg = fmt.Sprintf(`
		%s | %s  
		
		Status: 	*%s*

		Repository:  	 %s 

		Issue Number:  %s	| [%s]

		Issue Title: 	%s

		Comment:		*%s*

		Link:		[%s](%s)

		Triggered by:   *%s* 

		Event:		 *%s*
		
		`, icon, text, status, repo, inum, prstate, ititle, icomment, workflow, link, person, event)
	}

	if event == "pull_request" {
		msg = fmt.Sprintf(`
		%s | %s  
		
		Status: 	*%s*

		Repository:  	 %s 

		PR Number:  %s	| [%s]

		PR Title: 	%s

		PR Body:		*%s*

		Link:		[%s](%s)

		Triggered by:   *%s* 

		Event:		 *%s*
		
		`, icon, text, status, repo, prnum, prstate, prtitle, prbody, workflow, link, person, event)
	}

	if event == "watch" {
		msg = fmt.Sprintf(`
		%s | %s 

		Status: 	*%s*

		Repository:  	 %s 

		Stars:		*%s*

		Forks:		%s

		Link:		[%s](%s)

		Triggered by:   *%s* 

		Event:		 *%s*
		
		`, icon, text, status, repo, stars, forks, workflow, link, person, event)
	}
	if event == "push" {
		msg = fmt.Sprintf(`
		%s | %s 
		
		Status: 	*%s*

		Repository:  	 %s 

		Link:		[%s](%s)

		Triggered by:   *%s* 

		Event:		 *%s*
		
		`, icon, text, status, repo, workflow, link, person, event)
	}

	// Send to chat using Markdown format
	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
