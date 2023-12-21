package pt.obs.kafka;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import static pt.obs.kafka.KafkaConfiguration.TOPIC_IN;

@Slf4j
@Component
public class AppKafkaConsumer {

    private Counter topicCounter = Metrics.counter("app2.topic.in.counter", "it-1", "it-2");

    @KafkaListener(id="app2:1", topics = TOPIC_IN)
    void onIn(String data){
        log.info(STR."Fetch data from topic \{TOPIC_IN}: \{data}");
        topicCounter.increment();
    }
}
