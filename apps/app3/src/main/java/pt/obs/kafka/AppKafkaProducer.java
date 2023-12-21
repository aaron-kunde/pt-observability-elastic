package pt.obs.kafka;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

import static pt.obs.kafka.KafkaConfiguration.TOPIC_OUT;
@Slf4j
@Component
public class AppKafkaProducer {

    private Counter topicCounter = Metrics.counter("app3.topic.out.counter", "it-1", "it-2");

    private final KafkaTemplate<String, String> kafkaTemplate;

    AppKafkaProducer(KafkaTemplate<String, String> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    public void send(String apiName, double data) {
        log.info(STR."Send data to topic \{TOPIC_OUT}: \{data}");
        kafkaTemplate.send(TOPIC_OUT, STR."app3;\{apiName};data:\{data}");
        topicCounter.increment();
    }
}
