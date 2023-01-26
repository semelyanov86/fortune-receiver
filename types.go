package main

type GetWebRequest interface {
	FetchBytes(typeFort int) ([]byte, error)
}

type FortuneResult struct {
	Content string
}
