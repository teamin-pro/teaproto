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

However, using a proprietary protocol can also have drawbacks. For example, it can limit interoperability with other messaging apps, as other apps may not be able to communicate using the same protocol. This can make it harder for users to switch between different messaging apps or communicate with users on other apps. Additionally, a proprietary protocol may require more resources to develop and maintain, as the app developers are solely responsible for ensuring the protocol is secure and up-to-date.
