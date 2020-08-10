FROM golang:1.14-alpine as stage-build
LABEL stage=stage-build
WORKDIR /build/kotf
ARG GOPROXY
ARG GOARCH

ENV GOPROXY=$GOPROXY
ENV GOARCH=$GOARCH
ENV GO111MODULE=on
ENV GOOS=linux
ENV CGO_ENABLED=0

RUN  apk update \
  && apk add git \
  && apk add make \
  && apk add bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build_server_linux GOARCH=$GOARCH

RUN wget https://releases.hashicorp.com/terraform/0.12.28/terraform_0.12.28_linux_$GOARCH.zip -O /tmp/terraform_0.12.28_linux_$GOARCH.zip \
    && cd /tmp \
    && unzip /tmp/terraform_0.12.28_linux_amd64.zip -d /build/kotf/

RUN mkdir -p /build/kotf/plugins/
COPY /resource/plugins/  /build/kotf/plugins/

FROM alpine:3.11

RUN mkdir -p /root/.terraform.d/plugins/
COPY --from=stage-build /build/kotf/terraform /usr/local/bin/
COPY --from=stage-build /build/kotf/plugins/ /root/.terraform.d/plugins/
COPY --from=stage-build /build/kotf/dist/etc/ /etc/
COPY --from=stage-build /build/kotf/dist/usr/ /usr/
COPY --from=stage-build /build/kotf/dist/var/ /var/

VOLUME ["/var/kotf/data"]

EXPOSE 8080

CMD ["kotf-server"]
