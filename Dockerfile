# Go application
FROM golang:alpine AS build
WORKDIR /go/src/griseo
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /griseo -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch AS runtime
WORKDIR /app
COPY --from=build /griseo /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
CMD ["/app/griseo","serve", "--http=0.0.0.0:8080"]
