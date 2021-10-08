FROM golang:1.14.6-alpine3.12 as build_base
WORKDIR /app

ENV GO111MODULE=on
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM build_base AS binary_builder
# Here we copy the rest of the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-fresh .

#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine AS database-migrator
# We add the certificates to be able to verify remote instances
RUN apk add ca-certificates
COPY --from=binary_builder /app /app
EXPOSE 3000 3000
ENTRYPOINT ["/app/hello-fresh"]