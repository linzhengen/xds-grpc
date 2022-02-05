FROM golang:1.17 as build

ARG GRPC_HEALTH_PROBE_VERSION=0.4.2
RUN wget -q -O /bin/grpc_health_probe "https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64" && \
    chmod +x /bin/grpc_health_probe

WORKDIR /src
COPY go.mod /src/
COPY go.sum /src/
RUN go mod download
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/helloworld-server -ldflags="-s -w" -trimpath .


FROM gcr.io/distroless/base:debug-nonroot

ENV GOTRACEBACK=single
ENV PORT=50051
EXPOSE 50051

COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe
COPY --from=build /bin/helloworld-server /bin/helloworld-server

ENTRYPOINT ["/bin/helloworld-server"]