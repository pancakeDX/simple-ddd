FROM docker.io/golang:1.21-alpine AS build

WORKDIR /app

ADD . .

RUN go build -o simple-ddd -ldflags '-w -s' cmd/main.go

FROM scratch

COPY --from=build /app/simple-ddd /bin/simple-ddd

ENTRYPOINT ["/bin/simple-ddd"]
