##
## Build
##
FROM golang:1.17 AS build
WORKDIR /app
COPY . ./
RUN go mod download
COPY *.go ./
RUN go build -o /refresh-token

##
## Deploy
##
FROM nginx
WORKDIR /app
COPY --from=build /refresh-token /refresh-token
EXPOSE 8000
ENTRYPOINT ["/refresh-token"]