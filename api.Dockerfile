# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.20.2-alpine as builder

# create a working directory inside the image
WORKDIR /app

# copy directory files i.e all files ending with .go
COPY api .

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

# compile application
# /sccapi: directory stores binaries file
RUN go build -o /sccapi ./cmd/serverd/main.go ./cmd/serverd/router.go

##
## STEP 2 - DEPLOY
##

FROM debian:buster

# Install ca-certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Install necessary dependencies, including ffmpeg
RUN apt-get update && \
    apt-get install -y ffmpeg && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /

COPY --from=builder /sccapi /sccapi

COPY api/certs/ /certs/

ENTRYPOINT ["./sccapi"]
