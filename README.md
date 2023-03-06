# gtproto
Protocol specification.

## How to use
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

## Why use own protocol not XMPP?
Firstly, creating own protocol can offer more flexibility and customization options for the app developers. They can design the protocol to specifically fit the needs of their app and its features, which may not be possible with a standardized protocol like XMPP. This can allow for faster innovation and development of new features.

Secondly, using a proprietary protocol can provide greater control over the security and privacy of the messaging app. By designing their own protocol, the app developers can implement security measures that may not be possible with a standardized protocol. This can help to prevent unauthorized access to user data, protect user privacy, and prevent malicious attacks.
