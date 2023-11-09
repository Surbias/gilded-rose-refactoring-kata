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
		if item.Name != agedBrie && item.Name != backstagePass {
			if item.Quality > 0 {
				if item.Name != sulfuras {
					item.Quality -= 1
				}
			}
		} else {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
				if item.Name == backstagePass {
					if item.SellIn < 11 {
						if item.Quality < 50 {
							item.Quality += 1
						}
					}
					if item.SellIn < 6 {
						if item.Quality < 50 {
							item.Quality += 1
						}
					}
				}
			}
		}

		if item.Name != sulfuras {
			item.SellIn -= 1
		}

		if item.SellIn < 0 {
			if item.Name != agedBrie {
				if item.Name != backstagePass {
					if item.Quality > 0 {
						if item.Name != sulfuras {
							item.Quality -= 1
						}
					}
				} else {
					item.Quality = 0
				}
			} else {
				if item.Quality < 50 {
					item.Quality += 1
				}
			}
		}
	}
}
