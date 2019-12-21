# Statusoli 🤑

Statusoli pings you on telegram about Build Updates and Star Updates to your Telegram account by a Bot you can create.

Since it is your repo and it should only be limited to you. Thats why you need to give your chat id to the bot. You can either give your **personal chat id** or a **channel chat id** and add the bot to it. 

You can add these details to the Repository Secrets by going to `<repo>/settings/secrets/`

## Notifications
Since I hail from Kerela, and I like memes, Most of the notification text is in Malayalam like if it passed or failed.
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
The Output will be as ![](/op1.png)

- You can use the Star notification at the branch `star` of this repo
```yml
    - name: Statusoli
      uses: athul/statusoli@star
      if: always()
      with:
        chat: ${{ secrets.chat }}
        token: ${{ secrets.token }}
        status: ${{ job.status }}
```
Add this to your workflow file and voila....

-----

Don't forget to chat with your bot first by `/start`
