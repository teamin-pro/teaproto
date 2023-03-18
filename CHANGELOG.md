## v1.3.0
 - `CreateGroupRequest` / `UpdateGroupSettingsRequest` now returns `Chat` object
 - `Group.chat_id` removed because group is always part of chat
 - removed `GroupDetailsRequest`, use `ChatDetailsRequest`
 - removed `GroupListRequest`, use `ChatListRequest`
 - `CreateGroupXXX` renamed to `CreateGroupChatXXX`
 - `UpdateGroupSettingsXXX` renamed to `UpdateGroupChatSettingsXXX`
 - `Chat` / `ShortChat` pair replaced with `Chat` that have optional `ChatContext` field. `LastMessage`, `Badge` and other fields moved to the `ChatContext`

## v1.2.2
 - `Chat` and `ShortChat` have new field `topic` with `User` or `Group` inside 

## v1.2.1
 - `Chat.with_user_id` replaced with `Chat.with_user`

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
