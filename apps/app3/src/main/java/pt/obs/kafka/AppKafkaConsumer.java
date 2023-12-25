package pt.obs.kafka;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestClient;
import pt.obs.db.DataEntity;
import pt.obs.db.DataRepository;

@Slf4j
@Component
public class AppKafkaConsumer {

    private Counter topicCounter;
    private final DataRepository dataRepository;
    private String restOutUrl;

    AppKafkaConsumer(DataRepository dataRepository,
                     @Value("${metrics.counter.topic-in.name}") String topicCounterName,
                     @Value("${rest.out.url}") String restOutUrl) {
        this.dataRepository = dataRepository;
        this.topicCounter = Metrics.counter(topicCounterName, "it-1", "it-2");
        this.restOutUrl = restOutUrl;
    }

    @KafkaListener(id="app3:1", topics = {"topic1", "topic3"})
    void listen(ConsumerRecord<String, String> record){
        log.info(STR."Fetch data from topic \{record.topic()}: \{record.key()}=\{record.value()}");
        topicCounter.increment();

        DataEntity data = new DataEntity();
        data.setData(STR."AppKafkaConsumer: \{topicCounter.count()}");
        log.info(STR."Write data to database: \{data}");
        dataRepository.save(data);

        log.info(STR."Call REST URL: \{restOutUrl}, result: \{RestClient.create(restOutUrl).get().retrieve().toEntity(String.class)}");
    }
}
