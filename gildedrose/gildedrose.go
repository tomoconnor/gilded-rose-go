package gildedrose

import "strings"

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		if strings.HasPrefix(item.Name, "Conjured") {
			item.Quality -= 2
		} else {

			switch item.Name {
			case "Aged Brie":
				updateAgedBrie(item)
			case "Backstage passes to a TAFKAL80ETC concert":
				updateBackstagePasses(item)
			case "Sulfuras, Hand of Ragnaros":
				// Do nothing for Sulfuras, it doesn't change.
			default:
				updateDefaultItem(item)
			}

			// Adjust the sell-in value for all items except Sulfuras.
			if item.Name != "Sulfuras, Hand of Ragnaros" {
				item.SellIn--
			}

			// Check post-sell-in adjustments for all items.
			if item.SellIn < 0 {
				postSellInAdjustment(item)
			}
		}
	}
}

func updateAgedBrie(item *Item) {
	if item.Quality < 50 {
		item.Quality++
	}
}

func updateBackstagePasses(item *Item) {
	if item.Quality < 50 {
		item.Quality++
		if item.SellIn < 11 {
			item.Quality++
		}
		if item.SellIn < 6 {
			item.Quality++
		}
	}
}

func updateDefaultItem(item *Item) {
	if item.Quality > 0 {
		item.Quality--
	}
}

func postSellInAdjustment(item *Item) {
	switch item.Name {
	case "Aged Brie":
		updateAgedBrie(item)
	case "Backstage passes to a TAFKAL80ETC concert":
		item.Quality = 0
	default:
		updateDefaultItem(item)
	}
}
