FROM golang:1.7.1

RUN mkdir /app
ADD . /app/ 
WORKDIR /app/bin

RUN go get github.com/alfredxing/calc/... && \
    go get golang.org/x/crypto/bcrypt

RUN cd ../src && go build -o ../bin/goq && \
    cd worker && go build -o ../../bin/worker && \
    cd ../client && go build -o ../../bin/client

CMD ["/app/bin/goq"]
EXPOSE 9001