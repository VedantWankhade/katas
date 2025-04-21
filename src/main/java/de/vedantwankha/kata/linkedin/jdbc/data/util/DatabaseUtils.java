package de.vedantwankha.kata.linkedin.jdbc.data.util;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.util.logging.Logger;

public class DatabaseUtils {
    private static final String URL = "jdbc:postgresql://localhost:5432/";
    private static final String USERNAME = "postgres";
    private static final String PASSWORD = "rootpass";
    private static final Logger LOGGER = Logger.getLogger(DatabaseUtils.class.getName());
    private static final String exceptionFormat = "exception in %s, message: %s, code: %s";
    private static Connection connection;

    public static Connection getConnection() {
        if (connection == null) {
            synchronized (DatabaseUtils.class) {
                if (connection == null) {
                    try {
                        connection = DriverManager.getConnection(URL, USERNAME, PASSWORD);
                    } catch (SQLException e) {
                        handleSQLException("DatabaseUtils.getConnection", e, LOGGER);
                    }
                }
            }
        }
        return connection;
    }

     public static void handleSQLException(String method, SQLException e, Logger logger) {
        logger.warning(String.format(exceptionFormat, method, e.getMessage(), e.getErrorCode()));
        throw new RuntimeException(e);
    }
}
