package gildedrose

const (
	agedBrie      = "Aged Brie"
	backstagePass = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras      = "Sulfuras, Hand of Ragnaros"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		if item.Name == sulfuras {
			continue
		}

		if item.Name == backstagePass {
			item.SellIn -= 1
			if item.Quality < 50 {
				item.Quality += 1
			}
			if item.Quality < 50 {
				if item.SellIn < 11 {
					item.Quality += 1
				}
				if item.SellIn < 6 && item.Quality < 50 {
					item.Quality += 1
				}
			}
			if item.SellIn < 0 {
				item.Quality = 0
			}
			continue
		}

		if item.Name == agedBrie {
			item.SellIn -= 1
			if item.Quality < 50 {
				item.Quality += 1
			}

			if item.SellIn < 0 && item.Quality < 50 {
				item.Quality += 1
			}
			continue
		}

		if item.Name != agedBrie && item.Name != backstagePass {
			if item.Quality > 0 {
				item.Quality -= 1
			}
			item.SellIn -= 1
			if item.SellIn < 0 && item.Quality > 0 {
				item.Quality -= 1

			}
			continue
		}
	}
}
