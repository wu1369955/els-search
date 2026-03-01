package service

import (
	"context"
	"backend/repository"
)

// SearchService 定义了搜索服务接口
type SearchService interface {
	Search(ctx context.Context, index, query string, fields []string, page, size int) (*repository.SearchResult, error)
}

// searchService 实现了 SearchService 接口
type searchService struct {
	esRepo repository.ESRepository
}

// NewSearchService 创建 SearchService 实例
func NewSearchService(esRepo repository.ESRepository) SearchService {
	return &searchService{esRepo: esRepo}
}

// Search 执行搜索操作
func (s *searchService) Search(ctx context.Context, index, query string, fields []string, page, size int) (*repository.SearchResult, error) {
	// 构建 Elasticsearch 查询
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": fields,
			},
		},
		"from": (page - 1) * size,
		"size": size,
	}

	// 执行搜索
	return s.esRepo.Search(ctx, index, esQuery)
}
