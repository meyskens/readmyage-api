FROM golang as build

ADD ./ /go/src/github.com/meyskens/readmyage-api

WORKDIR /go/src/github.com/meyskens/readmyage-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o readmyage-api ./


FROM multiarch/alpine:arm64-edge

COPY --from=build /go/src/github.com/meyskens/readmyage-api/readmyage-api /usr/local/bin/readmyage-api

CMD [ "readmyage-api" ]