package tst.obs.rest;

import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RestController
public class AppRestController {

    @GetMapping("/api-1")
    void api1() {
        log.info("Calling api 1");
    }
}
