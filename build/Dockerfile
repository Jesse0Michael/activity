FROM golang:1.22-alpine AS build

# Install git
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git 

# Copy project files
WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN mkdir /dir

# Fetch dependencies
RUN go mod download
COPY . .

# Build project
ENV CGO_ENABLED=0
RUN go build -o app ./cmd/activity-api

FROM gcr.io/distroless/static

COPY --from=build /go/src/app ./

ENTRYPOINT ["./app"]
