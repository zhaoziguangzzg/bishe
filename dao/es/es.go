package es

import "github.com/elastic/go-elasticsearch/v7"

var EsClient *elasticsearch.Client

func InitEs() (err error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return
	}

	// 测试连接
	res, err := client.Info()
	if err != nil {
		return
	}
	defer res.Body.Close()

	EsClient = client
	return
}
