# -*- coding: utf-8 -*-

from typing import Final


class Item:
    def __init__(self, name: str, sell_in: int, quality: int):
        self.name = name
        self.sell_in = sell_in
        self.quality = quality

    def __repr__(self):
        return "%s, %s, %s" % (self.name, self.sell_in, self.quality)

    def is_quality_below_average(self) -> bool:
        return self.quality < 50

    def increase_quality(self):
        self.quality += 1

    def decrease_quality(self):
        self.quality -= 1

    def reset_quality(self):
        self.quality = 0

    def has_zero_or_above_sell_in(self) -> bool:
        return self.sell_in >= 0


class GildedRose:
    __item_name_aged_brie: Final[str] = "Aged Brie"
    __item_name_backstage_pass: Final[str] = "Backstage passes to a TAFKAL80ETC concert"
    __item_name_sulfuras: Final[str] = "Sulfuras, Hand of Ragnaros"
    __known_item_names: list[str] = [
        __item_name_aged_brie,
        __item_name_backstage_pass,
        __item_name_sulfuras,
    ]

    def __init__(self, items: list[Item]) -> None:
        self.items = items

    def update_quality(self) -> None:
        for item in self.items:
            self.__update_item_quality(item)

    def __update_item_quality(self, item: Item) -> None:
        if not self.__is_known_item_name(item) and item.quality > 0:
            item.decrease_quality()

        if self.__is_known_item_name(item) and item.is_quality_below_average():
            item.increase_quality()
            if (
                item.name == self.__item_name_backstage_pass
                and item.is_quality_below_average()
                and item.sell_in < 11
            ):
                item.increase_quality()
                if item.sell_in < 6:
                    item.increase_quality()

        if item.name != self.__item_name_sulfuras:
            item.sell_in -= 1

        if item.has_zero_or_above_sell_in():
            return

        if item.name == self.__item_name_backstage_pass:
            item.reset_quality()

        if not self.__is_known_item_name(item) and item.quality > 0:
            item.decrease_quality()

        is_brie_below_average = (
            item.name == self.__item_name_aged_brie and item.is_quality_below_average()
        )
        if is_brie_below_average:
            item.increase_quality()

    def __is_known_item_name(self, item: Item) -> bool:
        return item.name in self.__known_item_names
