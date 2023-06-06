# Authentication

## By password
`StateResponse.auth_methods` must contain `EMAIL_PASSWORD_AUTH_METHOD`

- send `EmailPasswordAuthResponse` with `email` and `password` fields

## By one-time code
`StateResponse.auth_methods` must contain `EMAIL_CODE_AUTH_METHOD`

- send `EmailCodeAuthRequest` with `email` field
- server will send email with code to `email` address
- send `EmailCodeAuthConfirmRequest` with `email` and `code` fields
