package com.gildedrose;

public class AgedBrieItem implements IItem {
    private final Item item;

    public AgedBrieItem(Item item) {
        this.item = item;
    }

    @Override
    public void update() {
        item.sellIn -= 1;

        if (item.quality > 50) return;
        item.quality += 2;
    }
}
