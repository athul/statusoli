# Statusoli 🤑

Statusoli sends you on telegram about Repo Updates to your Telegram account by a Bot you can create.

Create a chatbot with **botfather** bot in telegram. Get your chat id by speaking to **jsondumpbot** in telegram.

Since it is your repo and it should only be limited to you. Thats why you need to give your chat id to the bot. You can either give your **personal chat id** or a **channel chat id** and add the bot to it. 

You can add these details to the Repository Secrets by going to `<repo>/settings/secrets/`

## Notifications
Since I hail from Kerela, and I like memes, Most of the notification text is in Malayalam like if it passed or failed.
- You can use the simple notifier at the master branch or the release tag like 
```yml
    - name: Statusoli
      uses: athul/statusoli@master
      env:
        STARGAZERS: ${{ github.event.repository.stargazers_count }}
        FORKERS: ${{ github.event.repository.forks_count }}
        IU_TITLE: ${{ github.event.issue.title }}
        IU_NUM: ${{ github.event.issue.number }}
        IU_ACTOR: ${{ github.event.issue.user.login }}
        IU_BODY: ${{ github.event.issue.body }}
        IU_COM: ${{github.event.comment.body}}
        PR_STATE: ${{ github.event.action }}
        PR_NUM: ${{ github.event.number }}
        PR_TITLE: ${{ github.event.pull_request.title }}
        PR_BODY: ${{ github.event.pull_request.body }}
      if: always()
      with:
        chat: ${{ secrets.chat }}
        token: ${{ secrets.token }}
        status: ${{ job.status }}

```
The `chat` is the chat id/channel id and you can get that by talking to the json dump bot. The `token` is the bot's API token and you can create a bot by speaking to Botfather bot in Telegram.    
The Output will be as ![](/op1.png)


-----

Don't forget to chat with your bot first by `/start`
