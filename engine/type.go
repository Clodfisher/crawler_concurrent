package engine

type ParserFuncType func(contents []byte) ParserResult

type Request struct {
	Url        string
	ParserFunc ParserFuncType
}

type ParserResult struct {
	RequestSlice []Request
	ItemSlice    []interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
