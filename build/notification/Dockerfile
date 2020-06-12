FROM golang:1.14 AS build
WORKDIR /go/src/app
COPY . .
RUN mkdir bin/notification & go build -o bin/notification ./cmd/notification

FROM golang:1.14 AS publish
WORKDIR /notification
COPY --from=build /go/src/app/bin/notification .
EXPOSE 5000
ENTRYPOINT [ "./notification" ]