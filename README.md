# Go API
This API is a project for me to learn to create a RESTful API using Golang. Where my main goal is to learn GO and also create a simple API that can be used in my app project.

---

## Usage

http://localhost:5006 or test on my hosted API https://go-api.noava.dev

Using these endpoints:
- `GET /severity?lat={lat}&lon={lon}&types={types}`
- `GET /pollen-info`


### Pollen Severity
*Shows the severity of the pollen*

`/severity?lat={lat}&lon={lon}&types={types}` Replace {lat} with the latitude, {lon} with the longitude, and {types} with the types of pollen you want to get. Seperated by comma. The types of pollen are:
- alder
- birch
- grass
- mugwort
- olive
- ragweed

This uses the API from [open-meteo.com](https://open-meteo.com/) to get the pollen data. Where i take the data and return it in a more readable format. Like how severe the pollen is and a message about how severe the pollen is.

This function returns:
```json
{
  "summary": {
    "{type}_pollen": "{severity}"
  },
  "interpretation": {
    "{type}": "{Message about how severe the pollen is.}"
  }
}
```

### Pollen Info
*Shows info about pollen*

`/pollen-info`

This shows info about pollen:
- What is pollen
- What is a pollen allergy 
- Symptoms
- Who is at risk
- Management tips
- Seasonal pollen

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
*Transparency*: Im logging the requests to see what users are requesting.