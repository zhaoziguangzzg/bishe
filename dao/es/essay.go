package es

import (
	"bishe/model"
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gin-gonic/gin"
)

// 在es添加essay
func AddEssayEsDoc(c *gin.Context, essay model.Essay) (err error) {
	jsonData, err := json.Marshal(essay)
	if err != nil {
		return
	}

	req := esapi.IndexRequest{
		Index:      "essay_info",
		DocumentID: strconv.Itoa(essay.Id),
		Body:       bytes.NewReader(jsonData),
	}

	res, err := req.Do(c, EsClient)
	if err != nil {
		return
	}

	/*
			{
		    "_index": "user_info",
		    "_type": "_doc",
		    "_id": "11",
		    "_version": 5,
		    "result": "updated",
		    "_shards": {
		        "total": 2,
		        "successful": 1,
		        "failed": 0
		    },
		    "_seq_no": 8,
		    "_primary_term": 1
		}
	*/

	if res.IsError() {
		err = errors.New("response err")
		return
	}

	return
}

// 从es获取文章
func GetEssayFromEs(cid int, word string, from int, size int) (essayList []model.Essay, err error) {
	_, wordErr := strconv.Atoi(word)
	var req map[string]interface{}
	if wordErr == nil {
		req = map[string]interface{}{
			"from": from,
			"size": size,
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"should": []map[string]interface{}{
						{
							"wildcard": map[string]interface{}{
								"title": "*" + word + "*",
							},
						},
						{
							"wildcard": map[string]interface{}{
								"content": "*" + word + "*",
							},
						},
					},
					"minimum_should_match": 1,
					"filter": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"circleId": cid,
							},
						},
					},
				},
			},
		}
	} else {
		req = map[string]interface{}{
			"from": from,
			"size": size,
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"should": []map[string]interface{}{
						{
							"match": map[string]interface{}{
								"title": word,
							},
						},
						{
							"match": map[string]interface{}{
								"content": word,
							},
						},
					},
					"minimum_should_match": 1,
					"filter": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"circleId": cid,
							},
						},
					},
				},
			},
		}
	}

	reqStr, err := json.Marshal(req)
	if err != nil {
		return
	}

	res, err := EsClient.Search(
		EsClient.Search.WithIndex("essay_info"),
		EsClient.Search.WithBody(strings.NewReader(string(reqStr))),
	)

	if err != nil {
		return
	}

	if res.IsError() {
		err = errors.New("response err")
		return
	}

	var result model.GetEsEssayResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return

	}

	essayList = make([]model.Essay, 0)
	if result.Hits.Total.Value == 0 {
		return
	}

	for _, v := range result.Hits.Hits {
		essayList = append(essayList, v.Source)
	}

	return
}
