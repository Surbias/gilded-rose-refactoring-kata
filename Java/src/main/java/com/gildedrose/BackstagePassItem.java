package com.gildedrose;

public class BackstagePassItem implements IItem {
    private final Item item;

    public BackstagePassItem(Item item) {
        this.item = item;
    }


    @Override
    public void update() {
        item.sellIn -= 1;
        item.quality = 0;
    }
}
