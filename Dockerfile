FROM golang:1.17

WORKDIR /mnt/HDD/Applications/JetBrains/Projects/Goland/password-manager-backend
COPY . .

# packages
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD go run main.go
