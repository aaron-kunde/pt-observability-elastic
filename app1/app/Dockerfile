FROM docker.io/curlimages/curl:8.4.0 as builder

RUN curl -L -O https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/latest/download/opentelemetry-javaagent.jar

FROM docker.io/eclipse-temurin:21.0.1_12-jdk-alpine

COPY --from=builder /home/curl_user/opentelemetry-javaagent.jar /opt/opentelemetry-javaagent.jar
COPY build/libs/app-0.0.1-SNAPSHOT.jar /opt/app.jar

CMD ["java", "-jar", "/opt/app.jar"]
