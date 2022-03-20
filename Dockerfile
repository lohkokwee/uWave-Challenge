# Multistage build to reduce image size
FROM golang:latest as builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./ 
RUN go mod download
COPY . .
RUN go build -o main .

# Reduce size of docker image
FROM gcr.io/distroless/base-debian11
COPY --from=builder /usr/src/app/main ./
COPY --from=builder /usr/src/app/static ./static
EXPOSE 80
ENTRYPOINT [ "./main" ]