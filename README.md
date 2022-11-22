# gtproto
Protocol specification.

Start from [actions.proto](https://github.com/gamarjoba-team/gtproto/blob/main/actions.proto).

## Synchronous workflow (for actions)
 - client send POST http request `/api/v2` with message `ClientRequest`
 - server respond with message `ServerResponse`.

## Asynchronous workflow (for events)
 - client create a websocket connection to `/api/v2/ws`;
 - client send message `WebsocketAuthRequest`;
 - server start sending `WebsocketEvent` messages.

## Notes
 - timestamps encoded in `uint64` with milliseconds in UTC;
 - all ids are string representation of `uuid`.
