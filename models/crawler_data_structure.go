package models

type CrawledData struct {
	DataSize int // the data size measuring unit should be bytes
	Data string // this is the texual version of the data. In future i might need to replace it with a file descriptor since the scrapped list can be in any size. Putting scrapped data all in memory at once can result in memory shortage.
	DataMap map[string]map[string]int // this is the mapped version can be useful for many tasks also faster to access.
	DataMapInterface interface{} // this is the optional type Data mapper that can be casted to different data types.
}

func (crawledData CrawledData) getDataSize() int {
	return crawledData.DataSize
}

func (crawledData CrawledData) getData() string {
	return crawledData.Data
}

