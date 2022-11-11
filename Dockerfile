# Builder
FROM golang:1.19-alpine as development
WORKDIR /go/src/app
RUN apk update && apk add --no-cache gcc musl-dev bash tzdata
COPY go.* ./
RUN go mod download -x
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o api cmd/main.go
CMD ["go", "run", "cmd/main.go"]

# Distribution
FROM alpine:latest as production
WORKDIR /app
RUN apk update && apk add --no-cache bash
COPY --from=development /go/src/app/api .
COPY --from=development /go/src/app/data ./data
COPY --from=development /usr/share/zoneinfo /usr/share/zoneinfo
CMD ["./api"]