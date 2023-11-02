package com.gildedrose;

public class ItemFactory {
    private static final String agedBrieItem = "Aged Brie";
    private static final String sulfurasItem = "Sulfuras, Hand of Ragnaros";
    private static final String backstagePassItem = "Backstage passes to a TAFKAL80ETC concert";

    public static IItem Item(Item item) {
        switch (item.name) {
            case agedBrieItem:
                return new AgedBrieItem(item);
            case backstagePassItem:
                return new BackstagePassItem(item);
            case sulfurasItem:
                return new LegendaryItem();
            default:
                return item;
        }
    }
}
