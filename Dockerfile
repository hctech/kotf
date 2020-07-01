FROM golang:1.14-alpine as stage-build
LABEL stage=stage-build
WORKDIR /build/koft
ARG GOPROXY
ENV GOPROXY=$GOPROXY
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
  && apk update \
  && apk add git \
  && apk add make \
  && apk add bash
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build_server_linux

COPY ./resource/install_terraform.sh /build/kotf/install_terraform.sh
RUN bash /build/kotf/install_terraform.sh

COPY . /build/kotf/


FROM golang:1.14-alpine

COPY --from=stage-build /build/kotf/dist/etc/ /etc/
COPY --from=stage-build /build/kotf/dist/usr/ /usr/
COPY --from=stage-build /build/kotf/dist/var/ /var/

VOLUME ["/var/koft/data"]

EXPOSE 8080

CMD ["kotf-server"]