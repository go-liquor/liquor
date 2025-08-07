package redis

import (
	"context"
	"errors"
	"sync"
	"time"
)

type inMemoryItem struct {
	value      interface{}
	expiration *time.Time
}

type inMemoryRedis struct {
	data       map[string]inMemoryItem
	hashData   map[string]map[string]string
	listData   map[string][]string
	mu         sync.RWMutex
	cleanupInt time.Duration
}

// NewInMemoryRedis creates a new in-memory implementation of RedisClient
func NewInMemoryRedis() RedisClient {
	r := &inMemoryRedis{
		data:       make(map[string]inMemoryItem),
		hashData:   make(map[string]map[string]string),
		listData:   make(map[string][]string),
		cleanupInt: time.Minute,
	}
	go r.cleanup()
	return r
}

func (r *inMemoryRedis) cleanup() {
	ticker := time.NewTicker(r.cleanupInt)
	for range ticker.C {
		r.mu.Lock()
		now := time.Now()
		for k, v := range r.data {
			if v.expiration != nil && now.After(*v.expiration) {
				delete(r.data, k)
			}
		}
		r.mu.Unlock()
	}
}

func (r *inMemoryRedis) Set(_ context.Context, key string, value interface{}, expiration time.Duration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var exp *time.Time
	if expiration > 0 {
		t := time.Now().Add(expiration)
		exp = &t
	}

	r.data[key] = inMemoryItem{
		value:      value,
		expiration: exp,
	}
	return nil
}

func (r *inMemoryRedis) Get(_ context.Context, key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, exists := r.data[key]
	if !exists {
		return "", errors.New("key not found")
	}

	if item.expiration != nil && time.Now().After(*item.expiration) {
		delete(r.data, key)
		return "", errors.New("key expired")
	}

	return item.value.(string), nil
}

func (r *inMemoryRedis) Delete(_ context.Context, keys ...string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, key := range keys {
		delete(r.data, key)
	}
	return nil
}

func (r *inMemoryRedis) Exists(_ context.Context, keys ...string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, key := range keys {
		if _, exists := r.data[key]; !exists {
			return false, nil
		}
	}
	return true, nil
}

func (r *inMemoryRedis) Expire(_ context.Context, key string, expiration time.Duration) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, exists := r.data[key]
	if !exists {
		return false, nil
	}

	t := time.Now().Add(expiration)
	item.expiration = &t
	r.data[key] = item
	return true, nil
}

func (r *inMemoryRedis) Incr(_ context.Context, key string) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var value int64
	if item, exists := r.data[key]; exists {
		if v, ok := item.value.(int64); ok {
			value = v
		}
	}

	value++
	r.data[key] = inMemoryItem{
		value:      value,
		expiration: nil,
	}
	return value, nil
}

func (r *inMemoryRedis) HSet(_ context.Context, key string, values ...interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(values)%2 != 0 {
		return errors.New("invalid number of arguments")
	}

	if _, exists := r.hashData[key]; !exists {
		r.hashData[key] = make(map[string]string)
	}

	for i := 0; i < len(values); i += 2 {
		field := values[i].(string)
		value := values[i+1].(string)
		r.hashData[key][field] = value
	}
	return nil
}

func (r *inMemoryRedis) HGet(_ context.Context, key, field string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if hash, exists := r.hashData[key]; exists {
		if value, exists := hash[field]; exists {
			return value, nil
		}
	}
	return "", errors.New("field not found")
}

func (r *inMemoryRedis) HGetAll(_ context.Context, key string) (map[string]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if hash, exists := r.hashData[key]; exists {
		result := make(map[string]string)
		for k, v := range hash {
			result[k] = v
		}
		return result, nil
	}
	return nil, errors.New("key not found")
}

func (r *inMemoryRedis) HDel(_ context.Context, key string, fields ...string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if hash, exists := r.hashData[key]; exists {
		for _, field := range fields {
			delete(hash, field)
		}
	}
	return nil
}

func (r *inMemoryRedis) LPush(_ context.Context, key string, values ...interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.listData[key]; !exists {
		r.listData[key] = make([]string, 0)
	}

	for _, v := range values {
		r.listData[key] = append([]string{v.(string)}, r.listData[key]...)
	}
	return nil
}

func (r *inMemoryRedis) LPop(_ context.Context, key string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if list, exists := r.listData[key]; exists && len(list) > 0 {
		value := list[0]
		r.listData[key] = list[1:]
		return value, nil
	}
	return "", errors.New("list is empty")
}

func (r *inMemoryRedis) RPush(_ context.Context, key string, values ...interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.listData[key]; !exists {
		r.listData[key] = make([]string, 0)
	}

	for _, v := range values {
		r.listData[key] = append(r.listData[key], v.(string))
	}
	return nil
}

func (r *inMemoryRedis) RPop(_ context.Context, key string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if list, exists := r.listData[key]; exists && len(list) > 0 {
		lastIdx := len(list) - 1
		value := list[lastIdx]
		r.listData[key] = list[:lastIdx]
		return value, nil
	}
	return "", errors.New("list is empty")
}

func (r *inMemoryRedis) LRange(_ context.Context, key string, start int64, end int64) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list, exists := r.listData[key]
	if !exists {
		return []string{}, nil
	}

	listLen := int64(len(list))
	if listLen == 0 {
		return []string{}, nil
	}

	if start < 0 {
		start = listLen + start
	}
	if end < 0 {
		end = listLen + end
	}

	if start < 0 {
		start = 0
	}
	if end >= listLen {
		end = listLen - 1
	}

	if start > end {
		return []string{}, nil
	}

	result := make([]string, end-start+1)
	copy(result, list[start:end+1])
	return result, nil
}
