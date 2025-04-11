FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

FROM scratch AS prod

COPY --from=builder /bin/api /bin/api
COPY --from=builder /app/views /bin/views

ENV VIEW_PATH="/bin/views"

ENTRYPOINT ["/bin/api"]

EXPOSE ${PORT}