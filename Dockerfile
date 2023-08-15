FROM jac18281828/godev:latest

ARG PROJECT=gocrucible
WORKDIR /workspaces/${PROJECT}
ENV GOMAXPROCS=10
COPY . .
RUN chown -R jac:jac .
USER jac
ENV GOPATH=/workspaces/${PROJECT}

RUN go install -v github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /workspaces/${PROJECT}/matrix

RUN go build
CMD go test

