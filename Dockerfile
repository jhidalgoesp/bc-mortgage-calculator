FROM golang:alpine as build
RUN mkdir /service
COPY go.* /service
WORKDIR /service
RUN go mod download
COPY . /service
WORKDIR /service/cmd
RUN go build
CMD ["./cmd"]


FROM alpine:latest
WORKDIR /service
COPY --from=build /service/cmd/cmd /service/cmd
EXPOSE 3000
CMD ["./cmd"]
