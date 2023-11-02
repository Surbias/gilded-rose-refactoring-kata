package com.gildedrose;

public class Item implements IItem {

    public String name;

    public int sellIn;

    public int quality;

    public Item(String name, int sellIn, int quality) {
        this.name = name;
        this.sellIn = sellIn;
        this.quality = quality;
    }

    public void update() {
        this.sellIn -= 1;
        this.quality = this.quality > 0 ? this.quality - 1 : this.quality;
        if (this.sellIn < 0)
            this.quality = this.quality > 0 ? this.quality - 1 : this.quality;
    }

    @Override
    public String toString() {
        return this.name + ", " + this.sellIn + ", " + this.quality;
    }
}
