FROM golang:1.23.3@sha256:017ec6b8081a283be9a45726f55e6e2830a7d66f94dbc09c557ca4ca1e6f7333

# Enviroment variable
WORKDIR /usr/src/some-api

RUN go install github.com/air-verse/air@latest

#Copying files to work directory
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

# Run and expose the server on port 3000
EXPOSE 3000

# CMD [ "nodemon", "build/app.js" ]
