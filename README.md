# Sport Score

Sport Score is an API to simulate matches between two football (the real one) teams, compute their scores and send notification to a user (email, sms, whatever).

# Requirements

- Docker
- Docker Compose
- Go 1.21

## How to start

Make a copy of `env.example` file with `.env` name in the project root.

```
cp env.example .env
```

The content in the `env.example` is enough, but you can modify it as you wish.


```
docker-compose up --build
```

When the application is up, it will be available for requests on `http://localhost:3000`

## Endpoints

The content type for all endpoints is `application/json`.

### Retrieve all countries

```
GET /countries

Response:
- status 200
```

### Get a country by id

```
GET /countries/:id

Response:
- status 200
```

### Create a country

```
POST /countries

Attributes:
- name (required)

Response:
- status 201
```

### Update a country

```
PATCH /countries/:id

Attributes:
- name (required)

Response:
- status 202
```

### Delete a country

```
DELETE /countries/:id

Response:
- status 204
```
