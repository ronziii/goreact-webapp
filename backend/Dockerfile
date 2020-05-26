FROM golang:1.12.0
RUN apt-get update
RUN apt-get install -y git python jq curl
RUN curl -sL https://deb.nodesource.com/setup_9.x | bash -
RUN apt-get update && apt-get install -y nodejs
RUN go get -u github.com/pressly/goose/cmd/goose


RUN mkdir -p /apps/backend
WORKDIR /app/backend
COPY . .
RUN npm install --silent
RUN npm install -g gulp-cli && npm install gulp@3.9.1
COPY . .
RUN wget https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && chmod +x wait-for-it.sh
ENTRYPOINT ./wait-for-it.sh postgres:5432 -- goose postgres "postgres://postgres:password@postgres:5432/ictlife_infra_interview_may_2020" status && gulp server

