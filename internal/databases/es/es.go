// Package es 提供 Elasticsearch 搜索引擎的客户端封装，
// 支持索引创建、文档增删改查和批量操作。
package es

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.uber.org/zap"

	"x-HanJin/pkg/log"
)

// Conf Elasticsearch 连接配置
type Conf struct {
	Addresses string `json:"addresses"` // 服务地址
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
}

// DocContent 文档内容，支持 upsert 操作
type DocContent struct {
	Doc         string `json:"doc"`           // 文档内容
	DocAsUpsert bool   `json:"doc_as_upsert"` // 是否作为 upsert 操作
}

// ElasticsearchClient Elasticsearch 客户端封装
type ElasticsearchClient struct {
	Client *elasticsearch.Client
}

// NewElasticsearchClient 创建并验证 Elasticsearch 客户端连接
func NewElasticsearchClient(addresses, username, password string) (*ElasticsearchClient, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 跳过 TLS 证书验证（生产环境建议关闭）
		},
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ResponseHeaderTimeout: 10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
	}

	cfg := elasticsearch.Config{
		Addresses: []string{addresses},
		Username:  username,
		Password:  password,
		Transport: tr,
	}

	esCli, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Error("创建 Elasticsearch 客户端失败", zap.Error(err))
		return nil, err
	}

	// 验证连接是否正常
	res, err := esCli.Info()
	if err != nil {
		log.Error("获取 Elasticsearch 信息失败", zap.Error(err))
		return nil, fmt.Errorf("error getting Elasticsearch info: %w", err)
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Error("关闭响应体失败", zap.Error(err))
		}
	}(res.Body)

	if res.IsError() {
		log.Error("Elasticsearch 返回错误状态", zap.String("status", res.Status()))
		return nil, fmt.Errorf("error response from Elasticsearch: %s", res.Status())
	}

	log.Info("Elasticsearch 客户端连接成功", zap.String("address", addresses))
	return &ElasticsearchClient{Client: esCli}, nil
}

// Init 初始化 Elasticsearch 客户端
func Init(addresses, userName, password string) {
	es, err := NewElasticsearchClient(addresses, userName, password)
	if err != nil {
		log.Error("Elasticsearch 初始化失败", zap.Error(err))
		panic(err)
	}
	_ = es // TODO: 存储为全局实例，初始化索引
}

// Search 执行搜索查询
func (es *ElasticsearchClient) Search(index, query string) (*esapi.Response, error) {
	res, err := es.Client.Search(
		es.Client.Search.WithContext(context.Background()),
		es.Client.Search.WithIndex(index),
		es.Client.Search.WithBody(strings.NewReader(query)),
		es.Client.Search.WithTrackTotalHits(true),
		es.Client.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("search query failed: %w", err)
	}
	return res, nil
}

// Index 向指定索引插入单条文档
func (es *ElasticsearchClient) Index(index, docId, jsonMessage string) (*esapi.Response, error) {
	data := []byte(jsonMessage)
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: docId,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}
	resp, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return nil, fmt.Errorf("index document failed: %w", err)
	}
	return resp, nil
}

// IndexBulk 批量索引文档，支持 upsert（更新或插入）
func (es *ElasticsearchClient) IndexBulk(index string, docs map[string]string) (*esapi.Response, error) {
	var bulkRequest bytes.Buffer
	for docId, jsonMessage := range docs {
		meta := fmt.Sprintf(`{ "update" : { "_id" : "%s" } }%s`, docId, "\n")
		bulkRequest.WriteString(meta)

		var doc map[string]interface{}
		if err := json.Unmarshal([]byte(jsonMessage), &doc); err != nil {
			return nil, fmt.Errorf("unmarshal JSON failed: %w", err)
		}

		updateDoc := map[string]interface{}{
			"doc":           doc,
			"doc_as_upsert": true,
		}
		updateDocJSON, err := json.Marshal(updateDoc)
		if err != nil {
			return nil, fmt.Errorf("marshal update doc failed: %w", err)
		}

		bulkRequest.Write(updateDocJSON)
		bulkRequest.WriteByte('\n')
	}

	req := esapi.BulkRequest{
		Index:   index,
		Body:    bytes.NewReader(bulkRequest.Bytes()),
		Refresh: "true",
	}
	resp, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return nil, fmt.Errorf("bulk index failed: %w", err)
	}
	return resp, nil
}

// Update 更新指定索引中的文档
func (es *ElasticsearchClient) Update(index, docId, jsonMessage string) (*esapi.Response, error) {
	data := []byte(jsonMessage)
	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: docId,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}
	resp, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return nil, fmt.Errorf("update document failed: %w", err)
	}
	return resp, nil
}

// CreateIndex 创建带映射的索引（如已存在则跳过）
func (es *ElasticsearchClient) CreateIndex(index, mapping string) error {
	exists, err := es.Client.Indices.Exists([]string{index})
	if err != nil {
		return fmt.Errorf("check index exists failed: %w", err)
	}
	if exists.StatusCode == 200 {
		log.Info("索引已存在，跳过创建", zap.String("index", index))
		return nil
	}

	req := esapi.IndicesCreateRequest{
		Index: index,
		Body:  bytes.NewReader([]byte(mapping)),
	}
	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return fmt.Errorf("create index failed: %w", err)
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Error("关闭响应体失败", zap.Error(err))
		}
	}(res.Body)

	if res.IsError() {
		return fmt.Errorf("create index error response: %s", res.String())
	}

	log.Info("索引创建成功", zap.String("index", index))
	return nil
}
