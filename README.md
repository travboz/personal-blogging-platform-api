# Personal Blogging Platform API
![Superhero Gopher - Project Title Image](https://raw.githubusercontent.com/egonelbre/gophers/63b1f5a9f334f9e23735c6e09ac003479ffe5df5/vector/superhero/standing.svg)

## Description

This is a RESTful API designed to power a personal blog. It provides endpoints to manage blog articles, allowing for full CRUD functionality — Create, Read, Update, and Delete. Articles can be listed, filtered by tags or publish date, retrieved individually by ID, and managed through standard HTTP methods.

The brief follows:
[Personal Blogging Platform API Project](https://roadmap.sh/backend/project-ideas#1-personal-blogging-platform-api:~:text=1.%20Personal%20Blogging%20Platform%20API)

## Features

- List all articles with optional filters (e.g., tags, publish date)
- Retrieve a single article by ID
- Create a new article
- Update an existing article by ID
- Delete an article by ID
- Implements basic CRUD operations over HTTP


## Getting Started

### Prerequisites
- Docker
- Docker Compose
- Go (1.18+ recommended)

## Installation

1. Clone this repository:
   ```sh
   git clone https://github.com/travboz/personal-blogging-platform-api.git
   cd personal-blogging-platform-api
   ```
2. Set up Go modules:
   ```sh
   go mod tidy
   ```   
3. Run docker container containing MongoDB instance:
    ```sh
    make compose/up
    ```
4. Seed MongoDB instance:
   ```sh
   make seed
   ```
5. Run server:
    ```sh
    make run
    ```
6. Navigate to `http://localhost<SERVER_PORT>` and call an endpoint

I will use example port `":7666"`.

### `.env` file
This server uses a `.env` environment file for configuration.
For an example, see `.env.example`.

## API Endpoints

| Method    | Endpoint           | Description                    |
|-----------|--------------------|--------------------------------|
| `GET`     | `/health`          | Health check                   |
| `POST`    | `/articles`        | Create a new article           |
| `GET`     | `/articles`        | Get all articles               |
| `GET`     | `/articles/:id`    | Get article by ID              |
| `PATCH`   | `/articles/:id`    | Update an article              |
| `DELETE`  | `/articles/:id`    | Delete an article              |

## Example usage

### JSON payload structures

#### Create article payload

```json
{
  "content": "this is the content for a new article",
  "tags": ["these", "are", "the", "tags"]
}
```

#### Update user payload

```json
{
  "content": "this is the NEW CONTENT for an existinf article",
  "tags": ["where", "did", "the", "old", "tags", "go?"]
}
```

### Endpoint example usage
#### Create a user
```sh
curl -X POST "http://localhost:8080/articles" \
     -H "Content-Type: application/json" \
     -d '{
        "content": "this is the content for a new article",
        "tags": ["these", "are", "the", "tags"]
     }'
```

#### Update a user
```sh
curl -X POST "http://localhost:8080/users/67a0a3eef39fc03fe52450b5" \
     -H "Content-Type: application/json" \
     -d '{
        "content": "this is the NEW CONTENT for an existinf article",
        "tags": ["where", "did", "the", "old", "tags", "go?"]
      }'
```

#### Get a user by id
```sh
curl -X GET "http://localhost:7666/articles/67a0a3eef39fc03fe52450b5"
```

#### Fetch all users
```sh
curl http://localhost:7666/articles
```

#### Delete a user
```sh
curl -X DELETE "http://localhost:7666/articles/67a0a3eef39fc03fe52450b5"
```

## Contributing
Feel free to fork and submit PRs!

## License: `MIT`


If there are any concerns regarding the licence, please contact me at `travis.bozic@hotmail.com`.


## Image
Image by [Egon Elbre](https://github.com/egonelbre), used under CC0-1.0 license.
