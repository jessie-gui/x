package xes

import (
	"testing"
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
