package com.gildedrose;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class GildedRoseTest {

    @Test
    void unknownItem() {
        Item[] items = new Item[]{
            new Item("foo", 0, 1),
            new Item("bar", 0, 2)
        };
        GildedRose app = new GildedRose(items);
        app.updateQuality();
        assertEquals("foo", app.items[0].name);
        assertEquals(0, app.items[0].quality);
        assertEquals(-1, app.items[0].sellIn);

        assertEquals("bar", app.items[1].name);
        assertEquals(0, app.items[1].quality);
        assertEquals(-1, app.items[1].sellIn);
    }

    @Test
    void sulfurasItem() {
        Item[] items = new Item[]{new Item("Sulfuras, Hand of Ragnaros", 0, 1)};

        GildedRose app = new GildedRose(items);
        app.updateQuality();
        assertEquals("Sulfuras, Hand of Ragnaros", app.items[0].name);
        assertEquals(1, app.items[0].quality);
        assertEquals(1, app.items[0].quality);
    }

    @Test
    void agedBrieItem() {
        Item[] items = new Item[]{new Item("Aged Brie", 0, 0)};

        GildedRose app = new GildedRose(items);
        app.updateQuality();
        assertEquals("Aged Brie", app.items[0].name);
        assertEquals(2, app.items[0].quality);
        assertEquals(-1, app.items[0].sellIn);
    }

    @Test
    void backstagePassItem() {
        Item[] items = new Item[]{new Item("Backstage passes to a TAFKAL80ETC concert", 0, 0)};

        GildedRose app = new GildedRose(items);
        app.updateQuality();
        assertEquals("Backstage passes to a TAFKAL80ETC concert", app.items[0].name);
        assertEquals(0, app.items[0].quality);
        assertEquals(-1, app.items[0].sellIn);
    }

    @Test
    void itemToString() {
        Item item = new Item("baz", 0, 0);
        assertEquals("baz, 0, 0", item.toString());
    }
}
