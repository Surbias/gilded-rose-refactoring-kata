package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
		{"bar", 1, 2},
		{"baz", -1, 2},
		{"qux", 1, -2},
		{"qux", 1, 55},
	}

	expected := []*gildedrose.Item{
		{"foo", -1, 0},
		{"bar", 0, 1},
		{"baz", -2, 0},
		{"qux", 0, -2},
		{"qux", 0, 54},
	}

	if len(items) != len(expected) {
		t.Errorf("expected items should have %d items", len(expected))
		return
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if item.Name != expected[i].Name {
			t.Errorf("Name: Expected %s but got %s ", expected[i].Name, item.Name)
		}
		if item.SellIn != expected[i].SellIn {
			t.Errorf("SellIn: Expected %d but got %d ", expected[i].SellIn, item.SellIn)
		}
		if item.Quality != expected[i].Quality {
			t.Errorf("Quality: Expected %d but got %d ", expected[i].Quality, item.Quality)
		}
	}
}

func Test_AgedBrie(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", -1, 2},
		{"Aged Brie", 0, 0},
		{"Aged Brie", 1, -1},
		{"Aged Brie", 10, 30},
		{"Aged Brie", 10, 55},
	}

	expected := []*gildedrose.Item{
		{"Aged Brie", -2, 4},
		{"Aged Brie", -1, 2},
		{"Aged Brie", 0, 0},
		{"Aged Brie", 9, 31},
		{"Aged Brie", 9, 55},
	}

	if len(items) != len(expected) {
		t.Errorf("expected items should have %d items", len(expected))
		return
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if item.Name != expected[i].Name {
			t.Errorf("Name: Expected %s but got %s ", expected[i].Name, item.Name)
		}
		if item.SellIn != expected[i].SellIn {
			t.Errorf("SellIn: Expected %d but got %d ", expected[i].SellIn, item.SellIn)
		}
		if item.Quality != expected[i].Quality {
			t.Errorf("Quality: Expected %d but got %d ", expected[i].Quality, item.Quality)
		}
	}
}

func Test_BackstagePass(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 0, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 1},
		{"Backstage passes to a TAFKAL80ETC concert", 1, -1},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 1},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 50},
	}

	expected := []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 3},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 2},
		{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
	}

	if len(items) != len(expected) {
		t.Errorf("expected items should have %d items", len(expected))
		return
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if item.Name != expected[i].Name {
			t.Errorf("Name: Expected %s but got %s ", expected[i].Name, item.Name)
		}
		if item.SellIn != expected[i].SellIn {
			t.Errorf("SellIn: Expected %d but got %d ", expected[i].SellIn, item.SellIn)
		}
		if item.Quality != expected[i].Quality {
			t.Errorf("Quality: Expected %d but got %d ", expected[i].Quality, item.Quality)
		}
	}
}

func Test_Sulfuras(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 0},
		{"Sulfuras, Hand of Ragnaros", 1, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 1},
		{"Sulfuras, Hand of Ragnaros", 1, -1},
		{"Sulfuras, Hand of Ragnaros", -1, 1},
		{"Sulfuras, Hand of Ragnaros", -1, 50},
	}

	expected := []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 0},
		{"Sulfuras, Hand of Ragnaros", 1, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 1},
		{"Sulfuras, Hand of Ragnaros", 1, -1},
		{"Sulfuras, Hand of Ragnaros", -1, 1},
		{"Sulfuras, Hand of Ragnaros", -1, 50},
	}

	if len(items) != len(expected) {
		t.Errorf("expected items should have %d items", len(expected))
		return
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if item.Name != expected[i].Name {
			t.Errorf("Name: Expected %s but got %s ", expected[i].Name, item.Name)
		}
		if item.SellIn != expected[i].SellIn {
			t.Errorf("SellIn: Expected %d but got %d ", expected[i].SellIn, item.SellIn)
		}
		if item.Quality != expected[i].Quality {
			t.Errorf("Quality: Expected %d but got %d ", expected[i].Quality, item.Quality)
		}
	}
}
