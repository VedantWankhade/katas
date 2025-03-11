package de.vedantwankha.kata.designpatterns.behavioral.strategy;

public class Map {
    private Router router;

    public Map(Router router) {
        this.router = router;
    }

    public void setRouter(Router router) {
        this.router = router;
    }

    public void route(String src, String dest) {
        System.out.printf("Path created from %s to %s for %s\n", src, dest, this.router.route());
    }

    public void route(String src, String dest, Router router) {
        System.out.printf("Path created from %s to %s for %s\n", src, dest, router.route());
    }
}
