# IOS Review Application
 
 Gathers the the Most Recent review for an application from the IOS store from the past 64 hours

## Tech Stack
- ViteJS
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
    -  Compared to ReactJS it has a fast development server, build times, and framerwork agnositic allowing greater flexibility if applications from other languages are added so I decided to ise it instead of create-react-app.

- Interfaces on FE
    - Allows for code to become module and have Services/APP that are based on functionality. Users can add more services/interfaces based on implementations(ex. Adding FE for a Android reviews Application).  

- Repository Pattern On BE
    - In order to seperate Application Logic from Database Logic and show a clear flow of how data is being moved
    - Can be further expanaded for a Domain Driven Design and Scalability


Packages added on BE
- CORS : github.com/rs/cors
    - Due to serving BE and FE from different ports this is needed to add the header for CORS.

## Tradeoffs

- Writing to JSON file vs Database
    - Initially testing the API on JSON file, I decided to do a more real world Full Stack application as I have not much expereince working in Go. Becasue of this it gave me a chance to learn alot about Golang as I built.


## Improvements

- Add Swagger Doc for API
- Add support for DB Migrations
- Create a table that stores Users and that has link to all reviews from user
- Improve Testing on BE with integration tests
- Cypress on FE to test entire flow