# stage 1: build
FROM golang:latest AS builder

LABEL maintainer="Md. Ishtiaq Islam <islam.ishtiaq99@gmail.com>"

ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o auth-server .

# stage 2: minimal runtime container
FROM gcr.io/distroless/static-debian11

COPY --from=builder /app/auth-server /auth-server

EXPOSE 3000

CMD ["/auth-server"]