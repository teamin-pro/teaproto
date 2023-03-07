## v1.4.0
 - tasks api first version

## v1.2.0
 - `OnlineStatus` was added

## v1.1.2
 - `ChatListRequest.min_updated_at` was added

## v1.1.1
 - `Badge` was renamed to `GlobalBadge`
 - `ChatBadge.badge` was renamed to `ChatBadge.counter`, and `ChatBadge.updated_at` added
 - `Chat.badge` and `ViewMessageResponse.badge` now returns `ChatBadge` not `uint32`

## v1.1.0
 - global badges

## v1.0.0
 - changelog was added
 - `Chat.chat_type` was renamed to just `Chat.type`
 - added new type `ChatMessage`. Currently only used for forwarded messages
 - forwards api was stabilized
