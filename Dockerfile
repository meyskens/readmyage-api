#Grab the latest alpine image
FROM arm64v8/golang:1.10-alpine

ARG QEMU_ARCH
COPY qemu-${QEMU_ARCH}-static /usr/bin/

RUN apk add --no-cache gcc

ADD ./ /go/src/github.com/meyskens/readmyage-api

RUN cd /go/src/github.com/meyskens/readmyage-api && go install

CMD /go/bin/readmyage-api