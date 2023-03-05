FROM golang:1.20 as builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go
COPY cmd/ cmd/
COPY cardversions/ cardversions/
COPY leagueclubs/ leagueclubs/
COPY nations/ nations/
COPY players/ players/

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o futbin main.go

FROM gcr.io/distroless/static:nonroot
USER 65532:65532
WORKDIR /
COPY --from=builder /workspace/futbin .

ENTRYPOINT ["/futbin"]
