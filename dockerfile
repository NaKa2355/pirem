FROM golang:1.20.7-alpine3.18

RUN apk update && \
    apk add bash make && \
    apk add --upgrade grep

WORKDIR ./pirem
COPY ./ ./
RUN make build
RUN make install
COPY ./config/mock_config/piremd.json /etc/piremd.json

WORKDIR ../
RUN rm -rf ./pirem
RUN go clean --modcache
RUN mkdir /var/run/piremd
RUN chmod 755 /var/run/piremd

CMD ["pirem",  "daemon"]