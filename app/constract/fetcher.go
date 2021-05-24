package constract


type Fetcher interface {
	Fetch(totalPage int) []map[string]interface{}
}