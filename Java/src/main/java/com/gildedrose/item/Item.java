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
        this.updateSellIn();
        this.updateQuality();
    }

    protected void updateSellIn() {
        this.sellIn -= 1;
    }

    private void updateQuality() {
        if (this.quality <= 0) return;
        this.quality -= 1;

        if (this.quality <= 0) return;
        if (this.sellIn < 0) this.quality -= 1;
    }


    @Override
    public String toString() {
        return this.name + ", " + this.sellIn + ", " + this.quality;
    }
}
