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
	}

	var (
		// inputs provided by Github Actions runtime
		// we should define them in action.yml
		token    = os.Getenv("INPUT_TOKEN")
		chat     = os.Getenv("INPUT_CHAT")
		status   = os.Getenv("INPUT_STATUS")
		stars    = os.Getenv("STARGAZERS")
		forks    = os.Getenv("FORKERS")
		ititle   = os.Getenv("IU_TITLE")
		inum     = os.Getenv("IU_NUM")
		ibody    = os.Getenv("IU_BODY")
		icomment = os.Getenv("IU_COM")
		prstate  = os.Getenv("PR_STATE")
		prnum    = os.Getenv("PR_NUM")
		prtitle  = os.Getenv("PR_TITLE")
		prbody   = os.Getenv("PR_BODY")

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

		Issue Number:  %s

		Issue Title: 	%s

		Issue Body:		%s

		Link:		[%s](%s)

		Triggered by:   *%s* 
		
		Event:		 *%s*
		
		`, icon, text, status, repo, inum, ititle, ibody, workflow, link, person, event)
	}
	if event == "issue_comment" {
		msg = fmt.Sprintf(`
		%s | %s  
		
		Status: 	*%s*
		Repository:  	 %s 

		Issue Number:  %s

		Issue Title: 	%s

		Comment:		%s

		Link:		[%s](%s)

		Triggered by:   *%s* 

		Event:		 *%s*
		
		`, icon, text, status, repo, inum, ititle, icomment, workflow, link, person, event)
	}

	if event == "pull_request" {
		msg = fmt.Sprintf(`
		%s | %s  
		
		Status: 	*%s*

		Repository:  	 %s 

		PR Number:  %s 	%s

		PR Title: 	%s

		PR Body:		%s

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

		Stars:		%s

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
