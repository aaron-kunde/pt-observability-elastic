FROM docker.io/eclipse-temurin:21.0.1_12-jre-alpine

ARG APP_NAME

RUN mkdir /opt/app

COPY build/libs/${APP_NAME}-0.0.1-SNAPSHOT.jar /opt/app/app.jar

RUN adduser -D -H -u 1000 app app \
    && chown -R app:app /opt/app

USER app

CMD java ${JAVA_OPTS} -jar /opt/app/app.jar
