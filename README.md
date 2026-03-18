# Trade-Vista

A small website to display crypto transactions on Bitvavo

## Structure of project

````
trade-vista/
|- templates/
   |- trading.html
|- transactions/
   |- 2025.csv
   |- 2024.csv
|- .gitignore
|- go.mod
|- main.go
|- models.go
|- README.md
|- utils.go
|- Dockerfile
|- .dockerignore
````

## Run go application

````
cd trade-vista
go run .
````

## Run Go Application with Docker

### Building Docker Images Using a Multi-Stage Dockerfile

This project uses a multi-stage Docker build to keep the final image lightweight and efficient.
The process consists of two stages:

#### Stage 1: Build go application

````
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
````

#### Stage 2: Create the Final Runtime Image

This stage creates a minimal image containing only the compiled binary and required assets.

````
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
````

### Build the Docker Image and Run a container

````
cd trade-vista

# build docker image trafe-vista

sudo docker build -t trade-vista .

# create docker container
sudo docker run -d --name trade-vista -v /mnt/ssd/data/bvv_transactions:/app/transactions:ro -p 8080:8080 trade-vista:latest

# login to container
sudo docker exec -it trade-vista /bin/sh

# delete container
sudo docker rm trade-vista

# delete image
sudo docker rmi trade-vista
````

NOTE: 

- The multi-stage build reduces the final image size by excluding build tools.
- The `/mnt/ssd/data/bvv_transactions` directory is mounted as a read-only volume inside the container.
- The application will be available at: `http://localhost:8080`

## Deployment with crypto-infra-deploy

see project crypto-infra-deploy