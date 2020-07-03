FROM golang:1.14-alpine as stage-build
LABEL stage=stage-build
WORKDIR /build/kotf
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
RUN bash install_terraform.sh

RUN mkdir -p /build/kotf/plugins/
COPY /resource/plugins/  /build/kotf/plugins/

FROM alpine:3.11

RUN mkdir -p /root/.terraform.d/plugins/
COPY --from=stage-build /build/kotf/plugins/ /root/.terraform.d/plugins/
COPY --from=stage-build /build/kotf/dist/etc/ /etc/
COPY --from=stage-build /build/kotf/dist/usr/ /usr/
COPY --from=stage-build /build/kotf/terraform /usr/local/bin/
COPY --from=stage-build /build/kotf/dist/var/ /var/

VOLUME ["/var/kotf/data"]

EXPOSE 8080

CMD ["kotf-server"]
