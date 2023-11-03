package com.gildedrose.item;

public class Item {

    public String name;

    public int sellIn;

    public int quality;

    public Item(String name, int sellIn, int quality) {
        this.name = name;
        this.sellIn = sellIn;
        this.quality = quality;
    }

    public void update() {
        this.decreaseQuality();

        this.sellIn -= 1;
        if (this.sellIn > 0) return;

        this.decreaseQuality();
    }

    private void decreaseQuality() {
        if (this.quality > 0) this.quality -= 1;
    }

    @Override
    public String toString() {
        return this.name + ", " + this.sellIn + ", " + this.quality;
    }
}
