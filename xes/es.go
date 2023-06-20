package xes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// ESClient 是封装的 Elasticsearch 客户端。
type ESClient struct {
	client *elasticsearch.Client
}

// NewESClient 创建一个新的封装的 Elasticsearch 客户端。
func NewESClient(addresses []string) (*ESClient, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ESClient{
		client: client,
	}, nil
}

// IndexDocument 创建一个文档索引。
func (c *ESClient) IndexDocument(indexName string, docID string, documentType string, document map[string]interface{}) (*esapi.Response, error) {
	body, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}

	req := esapi.IndexRequest{
		Index:        indexName,
		DocumentID:   docID,
		DocumentType: documentType, // v8后取消了文档类型，默认为_doc
		Body:         strings.NewReader(string(body)),
		Refresh:      "true",
	}

	return req.Do(context.Background(), c.client)
}

// SearchDocuments 搜索文档。
func (c *ESClient) SearchDocuments(indexName string, query map[string]interface{}) ([]byte, error) {
	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  c.encodeBody(query),
	}

	res, err := req.Do(context.Background(), c.client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("搜索失败：%s", res.String())
	}

	result := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return json.Marshal(result)
}

// DeleteDocument 删除文档。
func (c *ESClient) DeleteDocument(indexName string, docID string) (*esapi.Response, error) {
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: docID,
		Refresh:    "true",
	}

	return req.Do(context.Background(), c.client)
}

// ScrollQuery 执行滚动查询(游标查询)
func (c *ESClient) ScrollQuery(indexName string, query map[string]interface{}, scrollTime time.Duration) ([]byte, error) {
	// 设置搜索选项，包括游标和时间间隔
	req := esapi.SearchRequest{
		Index:  []string{indexName},
		Body:   c.encodeBody(query),
		Scroll: scrollTime,
	}

	res, err := req.Do(context.Background(), c.client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("搜索失败：%s", res.String())
	}

	// 解析第一次滚动的结果
	result := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	scrollID := result["_scroll_id"].(string)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	var documents []map[string]interface{}
	for _, hit := range hits {
		document := hit.(map[string]interface{})["_source"].(map[string]interface{})
		documents = append(documents, document)
	}

	// 开始滚动查询
	for len(hits) > 0 {
		scrollReq := esapi.ScrollRequest{
			Scroll:   scrollTime,
			ScrollID: scrollID,
		}

		res, err := scrollReq.Do(context.Background(), c.client)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.IsError() {
			return nil, fmt.Errorf("滚动查询失败：%s", res.String())
		}

		// 解析滚动的结果
		result := make(map[string]interface{})
		if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
			return nil, err
		}

		scrollID = result["_scroll_id"].(string)
		hits = result["hits"].(map[string]interface{})["hits"].([]interface{})

		for _, hit := range hits {
			document := hit.(map[string]interface{})["_source"].(map[string]interface{})
			documents = append(documents, document)
		}
	}

	return json.Marshal(documents)
}

// BulkDocuments 批量操作文档。
func (c *ESClient) BulkDocuments(indexName string, query []map[string]interface{}) ([]byte, error) {
	reqBody := ""
	for _, q := range query {
		rb, err := json.Marshal(q)
		if err != nil {
			log.Fatal(err)
		}

		reqBody += string(rb) + "\n"
	}

	// 创建 _bulk 请求。
	req := esapi.BulkRequest{
		Index: indexName,
		Body:  strings.NewReader(reqBody),
	}

	// 执行 _bulk 请求。
	res, err := req.Do(context.Background(), c.client)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("批量操作失败：%s", res.String())
	}

	// 解析响应结果。
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	return json.Marshal(result)
}

// encodeBody 编码请求体。
func (c *ESClient) encodeBody(query map[string]interface{}) *strings.Reader {
	body, err := json.Marshal(query)
	if err != nil {
		log.Fatal(err)
	}
	return strings.NewReader(string(body))
}
