FROM golang:1.18-alpine
WORKDIR /app
COPY ./ ./
RUN go mod download
WORKDIR /app/cmd
RUN go build -buildvcs=false -o urlshortener 
EXPOSE 8081
ENTRYPOINT ["./urlshortener"]