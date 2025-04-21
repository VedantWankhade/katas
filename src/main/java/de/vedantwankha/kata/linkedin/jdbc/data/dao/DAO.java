package de.vedantwankha.kata.linkedin.jdbc.data.dao;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

public interface DAO<T, ID extends UUID> {
    List<T> getAll();
    T create(T entity);
    Optional<T> getOne(ID id);
    T update(T entity);
    void delete(ID id);
}
