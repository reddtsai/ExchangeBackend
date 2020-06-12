FROM golang:1.14 AS build
WORKDIR /go/src/app
COPY . .
RUN mkdir bin/trading & go build -o bin/trading ./cmd/trading

FROM golang:1.14 AS publish
WORKDIR /trading
COPY --from=build /go/src/app/bin/trading .
EXPOSE 5001
ENTRYPOINT [ "./trading" ]