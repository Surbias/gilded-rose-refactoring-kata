# -*- coding: utf-8 -*-

from typing import Final


class Item:
    def __init__(self, name, sell_in, quality):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)


class GildedRose:
    
    __item_name_aged_brie: Final[str] = "Aged Brie"
    __item_name_backstage_pass: Final[str] = "Backstage passes to a TAFKAL80ETC concert"
    __item_name_sulfuras: Final[str] = "Sulfuras, Hand of Ragnaros"

    def __init__(self, items: list[Item]) -> None:
        self.items = items

    def update_quality(self) -> None:
        for item in self.items:
            if item.name != self.__item_name_aged_brie and item.name != self.__item_name_backstage_pass:
                if item.quality > 0:
                    if item.name != self.__item_name_sulfuras:
                        item.quality = item.quality - 1
            else:
                if item.quality < 50:
                    item.quality = item.quality + 1
                    if item.name == self.__item_name_backstage_pass:
                        if item.sell_in < 11:
                            if item.quality < 50:
                                item.quality = item.quality + 1
                        if item.sell_in < 6:
                            if item.quality < 50:
                                item.quality = item.quality + 1
            if item.name != self.__item_name_sulfuras:
                item.sell_in = item.sell_in - 1
            if item.sell_in < 0:
                if item.name != self.__item_name_aged_brie:
                    if item.name != self.__item_name_backstage_pass:
                        if item.quality > 0:
                            if item.name != self.__item_name_sulfuras:
                                item.quality = item.quality - 1
                    else:
                        item.quality = item.quality - item.quality
                else:
                    if item.quality < 50:
                        item.quality = item.quality + 1


