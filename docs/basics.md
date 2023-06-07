# Basic usage

## Synchronous workflow (for actions)

To initiate a synchronous workflow for actions:

 * The client must send a `ClientRequest` message via a POST HTTP request to the endpoint [https://api.gamarjoba.team/](https://api.gamarjoba.team/)
 * The server responds with a `ServerResponse` message.

## Asynchronous workflow (for events)

To initiate an asynchronous workflow for events:

 * The client must create a WebSocket connection to the endpoint https://api.gamarjoba.team/ws.
 * The client sends a `WebsocketAuthRequest` message.
 * The server starts sending `WebsocketEvent` messages.

## Notes
 * Timestamps are encoded in `uint64` format and represent milliseconds in UTC.
 * All IDs are represented as strings using the uuid format.

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
