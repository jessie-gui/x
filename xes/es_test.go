package xes

import (
	"testing"
	"time"
)

func TestIndexDocument(t *testing.T) {
	esClient, err := NewESClient([]string{"http://localhost:9200"})
	if err != nil {
		t.Fatal("无法创建 Elasticsearch 客户端:", err)
	}

	// 索引文档
	document := map[string]interface{}{
		"title":  "Hello World",
		"author": "John Doe",
	}

	_, err = esClient.IndexDocument("my-index", "1", "_doc", document)
	if err != nil {
		t.Fatal("索引文档失败:", err)
	}
}

func TestSearchDocuments(t *testing.T) {
	esClient, err := NewESClient([]string{"http://localhost:9200"})
	if err != nil {
		t.Fatal("无法创建 Elasticsearch 客户端:", err)
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "hello",
			},
		},
	}

	searchRes, err := esClient.SearchDocuments("my-index", query)
	if err != nil {
		t.Fatal("搜索文档失败:", err)
	}

	t.Log("查询结果:", string(searchRes))
}

func TestDeleteDocument(t *testing.T) {
	esClient, err := NewESClient([]string{"http://localhost:9200"})
	if err != nil {
		t.Fatal("无法创建 Elasticsearch 客户端:", err)
	}

	deleteRes, err := esClient.DeleteDocument("my-index", "1")
	if err != nil {
		t.Fatal("删除文档失败:", err)
	}

	t.Log("文档已删除:", deleteRes.Status())
}

func TestInsertDocument(t *testing.T) {
	esClient, err := NewESClient([]string{"http://localhost:9200"})
	if err != nil {
		t.Fatal("无法创建 Elasticsearch 客户端:", err)
	}

	// 索引文档
	document := map[string]interface{}{
		"first_name": "Jane",
		"last_name":  "Smith",
		"age":        32,
		"about":      "I like to collect rock albums",
		"interests": []string{
			"music",
		},
	}

	_, err = esClient.IndexDocument("my-index", "3", "employee", document)
	if err != nil {
		t.Fatal("索引文档失败:", err)
	}
}

func TestScrollQuery(t *testing.T) {
	esClient, err := NewESClient([]string{"http://localhost:9200"})
	if err != nil {
		t.Fatal("无法创建 Elasticsearch 客户端:", err)
	}

	query := map[string]interface{}{
		"size": 1000,
		"sort": []map[string]interface{}{
			{
				"_doc": "asc",
			},
		},
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	searchRes, err := esClient.ScrollQuery("employee", query, 1*time.Minute)
	if err != nil {
		t.Fatal("游标查询失败:", err)
	}

	t.Log("查询结果:", string(searchRes))
}

func TestBulkDocuments(t *testing.T) {
	esClient, err := NewESClient([]string{"http://localhost:9200"})
	if err != nil {
		t.Fatal("无法创建 Elasticsearch 客户端:", err)
	}

	query := []map[string]interface{}{
		{
			"index": map[string]interface{}{
				"_id": 1,
			},
			"price":     10,
			"productID": "XHDK-A-1293-#fJ3",
		},
		{
			"index": map[string]interface{}{
				"_id": 2,
			},
			"price":     20,
			"productID": "KDKE-B-9947-#kL5",
		},
		{
			"index": map[string]interface{}{
				"_id": 3,
			},
			"price":     30,
			"productID": "JODL-X-1937-#pV7",
		},
		{
			"index": map[string]interface{}{
				"_id": 4,
			},
			"price":     30,
			"productID": "QQPX-R-3956-#aD8",
		},
	}

	res, err := esClient.BulkDocuments("my_store", query)
	if err != nil {
		t.Fatal("索引文档失败:", err)
	}

	t.Log("插入结果:", string(res))
}
