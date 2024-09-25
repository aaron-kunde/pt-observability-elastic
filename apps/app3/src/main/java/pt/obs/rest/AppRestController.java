package pt.obs.rest;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import pt.obs.db.DataEntity;
import pt.obs.db.DataRepository;
import pt.obs.kafka.AppKafkaProducer;

@Slf4j
@RestController
class AppRestController {

    private final Counter api1Counter;
    private final Counter api2Counter;
    private final Counter api3Counter;
    private final AppKafkaProducer kafkaProducer;
    private final DataRepository dataRepository;

    public AppRestController(AppKafkaProducer kafkaProducer,
                             DataRepository dataRepository,
                             @Value("${metrics.counter.api-1.name}") String api1CounterName,
                             @Value("${metrics.counter.api-2.name}") String api2CounterName,
                             @Value("${metrics.counter.api-3.name}") String api3CounterName) {
        this.kafkaProducer = kafkaProducer;
        this.dataRepository = dataRepository;
        this.api1Counter = Metrics.counter(api1CounterName, "it-1", "it-2", "type", "FACHLICH");
        this.api2Counter = Metrics.counter(api2CounterName, "type", "TECHNISCH");
        this.api3Counter = Metrics.counter(api3CounterName, "it-1", "it-22", "type", "keks");
    }

    @GetMapping("/api-1")
    void api1() {
        String apiName = "API 1";
        log.info("Calling {}", apiName);
        api1Counter.increment();

        double count = api1Counter.count();
        kafkaProducer.send(apiName, count);

        save("AppRestController-1: " + count);
    }

    @GetMapping("/api-2")
    void api2() {
        log.info("Calling API 2");
        api2Counter.increment();

        throw new RuntimeException("An unexpected error occurred");
    }

    @GetMapping("/api-3")
    void api3() {
        String apiName = "API 3";
        log.info("Calling {}", apiName);
        api3Counter.increment();

        save("AppRestController-3: " + api3Counter.count());
    }

    private void save(String data) {
        DataEntity entity = new DataEntity();
        entity.setData(data);
        log.info("Write data to database: {}", entity);
        dataRepository.save(entity);
    }
}
