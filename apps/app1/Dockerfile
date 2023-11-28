FROM docker.io/curlimages/curl:8.4.0 as builder

RUN curl -L -O https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/latest/download/opentelemetry-javaagent.jar

FROM docker.io/eclipse-temurin:21.0.1_12-jre-alpine

RUN mkdir /opt/app

COPY --from=builder /home/curl_user/opentelemetry-javaagent.jar /opt/opentelemetry-javaagent.jar
COPY build/libs/app1-0.0.1-SNAPSHOT.jar /opt/app/app.jar

RUN adduser -D -H -u 1000 app app \
    && chown -R app:app /opt/app

USER app

CMD ["java", "-jar", "/opt/app/app.jar"]
