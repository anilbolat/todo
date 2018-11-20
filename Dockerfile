FROM golang:1.11 as builder

RUN groupadd -g 999 appuser && \
    useradd -r -u 999 -g appuser appuser

WORKDIR /todo
COPY . ./
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN make build

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /todo/builds/todo /todo
USER appuser
ENTRYPOINT [ "/todo" ]
