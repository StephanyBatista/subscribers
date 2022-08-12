FROM golang:alpine AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux
RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN ls
RUN go build -ldflags="-s -w" -o api main.go

FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/America/Sao_Paulo /usr/share/zoneinfo/America/Sao_Paulo
ENV TZ America/Sao_Paulo
WORKDIR /app
COPY --from=builder /app/api /app/api
COPY --from=builder /app/.env /app
EXPOSE 6004
CMD ["./api"]