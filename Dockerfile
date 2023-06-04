FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/package/myapp/

RUN addgroup --system --gid 1001 golang
RUN adduser --system --uid 1001 app

COPY . .

RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /server

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /server /server

USER app:golang

ENTRYPOINT ["/server"]
