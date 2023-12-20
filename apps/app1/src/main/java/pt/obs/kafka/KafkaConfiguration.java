package pt.obs.kafka;

import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.boot.ApplicationRunner;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;
import org.springframework.kafka.core.KafkaTemplate;

@Configuration
class KafkaConfiguration {

    static final String TOPIC_1 = "topic1";

    @Bean
    public NewTopic topic() {
        return TopicBuilder.name(TOPIC_1)
                .replicas(1)
                .partitions(1)
                .build();
    }

    @Bean
    public ApplicationRunner runner(KafkaTemplate<String, String> template) {
        return args -> {
            template.send(TOPIC_1, "test");
        };
    }

}
