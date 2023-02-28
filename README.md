# gtproto
Protocol specification.

## How to use
Start from [actions.proto](https://github.com/gamarjoba-team/gtproto/blob/main/actions.proto).

### Synchronous workflow (for actions)
 - client send POST http request to `https://api.gamarjoba.team/` with message `ClientRequest`
 - server respond with message `ServerResponse`.

### Asynchronous workflow (for events)
 - client create a websocket connection to `https://api.gamarjoba.team/ws`;
 - client send message `WebsocketAuthRequest`;
 - server start sending `WebsocketEvent` messages.

### Notes
 - timestamps encoded in `uint64` with milliseconds in UTC;
 - all ids are string representation of `uuid`.

## Why use own protocol not XMPP?
Firstly, creating own protocol can offer more flexibility and customization options for the app developers. They can design the protocol to specifically fit the needs of their app and its features, which may not be possible with a standardized protocol like XMPP. This can allow for faster innovation and development of new features.

Secondly, using a proprietary protocol can provide greater control over the security and privacy of the messaging app. By designing their own protocol, the app developers can implement security measures that may not be possible with a standardized protocol. This can help to prevent unauthorized access to user data, protect user privacy, and prevent malicious attacks.
