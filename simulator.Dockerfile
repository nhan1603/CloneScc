# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.20.2-alpine as builder_simulator

# create a working directory inside the image
WORKDIR /app

# copy directory files i.e all files ending with .go
COPY api .

# compile application
# /simulator: directory stores binaries file
RUN go build -o /simulator ./cmd/simulator/main.go

##
## STEP 2 - DEPLOY
##
FROM alpine:latest
WORKDIR /
COPY --from=builder_simulator /simulator /simulator

ENTRYPOINT ["./simulator"]