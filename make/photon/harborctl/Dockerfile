FROM golang:1.12 as builder

WORKDIR $GOPATH/src/github.com/heww/harborctl
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/harborctl .


FROM photon:2.0

RUN mkdir /harbor/ \
    && tdnf install sudo -y >> /dev/null\
    && tdnf clean all \
    && groupadd -r -g 10000 harbor && useradd --no-log-init -r -g 10000 -u 10000 gru

COPY --from=builder /go/bin/harborctl /harbor/harborctl

RUN chmod u+x /harbor/harborctl
WORKDIR /harbor
ENTRYPOINT ["/harbor/harborctl"]