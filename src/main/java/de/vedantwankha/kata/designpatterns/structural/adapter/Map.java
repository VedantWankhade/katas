package de.vedantwankha.kata.designpatterns.structural.adapter;

public class Map {
    private StandardImageProvider imageProvider;

    public Map(StandardImageProvider imageProvider) {
        this.imageProvider = imageProvider;
    }

    public void showImage(int x, int y) {
        System.out.println(this.imageProvider.getImage(x, y));
    }
}
