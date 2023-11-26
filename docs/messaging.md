### Sending a message to a chat

More useful example. This request requires authentication and sends a message to a chat.

```shell
curl -X POST https://api.teamin.pro/ \
  -H 'Content-Type: application/json' \
  -d '{"token":"<TOKEN>","actions":[{"sendMessageRequest":{"chatId":"<CHAT_ID>","text":"<MESSAGE_TEXT>"}}]}'
```
