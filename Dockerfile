FROM golang:1.18
EXPOSE 3000
WORKDIR /app
COPY . .
RUN go get -d github.com/gorilla/mux

CMD ["go","run","main.go"]