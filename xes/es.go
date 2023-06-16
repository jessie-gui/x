package xes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

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
		Body:  strings.NewReader(c.encodeBody(query)),
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

// encodeBody 编码请求体。
func (c *ESClient) encodeBody(query map[string]interface{}) string {
	body, err := json.Marshal(query)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
