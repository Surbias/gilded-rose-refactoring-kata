package gildedrose

type ItemUpdater interface {
	UpdateQuality() *Item
}

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) UpdateQuality() *Item {
	i.SellIn -= 1

	if i.Quality > 0 {
		i.Quality -= 1
	}

	if i.SellIn < 0 && i.Quality > 0 {
		i.Quality -= 1
	}

	return i
}

type ItemLegendary struct {
	Item
}

func (i *ItemLegendary) UpdateQuality() *Item {
	return &i.Item
}

type ItemAgedBrie struct {
	Item
}

func (i *ItemAgedBrie) UpdateQuality() *Item {
	i.SellIn -= 1

	if i.Quality < 50 {
		i.Quality += 1
	}

	if i.SellIn < 0 && i.Quality < 50 {
		i.Quality += 1
	}
	return &i.Item
}

type ItemBackstagePass struct {
	Item
}

func (i *ItemBackstagePass) UpdateQuality() *Item {
	i.SellIn -= 1

	if i.Quality < 50 {
		i.Quality += 1
	}

	if i.SellIn < 11 && i.Quality < 50 {
		i.Quality += 1
	}

	if i.SellIn < 6 && i.Quality < 50 {
		i.Quality += 1
	}

	if i.SellIn < 0 {
		i.Quality = 0
	}

	return &i.Item
}

func createItemUpdater(item *Item) ItemUpdater {
	const (
		agedBrie      = "Aged Brie"
		backstagePass = "Backstage passes to a TAFKAL80ETC concert"
		sulfuras      = "Sulfuras, Hand of Ragnaros"
	)

	switch item.Name {
	case sulfuras:
		return &ItemLegendary{*item}
	case agedBrie:
		return &ItemAgedBrie{*item}
	case backstagePass:
		return &ItemBackstagePass{*item}
	default:
		return item
	}

}

func UpdateQuality(items []*Item) {
	for i, item := range items {
		itemUpdater := createItemUpdater(item)
		items[i] = itemUpdater.UpdateQuality()
	}
}
