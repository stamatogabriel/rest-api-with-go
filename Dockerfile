FROM golang:1.24-alpine3.22

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]