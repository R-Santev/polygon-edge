FROM rsantev/polygon-edge:latest

WORKDIR /app

USER root

RUN apk --no-cache add ca-certificates jq

COPY ./polygon-edge.sh ./

# Expose json-rpc, libp2p, grpc and prometheus ports
EXPOSE 8545 9632 1478 5001

ENTRYPOINT ["./polygon-edge.sh"]
