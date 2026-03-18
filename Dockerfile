# Build Stage
FROM golang:1.22-alpine AS build

WORKDIR /app

# Dependencies first for docker cache
COPY go.mod ./

RUN go mod download

# copy rest of code
COPY . .

# build go application
RUN go build -o trade-vista

# Final Stage
FROM alpine:latest

WORKDIR /app

# copy binary
COPY --from=build /app/trade-vista .

# copy templates
COPY --from=build /app/templates ./templates

# Port and CMD
EXPOSE 8080

CMD ["./trade-vista"]

# metadate
LABEL maintainer="timit06@googlemail.com"
LABEL version="1.0"
LABEL description="Trade Vista - A modern trading application"
LABEL repository="https://github.com/kranjuli/trade-vista"
