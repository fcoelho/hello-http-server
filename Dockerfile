FROM golang:alpine as base

COPY server.go /server.go
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /server /server.go

FROM scratch

COPY --from=base /server /server

ENTRYPOINT ["/server"]
