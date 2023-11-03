package com.gildedrose.item;

public class ItemAgedBrie implements IItem {
    private final Item item;

    public ItemAgedBrie(Item item) {
        this.item = item;
    }

    @Override
    public void update() {
        item.sellIn -= 1;

        if (item.quality > 50) return;
        item.quality += 2;
    }
}
