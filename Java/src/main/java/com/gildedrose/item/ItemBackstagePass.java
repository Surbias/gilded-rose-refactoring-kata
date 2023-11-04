package com.gildedrose.item;

public class ItemBackstagePass extends Item {


    public ItemBackstagePass(String name, int sellIn, int quality) {
        super(name, sellIn, quality);
    }

    @Override
    public void update() {
        this.updateSellIn();
        this.updateQuality();
    }

    private void updateQuality() {
        this.quality = 0;
    }
}
