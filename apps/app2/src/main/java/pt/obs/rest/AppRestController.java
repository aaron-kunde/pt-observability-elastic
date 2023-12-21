package pt.obs.rest;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.Metrics;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RestController
class AppRestController {

    private Counter api1Counter = Metrics.counter("app2.api.1.counter", "it-1", "it-2");
    private Counter api2Counter = Metrics.counter("app2.api.2.counter", "it-1", "it-3");

    @GetMapping("/api-1")
    void api1() {
        log.info("Calling API 1");
        api1Counter.increment();
    }

    @GetMapping("/api-2")
    void api2() {
        log.info("Calling API 2");
        api2Counter.increment();

        throw new RuntimeException("An unexpected error occurred");
    }

}
