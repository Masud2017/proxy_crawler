package models

type CrawledData struct {
	DataSize int
	Data string // this is the texual version of the data. In future i might need to replace it with a file descriptor since the scrapped list can be in any size. Putting scrapped data all in memory at once can result in memory shortage.
	DataMap map[string]map[string]int // this is the mapped version can be useful for many tasks also faster to access.
}


