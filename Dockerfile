#Grab the latest alpine image
FROM golang:1.9-alpine

# Install python and pip
ADD ./ /opt/readmyage

RUN cd /opt/readmyage && go build ./

WORKDIR /opt/readmyage

CMD ./readmyage