package pt.obs.rest;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.test.web.server.LocalServerPort;
import pt.obs.db.DataRepository;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class AppRestControllerRequestTest {

    private static final String BASE_URL = "http://localhost";

    @LocalServerPort
    private int port;

    @Autowired
    private TestRestTemplate restTemplate;

    @MockBean
    private DataRepository dataRepository;

    @Test
    void api1_should_return_200_and_no_content() {
        var entity = restTemplate.getForEntity(BASE_URL + ':' + port + '/' + "api-1", String.class);

        assertThat(entity.getStatusCode().is2xxSuccessful()).isTrue();
        assertThat(entity.hasBody()).isFalse();
    }
}
