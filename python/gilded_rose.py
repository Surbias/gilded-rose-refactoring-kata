# -*- coding: utf-8 -*-

from typing import Final


class Item:
    def __init__(self, name: str, sell_in: int, quality: int):
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
            self.__update_item_sell_in(item)
            self.__update_item_quality(item)

    def __update_item_sell_in(self, item: Item) -> None:
        if item.name != self.__item_name_sulfuras:
            item.sell_in = item.sell_in - 1
        return item

    def __update_item_quality(self, item: Item) -> None:
        if item.name not in [
            self.__item_name_aged_brie,
            self.__item_name_backstage_pass,
        ]:
            if item.quality > 0:
                if item.name != self.__item_name_sulfuras:
                    item.quality = self.__decrease_quality(item.quality)
        else:
            if item.quality < 50:
                item.quality = self.__increase_quality(item.quality)
                if item.name == self.__item_name_backstage_pass:
                    if item.sell_in < 11:
                        if item.quality < 50:
                            item.quality = self.__increase_quality(item.quality)
                    if item.sell_in < 6:
                        if item.quality < 50:
                            item.quality = self.__increase_quality(item.quality)
        if item.sell_in < 0:
            if item.name != self.__item_name_aged_brie:
                if item.name != self.__item_name_backstage_pass:
                    if item.quality > 0:
                        if item.name != self.__item_name_sulfuras:
                            item.quality = self.__decrease_quality(item.quality)
                else:
                    item.quality = item.quality - item.quality
            else:
                if item.quality < 50:
                    item.quality = self.__increase_quality(item.quality)

    def __increase_quality(self, quality: int) -> int:
        return quality + 1

    def __decrease_quality(self, quality: int) -> int:
        return quality - 1
