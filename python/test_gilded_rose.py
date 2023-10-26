# -*- coding: utf-8 -*-
import unittest

from gilded_rose import Item, GildedRose


class GildedRoseTest(unittest.TestCase):
    def test_foo(self):
        items = [
            Item(name="foo", sell_in=0, quality=0),
        ]
        gilded_rose = GildedRose(items)
        gilded_rose.update_quality()
        self.assertEqual("foo", items[0].name)

    def test_some_item_should_decrease_sell_in_and_quality(self):
        items = [
            Item(name="some item", sell_in=10, quality=80),
            Item(name="some item", sell_in=-1, quality=80),
        ]
        gilded_rose = GildedRose(items)
        gilded_rose.update_quality()
        
        self.assertEqual(79, items[0].quality)
        self.assertEqual(9, items[0].sell_in)

        self.assertEqual(78, items[1].quality)
        self.assertEqual(-2, items[1].sell_in)

    def test_aged_brie(self):
        name = "Aged Brie"
        items = [
            Item(name=name, sell_in=10, quality=80),
            Item(name=name, sell_in=10, quality=49),
            Item(name=name, sell_in=-1, quality=48),
        ]
        gilded_rose = GildedRose(items)
        gilded_rose.update_quality()
        
        self.assertEqual(80, items[0].quality)
        self.assertEqual(9, items[0].sell_in)

        self.assertEqual(50, items[1].quality)
        self.assertEqual(9, items[1].sell_in)

        self.assertEqual(50, items[2].quality)
        self.assertEqual(-2, items[2].sell_in)

    def test_backstage_pass(self):
        name = "Backstage passes to a TAFKAL80ETC concert"
        items = [
            Item(name=name, sell_in=10, quality=80),
            Item(name=name, sell_in=10, quality=49),
            Item(name=name, sell_in=5, quality=49),
            Item(name=name, sell_in=-1, quality=49),
            Item(name=name, sell_in=10, quality=48),
            Item(name=name, sell_in=1, quality=0),
        ]
        gilded_rose = GildedRose(items)
        gilded_rose.update_quality()
        
        self.assertEqual(80, items[0].quality)
        self.assertEqual(9, items[0].sell_in)

        self.assertEqual(50, items[1].quality)
        self.assertEqual(9, items[1].sell_in)
        
        self.assertEqual(50, items[2].quality)
        self.assertEqual(4, items[2].sell_in)

        self.assertEqual(0, items[3].quality)
        self.assertEqual(-2, items[3].sell_in)

        self.assertEqual(50, items[4].quality)
        self.assertEqual(9, items[4].sell_in)

        self.assertEqual(3, items[5].quality)
        self.assertEqual(0, items[5].sell_in)

    def test_item_repr(self):
        item = Item(name="foo", sell_in=0, quality=0),
        self.assertEqual("(foo, 0, 0,)", repr(item))

if __name__ == "__main__":
    unittest.main()
