package utils

import (
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	"math"
	"strings"
)

const (
	DefaultPageSize = 10
	DefaultMaxSize  = 100
	DefaultPage     = 0
	DefaultSort     = "updated_at"
	DefaultOrder    = "desc"
)

// NormalizePageable is a function that normalizes a pageable object.
func NormalizePageable(pageable *common.Pageable) *common.Pageable {
	if pageable == nil {
		return &common.Pageable{
			Page:      DefaultPage,
			Size:      DefaultPageSize,
			Sort:      DefaultSort,
			Direction: DefaultOrder,
		}
	}
	return &common.Pageable{
		Page:      AsPage(pageable.Page),
		Size:      AsPageSize(pageable.Size),
		Sort:      AsSort(pageable.Sort),
		Direction: AsOrder(pageable.Direction),
	}
}

// AsPageMetaData is a function that maps a pageable object to a page metadata object.
func AsPageMetaData(pageable *common.Pageable, total int64) *common.PageMetaData {
	resp := &common.PageMetaData{
		TotalElements: total,
	}

	if pageable.UnPaged {
		resp.TotalPages = 1
		resp.Page = 1
		resp.Size = int32(total)
		resp.Previous = false
		resp.Next = false
	} else {
		totalPages := AsTotalPages(int32(total), pageable.Size)
		resp.TotalPages = int64(totalPages)
		resp.Page = pageable.Page + 1
		resp.Size = pageable.Size
		resp.Previous = pageable.Page > 0
		resp.Next = pageable.Page < totalPages-1
	}

	return resp
}

func AsPage(page int32) int32 {
	p := int32(DefaultPage)
	if page > 0 {
		p = page - 1
	}
	return p
}

func AsPageSize(pageSize int32) int32 {
	ps := int32(DefaultPageSize)
	if pageSize > 0 {
		ps = int32(math.Min(float64(pageSize), float64(DefaultMaxSize)))
	}
	return ps
}

func AsSort(sort string) string {
	if sort == "" {
		return DefaultSort
	}
	return sort
}

func AsOrder(order string) string {
	if strings.EqualFold(order, "asc") {
		return "asc"
	}
	return DefaultOrder
}

func AsOffset(page int32, size int32) int32 {
	return page * size
}

func AsTotalPages(total int32, size int32) int32 {
	totalPages := total / size
	if total%size > 0 {
		totalPages++
	}
	return totalPages
}
