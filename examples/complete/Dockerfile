FROM golang:1.24 AS builder

ARG TARGETOS
ARG TARGETARCH
ARG COMMIT
ARG DATE

ARG VERSION=dev
ENV VERSION=$VERSION

ENV COMMIT=$COMMIT

ENV DATE=$DATE

WORKDIR /app/

# Copy only the files required for go mod download
COPY ./go.* .

# Download dependencies
RUN go mod download

# Copy the rest of the files
COPY . .

# Build router
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags "-extldflags=-static \
    -X 'github.com/wundergraph/cosmo/router/core.Version=${VERSION}' \
    -X 'github.com/wundergraph/cosmo/router/core.Commit=${COMMIT}' \
    -X 'github.com/wundergraph/cosmo/router/core.Date=${DATE}'" \
    -a -o router main.go

FROM gcr.io/distroless/static:latest

COPY --from=builder /app/config.yaml /config.yaml
COPY --from=builder /app/router /router

ENTRYPOINT ["/router"]

EXPOSE 3002
