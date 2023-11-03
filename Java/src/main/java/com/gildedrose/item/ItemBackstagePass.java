package com.gildedrose.item;

public class ItemBackstagePass implements IItem {
    private final Item item;

    public ItemBackstagePass(Item item) {
        this.item = item;
    }


    @Override
    public void update() {
        item.sellIn -= 1;
        item.quality = 0;
    }
}
