# Go API
This API is a project for me to learn to create a RESTful API using Golang. Where my main goal is to learn GO and also create a simple API that can be used in my app project.

---

## Usage

http://localhost:5006 or test on my hosted API https://go-api.noava.dev

Using these endpoints:
- GET /hello
- GET /bye?name={name}

### Hello
/hello

This function returns:
```
"Hello, World!"
```

### Bye
/bye?name={name} takes a name as a query variable. Where you can replace {name} with your name. Where no name is provided, the default name is "World".

This function returns:
```
"Bye, {name}!"
```

---

## Rate Limiter
This API has a rate limiter that limits the number of requests to 10 requests per second. Where if the rate limit is exceeded, it will return a 429 status code with the message "Rate limit exceeded".

---

## Prerequisites

- Docker or run in terminal
- Go

### Build and Run

#### Run in Docker
```terminal
docker build -t go-api .
```

```terminal
docker run -p 5006:5006 go-api
```
#### Run in terminal
```terminal
go run main.go
```

You can use postman or the browser to test the API.
