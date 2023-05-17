## v1.9.0
 - `ClientRequest` allow multiple actions in one request for each action type
 - `StateResponse.max_actions_in_one_request` was added
 - `SendMessageListRequest` removed, use `SendMessageRequests`
 - `TOO_MANY_MESSAGES_IN_LIST_ERROR` removed in favor of `MAX_ACTIONS_IN_ONE_REQUEST_ERROR`
 - `INVALID_ACTION_ERROR` replaced with `NO_ACTIONS_IN_REQUEST_ERROR`

## v1.8.2
 - `SendMessageListRequest` restrictions were added: no more than `StateResponse.max_results_on_page` allowed

## v1.8.1
 - `MessageViewersRequest` parameters `limit` and `offset` were added

## v1.8.0
 - `Message.system_message` was added
 - all .proto files were moved to `protobuf` folder
 - experimental field `SendMessageResponse.errors_hint` was added
 - `MessagesListRequest.skip_system_messages` filter

## v1.7.0
 - `ChatMemberXXX` messages renamed to `GroupMemberXXX`
 - `GroupChatMemberDetailsRequest` removed, use `GroupMembersListResponse` with filter `user_id`
 -  introduced group roles: available in `GroupMembersListResponse.roles` and `GroupRolesResponse.roles`
 - `UpdateGroupMemberRequest` for changing member role
 - `StateResponse.gtproto_version` was added

## v1.6.0
 - `SendMessageListRequest` was added. Useful for forwarding and resending messages

## v1.5.0
 - `StateResponse.push_notifications_enabled` was added. It's `true` when push notifications are enabled for current server
 - `StateResponse.auth_methods`, `StateResponse.min_password_length` and `StateResponse.max_password_length` 
 - `Profile.has_password` 
 - `SetPasswordRequest`, `SetEmailRequest` and `SetEmailConfirmationRequest`

## v1.4.5
 - `SetMessageReactionResponse` now contains `recent` and `default_reaction` fields like `ReactionsResponse`

## v1.4.4
 - `UsersListRequest` extended with filters: `only_bots`, `only_online`, `only_my_contacts`

## v1.4.3
 - `UploadRequest.as_file` was renamed to `UploadRequest.disable_modifications`

## v1.4.2
 - `ReactionsResponse.default_reaction` was added

## v1.4.1
 - `Device.disable_pushes_when_online` field

## v1.4.0
 - bot commands were added
 - `ChatContext` was renamed to `ChatState`

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
