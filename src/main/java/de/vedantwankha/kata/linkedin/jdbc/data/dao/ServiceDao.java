package de.vedantwankha.kata.linkedin.jdbc.data.dao;

import de.vedantwankha.kata.linkedin.jdbc.data.entity.Service;
import de.vedantwankha.kata.linkedin.jdbc.data.util.DatabaseUtils;

import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.UUID;
import java.util.logging.Logger;

public class ServiceDao implements DAO<Service, UUID> {

    private static final Logger LOGGER = Logger.getLogger(ServiceDao.class.getName());
    private static final String GET_ALL = "select service_id, name, price from wisdom.services";
    private static final String GET_BY_ID = "select service_id, name, price from wisdom.services where service_id = ?";

    @Override
    public List<Service> getAll() {
        List<Service> services = new ArrayList<Service>();
        var conn = DatabaseUtils.getConnection();
        try(Statement stmt = conn.createStatement()) {
            var rs = stmt.executeQuery(GET_ALL);
            services = this.processResultSet(rs);
        } catch (SQLException e) {
            DatabaseUtils.handleSQLException("ServiceDao.getAll", e, LOGGER);
        }
        return services;
    }

    @Override
    public Service create(Service entity) {
        return null;
    }

    @Override
    public Optional<Service> getOne(UUID uuid) {
        var conn = DatabaseUtils.getConnection();
        try(PreparedStatement stmt = conn.prepareStatement(GET_BY_ID)) {
            stmt.setObject(1, uuid);
            var rs = stmt.executeQuery();
            var services = this.processResultSet(rs);
            if (services.isEmpty()) {
                return Optional.empty();
            } else {
                return Optional.of(services.get(0));
            }
        } catch (SQLException e) {
            DatabaseUtils.handleSQLException("ServiceDao.getOne", e, LOGGER);
            return Optional.empty();
        }
    }

    @Override
    public Service update(Service entity) {
        return null;
    }

    @Override
    public void delete(UUID uuid) {

    }

    private List<Service> processResultSet(ResultSet rs) throws SQLException {
        var services = new ArrayList<Service>();
        while (rs.next()) {
            var service = new Service();
            service.setServiceId((UUID) rs.getObject("service_id"));
            service.setName(rs.getString("name"));
            service.setPrice(rs.getBigDecimal("price"));
            services.add(service);
        }
        return services;
    }
}
