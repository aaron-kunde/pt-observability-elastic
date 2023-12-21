package pt.obs.rest;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import pt.obs.kafka.AppKafkaProducer;

@Slf4j
@RestController
class AppRestController {

    private Counter api1Counter = Metrics.counter("app3.api.1.counter", "it-1", "it-2");
    private Counter api2Counter = Metrics.counter("app3.api.2.counter", "it-1", "it-3");
    private final AppKafkaProducer kafkaProducer;

    public AppRestController(AppKafkaProducer kafkaProducer) {
        this.kafkaProducer = kafkaProducer;
    }

    @GetMapping("/api-1")
    void api1() {
        String apiName = "API 1";
        log.info(STR."Calling \{apiName}");
        api1Counter.increment();
        kafkaProducer.send(apiName, api1Counter.count());
    }

    @GetMapping("/api-2")
    void api2() {
        log.info("Calling API 2");
        api2Counter.increment();

        throw new RuntimeException("An unexpected error occurred");
    }

}
