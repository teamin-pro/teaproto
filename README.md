# gtproto

Extensible protocol for GTNode and client. Based on protobuf over HTTP(s) and WebSocket. GRPC is not used.

To use the gtproto protocol, start by referring to the [actions.proto](https://github.com/gamarjoba-team/gtproto/blob/main/actions.proto) file.

### Synchronous workflow (for actions)
To initiate a synchronous workflow for actions:

 - The client must send a ClientRequest message via a POST HTTP request to the endpoint https://api.gamarjoba.team/.
 - The server responds with a `ServerResponse` message.

### Asynchronous workflow (for events)
To initiate an asynchronous workflow for events:
- The client must create a WebSocket connection to the endpoint https://api.gamarjoba.team/ws.
- The client sends a `WebsocketAuthRequest` message.
- The server starts sending `WebsocketEvent` messages.

### Notes
 - Timestamps are encoded in uint64 format and represent milliseconds in UTC.
 - All IDs are represented as strings using the uuid format.

## Some examples of interaction

### Authentication

#### By password
`StateResponse.auth_methods` must contain `EMAIL_PASSWORD_AUTH_METHOD`
 
 - send `EmailPasswordAuthResponse` with `email` and `password` fields

##### By one-time code
`StateResponse.auth_methods` must contain `EMAIL_CODE_AUTH_METHOD`

 - send `EmailCodeAuthRequest` with `email` field
 - server will send email with code to `email` address
 - send `EmailCodeAuthConfirmRequest` with `email` and `code` fields

## JSON support
For command line tools, it is easier to use JSON instead of protobuf. To use JSON, simply set `Content-Type` 
header to `application/json` and send the JSON representation of the protobuf message. 

If content type is omitted and request starts with `{` (without leading spaces) it will be treated as JSON.

### Getting server info
This is simplest request. It does not require authentication and returns server info.

```shell
curl -X POST https://api.gamarjoba.team/ \
  -H 'Content-Type: application/json' \
  -d '{"stateRequest":{}}'
```

If request is JSON, response also will be JSON:
```json
{
  "stateResponse": {
    "now": "1679507631905",
    "maxResultsOnPage": 100,
    "maxGroupTitleLength": 40,
    "maxUsernameLength": 25,
    "maxMessageLength": 5000,
    "maxFileSize": 524288000,
    "maxUploadsForMessage": 10,
    "minIconSize": 256
  }
}
```

### Sending a message to a chat

More useful example. This request requires authentication and sends a message to a chat.

```shell
curl -X POST https://api.gamarjoba.team/ \
  -H 'Content-Type: application/json' \
  -d '{"token": "<TOKEN>", "sendMessageRequest": {"chatId": "<CHAT_ID>", "text": "<MESSAGE_TEXT>"}}'
```

## FAQ

### Why use own protocol not XMPP?
Firstly, creating own protocol can offer more flexibility and customization options for the app developers. They can design the protocol to specifically fit the needs of their app and its features, which may not be possible with a standardized protocol like XMPP. This can allow for faster innovation and development of new features.

Secondly, using a proprietary protocol can provide greater control over the security and privacy of the messaging app. By designing their own protocol, the app developers can implement security measures that may not be possible with a standardized protocol. This can help to prevent unauthorized access to user data, protect user privacy, and prevent malicious attacks.
