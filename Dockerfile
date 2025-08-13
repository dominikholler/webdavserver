FROM docker.io/library/golang:1.25-alpine as builder
#FROM registry.redhat.io/rhel9/go-toolset:1.20 as builder
WORKDIR /opt/app-root/src/

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /opt/app-root/webdavserver
RUN mkdir /tmp/webdav


FROM scratch
ARG USER=1001

COPY --from=builder /opt/app-root/webdavserver /webdavserver
COPY --from=builder --chown=$USER:$USER /tmp/webdav /srv/webdav

USER $USER:$USER

CMD ["-webdavDir", "/srv/webdav/"]
ENTRYPOINT ["/webdavserver"]
