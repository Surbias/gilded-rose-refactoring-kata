package com.gildedrose.item;

public class IItemFactory {
    private static final String agedBrieItem = "Aged Brie";
    private static final String sulfurasItem = "Sulfuras, Hand of Ragnaros";
    private static final String backstagePassItem = "Backstage passes to a TAFKAL80ETC concert";

    public static IItem Item(Item item) {
        switch (item.name) {
            case agedBrieItem:
                return new ItemAgedBrie(item);
            case backstagePassItem:
                return new ItemBackstagePass(item);
            case sulfurasItem:
                return new ItemLegendary();
            default:
                return item;
        }
    }
}
