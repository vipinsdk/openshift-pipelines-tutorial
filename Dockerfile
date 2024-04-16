FROM alpine AS build

RUN apk update \
    && apk add --no-cache go --repository http://dl-cdn.alpinelinux.org/alpine/edge/community/ make git gcc musl-dev build-base

WORKDIR /home/

COPY src /home/

RUN make

FROM alpine AS prod

ENV GODEBUG=madvdontneed=1

RUN mkdir -p /home/ \
    && apk update \
    && chown -R 1001:1001 /home/ \
    && chmod -R 777 /home/

WORKDIR /home/

COPY --chown=1001:1001 --from=build /home/httpserver /home/

USER 1001

CMD ./httpserver