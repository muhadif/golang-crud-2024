FROM golang:1.21
LABEL authors="muhadif"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make run

EXPOSE 8080

# Command to run the executable
CMD ["./main"]
