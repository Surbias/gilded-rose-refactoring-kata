using NUnit.Framework;
using System.Collections.Generic;

namespace csharp
{
    [TestFixture]
    public class GildedRoseTest
    {
        [Test]
        public void item()
        {
            IList<Item> Items = new List<Item> {
                new Item { Name = "foo", SellIn = 0, Quality = 0 },
                new Item { Name = "foo", SellIn = 0, Quality = 1 },
                new Item { Name = "foo", SellIn = 1, Quality = 0 },
                new Item { Name = "foo", SellIn = -1, Quality = 10 },
                new Item { Name = "foo", SellIn = 0, Quality = -1 },
                new Item { Name = "foo", SellIn = 1, Quality = 50 },
                new Item { Name = "foo", SellIn = -5, Quality = 51 }
            };

            IList<Item> ExpectedItems = new List<Item> {
                new Item { Name = "foo", SellIn = -1, Quality = 0 },
                new Item { Name = "foo", SellIn = -1, Quality = 0 },
                new Item { Name = "foo", SellIn = 0, Quality = 0 },
                new Item { Name = "foo", SellIn = -2, Quality = 8 },
                new Item { Name = "foo", SellIn = -1, Quality = -1 },
                new Item { Name = "foo", SellIn = 0, Quality = 49 },
                new Item { Name = "foo", SellIn = -6, Quality = 49 }
            };

            GildedRose app = new GildedRose(Items);
            app.UpdateQuality();

            for (int i = 0; i < ExpectedItems.Count; i++)
            {
                Assert.AreEqual(ExpectedItems[i].Name, Items[i].Name);
                Assert.AreEqual(ExpectedItems[i].SellIn, Items[i].SellIn);
                Assert.AreEqual(ExpectedItems[i].Quality, Items[i].Quality);
            }

        }

    [Test]
        public void itemAgedBrie()
        {
            IList<Item> Items = new List<Item> {
                new Item { Name = "Aged Brie", SellIn = 0, Quality = 0 },
                new Item { Name = "Aged Brie", SellIn = 0, Quality = 1 },
                new Item { Name = "Aged Brie", SellIn = 1, Quality = 0 },
                new Item { Name = "Aged Brie", SellIn = -1, Quality = 10 },
                new Item { Name = "Aged Brie", SellIn = 0, Quality = -1 },
                new Item { Name = "Aged Brie", SellIn = 1, Quality = 50 },
                new Item { Name = "Aged Brie", SellIn = -5, Quality = 51 }
            };

            IList<Item> ExpectedItems = new List<Item> {
                new Item { Name = "Aged Brie", SellIn = -1, Quality = 2 },
                new Item { Name = "Aged Brie", SellIn = -1, Quality = 3 },
                new Item { Name = "Aged Brie", SellIn = 0, Quality = 1 },
                new Item { Name = "Aged Brie", SellIn = -2, Quality = 12 },
                new Item { Name = "Aged Brie", SellIn = -1, Quality = 1 },
                new Item { Name = "Aged Brie", SellIn = 0, Quality = 50 },
                new Item { Name = "Aged Brie", SellIn = -6, Quality = 51 }
            };

            GildedRose app = new GildedRose(Items);
            app.UpdateQuality();

            for (int i = 0; i < ExpectedItems.Count; i++)
            {
                Assert.AreEqual(ExpectedItems[i].Name, Items[i].Name);
                Assert.AreEqual(ExpectedItems[i].SellIn, Items[i].SellIn);
                Assert.AreEqual(ExpectedItems[i].Quality, Items[i].Quality);
            }

        }

    [Test]
        public void itemBackStagePass()
        {
            IList<Item> Items = new List<Item> {
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = 0, Quality = 0 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = 0, Quality = 1 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = 1, Quality = 0 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = -1, Quality = 10 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = 0, Quality = -1 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = 1, Quality = 50 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = -5, Quality = 51 }
            };

            IList<Item> ExpectedItems = new List<Item> {
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = -1, Quality = 0 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = -1, Quality = 0 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = 0, Quality = 3 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = -2, Quality = 0 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = -1, Quality = 0 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = 0, Quality = 50 },
                new Item { Name = "Backstage passes to a TAFKAL80ETC concert", SellIn = -6, Quality = 0 }
            };

            GildedRose app = new GildedRose(Items);
            app.UpdateQuality();

            for (int i = 0; i < ExpectedItems.Count; i++)
            {
                Assert.AreEqual(ExpectedItems[i].Name, Items[i].Name);
                Assert.AreEqual(ExpectedItems[i].SellIn, Items[i].SellIn);
                Assert.AreEqual(ExpectedItems[i].Quality, Items[i].Quality);
            }

        }

    [Test]
        public void itemSulfuras()
        {
            IList<Item> Items = new List<Item> {
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 0, Quality = 0 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 0, Quality = 1 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 1, Quality = 0 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = -1, Quality = 10 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 0, Quality = -1 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 1, Quality = 50 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = -5, Quality = 51 }
            };

            IList<Item> ExpectedItems = new List<Item> {
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 0, Quality = 0 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 0, Quality = 1 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 1, Quality = 0 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = -1, Quality = 10 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 0, Quality = -1 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = 1, Quality = 50 },
                new Item { Name = "Sulfuras, Hand of Ragnaros", SellIn = -5, Quality = 51 }
            };

            GildedRose app = new GildedRose(Items);
            app.UpdateQuality();

            for (int i = 0; i < ExpectedItems.Count; i++)
            {
                Assert.AreEqual(ExpectedItems[i].Name, Items[i].Name);
                Assert.AreEqual(ExpectedItems[i].SellIn, Items[i].SellIn);
                Assert.AreEqual(ExpectedItems[i].Quality, Items[i].Quality);
            }

        }
    }
}
