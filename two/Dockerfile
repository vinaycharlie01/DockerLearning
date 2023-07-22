FROM golang:alpine

WORKDIR /app

COPY go.mod .

ENV GO111MODULE=on
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o main /app/add/main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o main2 /app/adde/main.go


# COPY --from=builder main main


CMD [ "/app/main" ,"/app/main2" ]