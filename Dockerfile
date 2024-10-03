FROM alpine:latest as builder
ARG TARGETPLATFORM
RUN echo "I'm building for $TARGETPLATFORM"

RUN apk add --no-cache gzip && \
    mkdir /clash.meta-config && \
    wget -O /clash.meta-config/geoip.metadb https://fastly.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geoip.metadb && \
    wget -O /clash.meta-config/geosite.dat https://fastly.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geosite.dat && \
    wget -O /clash.meta-config/geoip.dat https://fastly.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geoip.dat

COPY docker/file-name.sh /clash.meta/file-name.sh
WORKDIR /clash.meta
COPY bin/ bin/
RUN FILE_NAME=`sh file-name.sh` && echo $FILE_NAME && \
    FILE_NAME=`ls bin/ | egrep "$FILE_NAME.*"|awk NR==1` && echo $FILE_NAME && \
    mv bin/$FILE_NAME clash.meta.gz && gzip -d clash.meta.gz && echo "$FILE_NAME" > /clash.meta-config/test
FROM alpine:latest
LABEL org.opencontainers.image.source="https://github.com/kitty314/1.17.0"

RUN apk add --no-cache ca-certificates tzdata iptables

VOLUME ["/root/.config/clash.meta/"]

COPY --from=builder /clash.meta-config/ /root/.config/clash.meta/
COPY --from=builder /clash.meta/clash.meta /clash.meta
RUN chmod +x /clash.meta
ENTRYPOINT [ "/clash.meta" ]
