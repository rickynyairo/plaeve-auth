# plaeve-auth/Dockerfile
FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/plaeve-auth

COPY . .

# RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o plaeve-auth -a -installsuffix cgo main.go handler.go repository.go database.go token_service.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/plaeve-auth .

CMD ["./plaeve-auth"]