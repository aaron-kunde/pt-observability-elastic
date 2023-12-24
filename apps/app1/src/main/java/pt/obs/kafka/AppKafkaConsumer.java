package pt.obs.kafka;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;
import pt.obs.db.DataEntity;
import pt.obs.db.DataRepository;

@Slf4j
@Component
public class AppKafkaConsumer {

    private Counter topicCounter = Metrics.counter("app1.topic.in.counter", "it-1", "it-2");
    private final DataRepository dataRepository;

    AppKafkaConsumer(DataRepository dataRepository) {
        this.dataRepository = dataRepository;
    }

    @KafkaListener(id="app1:1", topics = {"topic2", "topic3"})
    void listen(ConsumerRecord<String, String> record){
        log.info(STR."Fetch data from topic \{record.topic()}: \{record.key()}=\{record.value()}");
        topicCounter.increment();
        DataEntity data = new DataEntity();
        data.setData(STR."AppKafkaConsumer: \{topicCounter.count()}");
        dataRepository.save(data);
    }
}
