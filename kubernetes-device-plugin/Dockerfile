FROM golang:1.16 AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o coffee-device-plugin .

FROM scratch

COPY --from=build /src/bin/coffee-device-plugin /usr/local/bin/coffee-device-plugin

ENTRYPOINT ["/usr/local/bin/coffee-device-plugin"]