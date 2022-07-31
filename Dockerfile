FROM golang:1.18
EXPOSE 8080
WORKDIR /app
COPY . .
RUN go get -d github.com/gorilla/mux

CMD ["go","run","main.go"]