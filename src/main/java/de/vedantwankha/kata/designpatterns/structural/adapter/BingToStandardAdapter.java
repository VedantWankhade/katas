package de.vedantwankha.kata.designpatterns.structural.adapter;

public class BingToStandardAdapter implements StandardImageProvider {
    private BingMap m = new BingMap();

    @Override
    public String getImage(int x, int y) {
        return m.image(y, x);
    }
}
