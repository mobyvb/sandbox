FROM golang:1.16
    WORKDIR /app
    COPY . .
    RUN go mod init simple-oauth-app
    RUN go mod tidy
    RUN go build -o simple-oauth-app
    EXPOSE 8080
    CMD [ "./simple-oauth-app" ]
    

  -