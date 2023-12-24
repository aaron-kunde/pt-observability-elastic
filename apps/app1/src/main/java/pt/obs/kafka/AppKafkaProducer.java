package pt.obs.kafka;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

import static pt.obs.kafka.KafkaConfiguration.TOPIC_OUT;
@Slf4j
@Component
public class AppKafkaProducer {

    private Counter topicCounter;

    private final KafkaTemplate<String, String> kafkaTemplate;

    AppKafkaProducer(KafkaTemplate<String, String> kafkaTemplate,
                     @Value("${metrics.counter.topic-out.name}") String topicCounterName) {
        this.kafkaTemplate = kafkaTemplate;
        this.topicCounter = Metrics.counter(topicCounterName, "it-1", "it-2");
    }

    public void send(String apiName, double data) {
        log.info(STR."Send data to topic \{TOPIC_OUT}: \{data}");
        kafkaTemplate.send(TOPIC_OUT, STR."app1;\{apiName};data:\{data}");
        topicCounter.increment();
    }
}
