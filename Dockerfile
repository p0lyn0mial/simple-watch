FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.16-openshift-4.8 AS builder
WORKDIR /go/src/github.com/p0lyn0mial/simple-watch
COPY . .
RUN go build -o ./app .

FROM debian
COPY --from=builder /go/src/github.com/p0lyn0mial/simple-watch/app /usr/bin/
ENTRYPOINT /usr/bin/app
