# Go MongoDB User API
![Superhero Gopher - Project Title Image](https://raw.githubusercontent.com/egonelbre/gophers/63b1f5a9f334f9e23735c6e09ac003479ffe5df5/vector/superhero/standing.svg)

A simple CRUD API using Go, Docker, and MongoDB. 
Initial project was highly-coupled and had a handful. 
The project has been refactored and implements the 
repository pattern to reduce that decouple the database and
service logic. 
Commit history includes multiple changes to the code. 


## Features

- RESTful API: Supports creating, retrieving, and deleting users via HTTP endpoints
- Abstraction of storage logic for decoupled database: use of repository pattern for decoupling
- Use `net/http` for routing and request/response handling

## Getting Started

### Prerequisites
- Docker
- Docker Compose
- Go (1.18+ recommended)

## Installation

1. Clone this repository:
   ```sh
   git clone https://github.com/travboz/go-mongodb-user-api.git
   cd go-mongodb-user-api
   ```
2. Run docker container:
    ```sh
    make up
    ```
3. Seed MongoDB instance:
   ```sh
   make seed-db
   ```
4. Run server:
    ```sh
    make run
    ```
5. Navigate to `http://localhost<SERVER_PORT>` and call an endpoint

### `.env` file
This server uses a `.env` file for basic configuration.
Here is an example of the `.env`:
   ```sh
   DB_CONTAINER_NAME=MONGO-USER-CRUD
   SERVER_PORT=":8080"
   MONGO_DB_NAME=mongo_user_crud
   MONGO_DB_USERNAME=<your-username>
   MONGO_DB_PASSWORD=<your-password>
   MONGODB_URI=mongodb://<your-username>:<your-password>@localhost:27017/mongo_user_crud?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false
   COMPASS_USER_MONGODB_URI=mongodb://<your-username>:<your-password>@localhost:27017/mongo_user_crud?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false
   ```
   
## API endpoints

| Method   | Endpoint        | Description          |
|----------|----------------|----------------------|
| `GET`    | `/`            | Welcome message/health check     |
| `POST`   | `/users`       | Create a new user   |
| `GET`    | `/users`       | Get all users       |
| `GET`    | `/users/{id}`  | Get user by ID      |
| `PUT`    | `/users/{id}`  | Update a user       |
| `DELETE` | `/users/{id}`  | Delete a user       |

## Example usage

### JSON payload structures

#### Create user payload

```json
{
  "name": "bob jones",
  "email": "bob@jones.com",
  "favourite_number": 25,
  "active": false
}
```

#### Update user payload

```json
{
  "name": "new jones",
  "email": "bob@jones.com",
  "favourite_number": 1000,
  "active": true
}
```

### Endpoint example usage
#### Create a user
```sh
curl -X POST "http://localhost:8080/users" \
     -H "Content-Type: application/json" \
     -d '{
       "name": "bob jones",
       "email": "bob@jones.com",
       "favourite_number": 25,
       "active": false
     }'
```

#### Update a user
```sh
curl -X POST "http://localhost:8080/users/67a0a3eef39fc03fe52450b5" \
     -H "Content-Type: application/json" \
     -d '{
       "name": "new jones",
       "email": "bob@jones.com",
       "favourite_number": 1000,
       "active": true
     }'
```

#### Get a user by id
```sh
curl -X GET "http://localhost:8080/users/67a0a3eef39fc03fe52450b5"
```

#### Fetch all users
```sh
curl http://localhost:8080/users
```

#### Delete a user
```sh
curl -X DELETE "http://localhost:8080/users/67a0a3eef39fc03fe52450b5"
```

## Contributing
Feel free to fork and submit PRs!

## License:
`MIT`


This should work for GitHub! Let me know if you need any tweaks. 


## Image
Image by [Egon Elbre](https://github.com/egonelbre), used under CC0-1.0 license.
