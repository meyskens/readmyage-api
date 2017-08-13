#Grab the latest alpine image
FROM golang:1.9-alpine

# Install python and pip
ADD ./ /go/src/github.com/meyskens/readmyage-api

RUN cd /go/src/github.com/meyskens/readmyage-api && go install

CMD /go/bin/readmyage-api