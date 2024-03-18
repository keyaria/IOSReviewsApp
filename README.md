# IOS Review Application
 
 Gathers the the Most Recent review for an application from the IOS store from the past 48 hours

## Tech Stack
- Vite
- Typescript
- Go
- PostgreSQL
- Docker

## Quick Start
```

// Build application in docker
docker compose build

// Run Docker Containers
docker compose up -d

// Stop Docker
docker componse stop

```
## Technical Decisions

- ViteJS
- Interfaces on FE

- Repository Pattern On BE


Packages added on BE
- CORS : github.com/rs/cors
    - Due to serving BE and FROM from different ports this is needed to add the header for CORS.

## Tradeoffs

- Writing to JSON file vs Database
    - Initially testing the API on JSON file, I decided to do a more real world Full Stack application as I have not much expereince working in Go. Becasue of this it gave me a chance to learn alot about Golang as I built.


## Improvements

- Add Swagger Doc for API
- Create a table that stores Users and that has link to all reviews from user