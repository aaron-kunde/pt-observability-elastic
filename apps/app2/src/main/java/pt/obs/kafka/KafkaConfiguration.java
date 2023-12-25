package pt.obs.kafka;

import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;

@Configuration
class KafkaConfiguration {

    @Value("${kafka.topic-out.name}") String topicOutName;

    @Bean
    public NewTopic topicOut() {
        return TopicBuilder.name(topicOutName)
                .replicas(1)
                .partitions(1)
                .build();
    }
}
