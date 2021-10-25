FROM golang:1.17.2-alpine

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

RUN go install github.com/spf13/cobra/cobra@latest
RUN apk add --update make

CMD ["tail", "-f", "/dev/null"]