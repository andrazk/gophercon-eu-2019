FROM golang:1.12.5-stretch as builder

COPY . /app

WORKDIR /app

ENV PROJECT github.com/andrazk/tenerife

RUN CGO_ENABLED=0 go build \
		-ldflags " \
	        -X ${PROJECT}/internal/diagnostics.revision=${COMMIT} \
            -X ${PROJECT}/internal/diagnostics.buildTime=${BUILD_TIME}" \
		-mod=vendor \
		-o ./bin/tenerife \
		${PROJECT}/cmd/tenerife

RUN useradd -u 1001 tenerife

FROM gcr.io/distroless/base as runnable

ENV TENERIFE_HOST 0.0.0.0
ENV TENERIFE_PORT 8080
ENV TENERIFE_DIAGNOSTICS_PORT 9090

EXPOSE $TENERIFE_PORT $TENERIFE_DIAGNOSTICS_PORT

COPY --from=builder /app/bin/tenerife /bin/tenerife

COPY --from=builder /etc/passwd /etc/passwd

USER tenerife

CMD ["/bin/tenerife"]
