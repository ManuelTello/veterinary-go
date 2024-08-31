FROM golang:1.23-bookworm
WORKDIR /veterinary
ENV PORT=8080
ENV GO_CWD="/veterinary"
ENV GIN_MODE="release"
ENV API_KEY="6b8c0cd648c7400f92a82871b5a8a318"
EXPOSE 8080
ADD source ./
RUN go mod download 
RUN go mod verify
RUN go build -o bin cmd/veterinary.go 
CMD ["/veterinary/bin/veterinary" ]
