FROM node:12.16.3-alpine

ARG GOURL="https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz"

# Install Golang and set enironment
RUN wget $GOURL -P /opt
WORKDIR /opt
RUN tar -xf go1.14.3.linux-amd64.tar.gz
RUN cp -r go /$HOME/ 
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin 
ENV GOPROXY=direct


# set working directory frontend
RUN apk -U upgrade
RUN mkdir -p /app/frontend
WORKDIR /app/frontend
COPY frontend/package.json ./
COPY frontend/package-lock.json ./
RUN npm install --silent
COPY frontend/ .


# set working directory backend
RUN mkdir -p /apps/backend
WORKDIR /app/backend
COPY backend/package.json ./
COPY backend/package-lock.json ./
RUN npm install --silent
RUN npm install -g gulp-cli && npm install gulp@3.9.1
RUN npm audit fix --force
COPY backend/ .
RUN nohup gulp server &
