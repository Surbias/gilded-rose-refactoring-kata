package com.gildedrose;

import com.gildedrose.item.Item;
import com.gildedrose.item.IItemFactory;

class GildedRose {
    Item[] items;


    public GildedRose(Item[] items) {
        this.items = items;
    }

    public void updateQuality() {
        for (Item item : items) {
            IItemFactory
                .Item(item)
                .update();
        }
    }

}
