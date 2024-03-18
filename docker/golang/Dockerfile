FROM golang:1.22.1-bookworm


COPY ./server /app/go
WORKDIR /app/go

# Download and install the dependencies:
RUN go get -d -v /app/go

RUN go build -o main .

EXPOSE 8080

CMD [ "/app/go/main" ]

