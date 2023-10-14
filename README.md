# Basic Auth

This API provides authorization and basic user management functionality using the Chi-Router for Go.

## Usage

When developing locally make sure that the mongo database is running and use (set flags if needed like in the [Config](#configuration) shown):

> go run ./cmd/api -cors="\*"

## User

- ID
- Username
- Email
- Password (hashed)
- ~~GoogleAccessToken~~
- EmailActivated
- EmailActivateToken
- EmailActivateExpires
- Avatar (link)
- Roles
- Inactive
- ResetPasswordToken
- ResetPasswordExpires
- CreatedAt
- UpdatedAt

## Endpoints

| Request | Path                           | Description             | Required Auth       | Done |
| ------- | ------------------------------ | ----------------------- | ------------------- | ---- |
| GET     | `/auth/api-status`             | Status of the API       | -                   | done |
| POST    | `/auth/sign-up`                | User register           | -                   | done |
| POST    | `/auth/sign-in`                | User login              | -                   | done |
| GET     | `/auth/sign-out`               | User logout             | -                   | done |
| POST    | `/auth/activate-email`         | Activate user email     | -                   | done |
| POST    | `/auth/reset-password`         | Resets password         | -                   | done |
| POST    | `/auth/reset-password/revoke`  | Revoke password request | auth/admin required | done |
| POST    | `/auth/reset-password/request` | Resets password request | -                   | done |
| GET     | `/users`                       | Gets user list          | auth/admin required | done |
| GET     | `/users/me`                    | Gets current user       | auth required       | done |
| DELETE  | `/users/me`                    | Deletes user            | auth required       | done |
| PATCH   | `/users/me/username`           | Updates user username   | auth required       | done |
| PATCH   | `/users/me/email`              | Updates user email      | auth required       | done |
| PATCH   | `/users/me/password`           | Updates user password   | auth required       | done |
| PATCH   | `/users/me/avatar`             | Updates user avatar     | auth required       | done |
| GET     | `/users/{id}`                  | Gets a user by ID       | auth required       | done |
| PUT     | `/users/{id}`                  | Updates user            | auth/admin required | done |
| DELETE  | `/users/{id}`                  | Deletes user            | auth/admin required | done |
| GET     | `/users/{id}/reactivate`       | Reactivate user         | auth/auth required  | done |
| GET     | `/users/avatar/{filename}`     | Gets the user's avatar  | auth required       | done |

The Baserouting for `/users` can get changed using the `baseRouting` flag in the settings. If `auth` is chosen for the baseRouting then `/users` turns into `/users/list`.

## Configuration

| Flag           | Default                   | Description                               |
| -------------- | ------------------------- | ----------------------------------------- |
| version        | 1.0.0                     | the app version                           |
| env            | dev                       | the app environment                       |
| servername     | http://localhost          | the api server name                       |
| port           | 8080                      | the used port                             |
| dsn            | mongodb://localhost:27017 | the database connection string            |
| dbName         | basic-auth                | te name of the used database              |
| jwt            | wonderfulsecretphrase     | the jwt token secret                      |
| cors           | http://\* https://\*      | the by cors allowed origin servers        |
| cookieName     | basic-auth                | the name of the cookie                    |
| cookieSameSite | lax                       | the cookie same site policy               |
| addRoles       | guest                     | the roles (+admin)                        |
| defaultRole    | guest                     | the default role when creating a new user |
| baseRouting    | users                     | the base routing for the users endpoints  |

## Password reset

1. Request the password reset (`/auth/reset-password/request`)
2. Get email (and token to reset, resetduration in db)
3. Use token to reset password (`/auth/reset-password?token=TOKEN`)
4. Can get revoked by admins (`/auth/reset-password/revoke`)

## Contributing

I am always happy for tips and suggestions to improve it.
There might be routings and settings that I have missed.

### TODO

- SignUp (send verification email)
- Email stuff ...
  - Better templates
- Logging to file and/or console (based on env)
- Swagger documentation
- Tests
- Postman?
- Dockerfile fixes (maybe docker-compose)

## License

This project is licensed under the [MIT](/LICENSE) license.
