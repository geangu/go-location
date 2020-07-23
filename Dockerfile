FROM golang

WORKDIR /app

COPY . .

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-redis/redis
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get go.mongodb.org/mongo-driver/mongo/options

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]