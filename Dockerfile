FROM golang as build

ADD ./ /go/src/github.com/meyskens/readmyage-api

WORKDIR /go/src/github.com/meyskens/readmyage-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o readmyage-api ./


FROM multiarch/alpine:amd64-edge

COPY --from=build /go/src/github.com/meyskens/readmyage-api/readmyage-api /usr/local/bin/readmyage-api

CMD [ "readmyage-api" ]
