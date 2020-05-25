ICTLife Infra May 2020 Interview Challenge
=========================================

Backend for merchant dashboard application


Getting Started
===============
- Clone repo
- Install Node.js v10 [Download](https://nodejs.org/dist/latest-v10.x/)
- Install Golang 1.12 [Download](https://golang.org/dl/#go1.12.13)
- Install Docker [Download](https://docs.docker.com/)
- Install goose: `go get -u github.com/pressly/goose/cmd/goose`
- Download npm packages: `npm install`
- Create local env file: `cp .env.development.sample .env.development`
- Update env file with correct configuration
- Create docker compose file `cp docker-compose.yml.sample docker-compose.yml`
  - Configure PGDATA path (ex: `/Users/{username}/apps/ictlife_infra_interview_may_2020/pgdata`)
- Run database in Docker: `docker-compose up`
- Setup database: `goose postgres -dir app/db/migrations up`
- Run server: `gulp server`
