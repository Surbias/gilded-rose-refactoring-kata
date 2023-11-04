package com.gildedrose;

import com.gildedrose.item.Item;
import com.gildedrose.item.ItemFactory;

class GildedRose {
    Item[] items;


    public GildedRose(Item[] items) {
        this.items = new Item[items.length];
        for (int i = 0; i < items.length; i++) {
            this.items[i] = ItemFactory.Item(items[i]);
        }
    }

    public void updateQuality() {
        for (Item item : items) {
            item.update();
        }
    }

}
