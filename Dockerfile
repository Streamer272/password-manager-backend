FROM golang:1.17

WORKDIR /mnt/HDD/Applications/JetBrains/Projects/Goland/password-manager-backend
COPY . .

# packages
RUN go get -d -v ./...
RUN go install -v ./...

CMD go run main.go
