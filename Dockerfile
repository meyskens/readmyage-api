#Grab the latest alpine image
FROM arm64v8/golang:1.10

ARG QEMU_ARCH
COPY qemu-${QEMU_ARCH}-static /usr/bin/

ADD ./ /go/src/github.com/meyskens/readmyage-api

RUN cd /go/src/github.com/meyskens/readmyage-api && go install

CMD /go/bin/readmyage-api