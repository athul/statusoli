# TELEWIRE is a Faster implemetation of Statusoli. Check that out at https://github.com/athul/telewire

----
# Statusoli 🤑

**_Rather than putting the tag, use the latest commit id to improve security_**  
eg:  `uses: athul/statusoli@37a2904`

Statusoli sends you on telegram about Repo Updates to your Telegram account by a Bot you can create.

Create a chatbot with **botfather** bot in telegram. Get your chat id by speaking to **jsondumpbot** in telegram.

Since it is your repo and it should only be limited to you. Thats why you need to give your chat id to the bot. You can either give your **personal chat id** or a **channel chat id** and add the bot to it. 

You can add these details to the Repository Secrets by going to `<repo>/settings/secrets/`

## Notifications
- You can use the simple notifier at the master branch or the release tag like 
```yml
    - name: Statusoli
      uses: athul/statusoli@master
      if: always()
      with:
        chat: ${{ secrets.chat }}
        token: ${{ secrets.token }}
        status: ${{ job.status }}

```
The `chat` is the chat id/channel id and you can get that by talking to the json dump bot. The `token` is the bot's API token and you can create a bot by speaking to Botfather bot in Telegram.    
The Output will be as 
- Push   
![](/op1.png)
- Stars   
![](/op2.png)
- PR Opened    
![](/op5.png)
- Issues   
![](/op3.png)
- Issue/ PR Comments   
![](/op4.png)


-----

Don't forget to chat with your bot first by `/start`   

Actions will only trigger on what you want to trigger. You might want to define all the triggers first. You can refer the workflow file of this repo for better guidance. or like this
```yml
name: Build and Notify
on:
  push:
  pull_request:
    types: [opened,closed]
  issues:
    types: [opened, closed, reopened]
  issue_comment:
    types: [created]
  watch:
    types: [started]
jobs:
```
