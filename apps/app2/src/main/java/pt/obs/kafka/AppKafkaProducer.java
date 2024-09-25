package pt.obs.kafka;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;
@Slf4j
@Component
public class AppKafkaProducer {

    private final Counter topicCounter;
    private final KafkaTemplate<String, String> kafkaTemplate;

    @Value("${kafka.topic-out.name}") String topicOutName;
    @Value("${spring.application.name}") String applicationName;

    AppKafkaProducer(KafkaTemplate<String, String> kafkaTemplate,
                     @Value("${metrics.counter.topic-out.name}") String topicCounterName) {
        this.kafkaTemplate = kafkaTemplate;
        this.topicCounter = Metrics.counter(topicCounterName, "it-1", "it-2");
    }

    public void send(String apiName, double data) {
        log.info("Send data to topic {}: {}", topicOutName, data);
        kafkaTemplate.send(topicOutName, applicationName + ';' + apiName + ";data:" + data);
        topicCounter.increment();
    }
}
