package pt.obs.kafka;

import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;

@Configuration
class KafkaConfiguration {

    static final String TOPIC_OUT = "topic3";

    @Bean
    public NewTopic topicOut() {
        return TopicBuilder.name(TOPIC_OUT)
                .replicas(1)
                .partitions(1)
                .build();
    }
}
