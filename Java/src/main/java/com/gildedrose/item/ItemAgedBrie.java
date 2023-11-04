package com.gildedrose.item;

public class ItemAgedBrie extends Item {


    public ItemAgedBrie(String name, int sellIn, int quality) {
        super(name, sellIn, quality);
    }

    @Override
    public void update() {
        this.updateSellIn();
        this.updateQuality();
    }

    private void updateQuality() {
        if (this.quality < 50) this.quality += 2;
    }
}
