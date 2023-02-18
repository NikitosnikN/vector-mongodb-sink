FROM golang:1.20.1-alpine as builder


WORKDIR /home/app

COPY ./src /home/app

RUN go mod download
RUN go build -o /vector-mongodb-sink ./cmd/vector-mongodb-sink


FROM scratch

COPY --from=builder /vector-mongodb-sink /vector-mongodb-sink

CMD [ "/vector-mongodb-sink" ]