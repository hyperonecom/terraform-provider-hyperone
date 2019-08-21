FROM golang:1.12
WORKDIR /go/github.com/hyperonecom/terraform-provider-hyperone
COPY ./ .
RUN make build
