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

    private final Counter topicCounter;
    private final DataRepository dataRepository;
    private final String restOutUrl;

    AppKafkaConsumer(DataRepository dataRepository,
                     @Value("${metrics.counter.topic-in.name}") String topicCounterName,
                     @Value("${rest.out.url}") String restOutUrl) {
        this.dataRepository = dataRepository;
        this.topicCounter = Metrics.counter(topicCounterName, "it-1", "it-2");
        this.restOutUrl = restOutUrl;
    }

    @KafkaListener(id="app1:1", topics = {"topic2", "topic3", "topic4"})
    void listen(ConsumerRecord<String, String> record){
        log.info("Fetch data from topic {}: {}={}", record.topic(), record.key(), record.value());
        topicCounter.increment();

        save("AppKafkaConsumer: " + topicCounter.count());

        log.info("Call REST URL: {}, result: {}", restOutUrl, RestClient.create(restOutUrl).get().retrieve().toEntity(String.class));
    }

    private void save(String data) {
        DataEntity entity = new DataEntity();
        entity.setData(data);
        log.info("Write data to database: {}", entity);
        dataRepository.save(entity);
    }
}
