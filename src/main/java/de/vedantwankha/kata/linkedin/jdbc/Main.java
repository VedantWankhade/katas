package de.vedantwankha.kata.linkedin.jdbc;

import de.vedantwankha.kata.linkedin.jdbc.data.dao.ServiceDao;
import de.vedantwankha.kata.linkedin.jdbc.data.entity.Service;

import java.util.UUID;

public class Main {
    public static void main(String[] args) {
        ServiceDao serviceDao = new ServiceDao();
        var services = serviceDao.getAll();
        System.out.println("-------------Services---------------");
        System.out.println("--------------Get All---------------");
        services.forEach(System.out::println);

        var service1 = serviceDao.getOne(services.getFirst().getServiceId());
        var service2 = serviceDao.getOne(new UUID(2, 2));
        System.out.println("-------Get One-----------");
        System.out.println(service1.orElseGet(Service::new));
        System.out.println(service2.orElseGet(Service::new));
    }
}
