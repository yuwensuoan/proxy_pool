package constract


type Fetcher interface {
	Fetch(maxPage int) []map[string]interface{}
}