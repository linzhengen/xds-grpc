FROM golang:1.17 as build

WORKDIR /src
COPY go.mod /src/
COPY go.sum /src/
RUN go mod download
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/helloworld-client -ldflags="-s -w" -trimpath .


FROM gcr.io/distroless/base:debug-nonroot

ENV GOTRACEBACK=single
ENV PORT=18080
EXPOSE 18080

COPY --from=build /bin/helloworld-client /bin/helloworld-client

ENTRYPOINT ["/bin/helloworld-client"]