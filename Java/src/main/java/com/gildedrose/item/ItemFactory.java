package com.gildedrose.item;

public class ItemFactory {
    private static final String agedBrieItem = "Aged Brie";
    private static final String sulfurasItem = "Sulfuras, Hand of Ragnaros";
    private static final String backstagePassItem = "Backstage passes to a TAFKAL80ETC concert";

    public static Item Item(Item item) {
        switch (item.name) {
            case agedBrieItem:
                return new ItemAgedBrie(item.name, item.sellIn, item.quality);
            case backstagePassItem:
                return new ItemBackstagePass(item.name, item.sellIn, item.quality);
            case sulfurasItem:
                return new ItemLegendary(item.name, item.sellIn, item.quality);
            default:
                return item;
        }
    }
}
