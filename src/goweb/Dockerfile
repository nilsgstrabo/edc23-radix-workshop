FROM golang:1.20-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /goweb

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /goweb /goweb

USER 101

ENTRYPOINT [ "/goweb" ]