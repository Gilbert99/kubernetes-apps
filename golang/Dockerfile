FROM golang:1.18
LABEL auther="my first app"
WORKDIR gilbert-golang
COPY go.mod .
RUN go mod download
COPY . .
EXPOSE 4000
RUN go build -o gilbert-golang .
CMD ["./gilbert-golang"] 
