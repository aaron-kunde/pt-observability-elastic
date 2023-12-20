package pt.obs.kafka;

import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

import static pt.obs.kafka.KafkaConfiguration.TOPIC_1;

@Component
public class AppKafkaProducer {

    private final KafkaTemplate<String, String> kafkaTemplate;

    public AppKafkaProducer(KafkaTemplate<String, String> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    public void send(String apiName, double data) {
        kafkaTemplate.send(TOPIC_1, STR."app1;\{apiName};data:\{data}");
    }
}
