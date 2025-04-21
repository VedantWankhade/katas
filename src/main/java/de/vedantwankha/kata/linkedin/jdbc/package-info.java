/**
 * <h3>Terminology</h3>
 * <ul>
 *     <li><code>DriverManager</code> - class that interacts with the driver for creating connections</li>
 *     <li><code>DataSource</code> - modern class that interacts with the driver for creating connections</li>
 *     <li><code>Connection</code> - the class that the developer interacts with, that manages the actual communication between the client and the database server</li>
 *     <li><code>Statement</code> - </li>
 *     <li><code>PreparedStatement</code> - </li>
 *     <li><code>ResultSet</code> - </li>
 *     <li>DTO - Single domain data (not necessarily single table - could be joins also): An object to transfer data</li>
 *     <li>DAO - Data Access Object: An object that interacts with database to handle operations on a data object (like DTO). Acts as an interface for data access logic</li>
 * </ul>
 *
 * <h3>Exceptions</h3>
 * <ul>
 *     <li>All JDBC operations throw a SQLException</li>
 *     <li>They are checked exceptions so you have to catch them</li>
 *     <li>SQLState code is available</li>
 *     <li>Vendor specific ErrorCode are available</li>
 * </ul>
 */
package de.vedantwankha.kata.linkedin.jdbc;