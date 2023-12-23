package pt.obs.db;

import org.springframework.data.repository.Repository;

public interface DataRepository extends Repository<DataEntity, Long> {
    DataEntity save(DataEntity data);
}
