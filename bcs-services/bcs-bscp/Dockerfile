FROM alpine
RUN apk --update --no-cache add ca-certificates bash vim curl
COPY bscp-ui /bscp-ui
COPY build/bk-bscp/bk-bscp-apiserver/bk-bscp-apiserver /bk-bscp/
COPY build/bk-bscp/bk-bscp-authserver/bk-bscp-authserver /bk-bscp/
COPY build/bk-bscp/bk-bscp-cacheservice/bk-bscp-cacheservice /bk-bscp/
COPY build/bk-bscp/bk-bscp-configserver/bk-bscp-configserver /bk-bscp/
COPY build/bk-bscp/bk-bscp-dataservice/bk-bscp-dataservice /bk-bscp/
COPY build/bk-bscp/bk-bscp-feedserver/bk-bscp-feedserver /bk-bscp/
ENTRYPOINT ["/bscp-ui"]
CMD []
