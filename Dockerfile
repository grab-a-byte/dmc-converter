FROM golang:alpine AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -o /build/dmc_convert


FROM alpine
COPY --from=build /build/dmc_convert /
EXPOSE 8080
ENTRYPOINT  ["/dmc_convert"]