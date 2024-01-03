FROM docker.io/alpine:3.19.0

ARG APP_NAME

RUN mkdir /opt/app

COPY build/${APP_NAME} /opt/app/app

RUN adduser -D -H -u 1000 app app \
    && chown -R app:app /opt/app

USER app

CMD /opt/app/app
