FROM alpine:3.2
ADD /cmd/main/fairfood .
RUN  mkdir -p  /opt/ff-service && mv  fairfood /opt/ff-service/ && mkdir /etc/dm && \
apk add --update ca-certificates openssl && update-ca-certificates && \
GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

WORKDIR /opt/ff-service/



EXPOSE 50051

ENTRYPOINT ["./fairfood"]