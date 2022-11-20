# gtproto
Protocol specification.

Start from [actions.proto](https://github.com/gamarjoba-team/gtproto/blob/main/actions.proto).

## Workflow
 - client send POST http request `/api/v2` with protobuf encoded message `ClientRequest`;
 - server respond with protobuf encoded message `ServerResponse`;
 - also server can broadcast `WebsocketEvent` message for each client have websocket connection to `/api/v2/ws`.

## Notes
 - timestamps encoded in `uint64` with milliseconds in UTC;
 - all ids are string representation of `uuid`.
