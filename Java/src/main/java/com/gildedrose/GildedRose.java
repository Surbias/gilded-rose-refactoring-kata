package com.gildedrose;

import com.gildedrose.item.Item;
import com.gildedrose.item.ItemFactory;

class GildedRose {
    Item[] items;


    public GildedRose(Item[] items) {
        this.items = items;
    }

    public void updateQuality() {
        for (Item item : items) {
            ItemFactory
                .Item(item)
                .update();
        }
    }

}
