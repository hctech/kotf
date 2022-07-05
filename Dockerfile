FROM golang:1.14-alpine as stage-build
LABEL stage=stage-build
WORKDIR /build/kotf
ARG GOPROXY
ARG GOARCH

ENV GO111MODULE=on \
    GOPROXY=$GOPROXY \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=$GOARCH

RUN  apk update \
  && apk add git \
  && apk add make \
  && apk add bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build_server_linux GOARCH=$GOARCH
RUN mkdir -p /build/kotf/plugins/
COPY /resource/plugins/  /build/kotf/plugins/

FROM kubeoperator/terraform:v0.12.28

RUN mkdir -p /root/.terraform.d/plugins/
COPY --from=stage-build /build/kotf/plugins/ /root/.terraform.d/plugins/
COPY --from=stage-build /build/kotf/dist/etc/ /etc/
COPY --from=stage-build /build/kotf/dist/usr/ /usr/
COPY --from=stage-build /build/kotf/dist/var/ /var/

VOLUME ["/var/kotf/data"]

EXPOSE 8080

CMD ["kotf-server"]
