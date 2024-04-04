FROM golang:1.22.1

WORKDIR /app/KeDuBack/

COPY . .

RUN go mod tidy

RUN go build -o KeDuBack

EXPOSE 8080

RUN ["./KeDuBack"]