FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /root/kpi-backend/

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/root/kpi-backend/binary"]