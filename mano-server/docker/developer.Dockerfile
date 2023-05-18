FROM alpine

WORKDIR /app

ADD mano-server/bin/mano-server /app/bin/mano-server

ENTRYPOINT bin/mano-server