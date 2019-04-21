# go-user-auth

This is a starter repo for token-based user authentication in a json web api - implemented in golang.

## Features

* Dockerized postgresql - migrations handled by [rambler](https://github.com/elwinar/rambler), DB interaction via [sqlx](https://github.com/jmoiron/sqlx).
* Hashed (and salted) password persistence via [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt).
* JWT claims via [jwt-go](https://github.com/dgrijalva/jwt-go).
* Routing and middleware with [gorilla/mux](https://github.com/gorilla/mux).

## Usage

* Clone repo to your gopath.
* Copy `.env.sample` to `.env` and set the variables.
* Run the database with `docker-compose up -d`
* Run migrations with `docker-compose run migrate`
* Install dependencies `go get`
* Build and run the go binary `./build-and-run.sh`
* API will be up at `localhost:8080`

## Endpoints

### POST /signup

Sign up like this:

```json
{
  "email": "me@example.com",
  "password": "password"
}
```

### POST /signin

Get a token. The request body is the same as `/signup`. Here's an example response:

```json
{
  "token":"eyJhbGciOiJIUzI1...",
  "expiration":"2020-04-20 18:14:21.282505 -0400 EDT m=+31539368.330614446"
}
```

### GET /users/me

This is an example endpoint which requires authentication. Put a valid token in the header: `Authorization: Bearer eyJhbGciOiJIUzI1...`. The response be a 401 if unauthenticated; if authenticated it will look like this:

```json
{
  "id": 1,
  "email": "me@example.com",
  "createdAt": "2019-04-21T01:35:34.820758Z"
}
```

## Todo

* Password length validation
* Email validation
* Token revocation
* Refresh token
* Password update/reset
