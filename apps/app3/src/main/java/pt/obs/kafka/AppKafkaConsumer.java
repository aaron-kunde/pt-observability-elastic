package pt.obs.kafka;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class AppKafkaConsumer {

    private Counter topicCounter = Metrics.counter("app3.topic.in.counter", "it-1", "it-2");

    @KafkaListener(id="app3:1", topics = {"topic1", "topic2"})
    void listen(ConsumerRecord<String, String> record){
        log.info(STR."Fetch data from topic \{record.topic()}: \{record.key()}=\{record.value()}");
        topicCounter.increment();
    }
}
