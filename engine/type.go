package engine

type ParserFuncType func(contents []byte) ParserResult

type Request struct {
	Url        string
	ParserFunc ParserFuncType
}

type ParserResult struct {
	RequestSlice []Request
	ItemSlice    []Item
}

type Item struct {
	Url string
	//珍爱网有珍爱网ID,其它网站有其它网站的ID
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
