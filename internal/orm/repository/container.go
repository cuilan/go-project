package repository

import (
	"log/slog"
	"sync"
)

// repoContainer 是依赖注入容器
type repoContainer struct {
	repoMap map[string]interface{}
	mu      sync.RWMutex
}

// Register 注册一个repository到容器中
func (c *repoContainer) register(name string, repo interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.repoMap[name] = repo
	slog.Debug("repository registered", "name", name)
}

// Get 从容器中获取repository
func (c *repoContainer) get(name string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	repo, exists := c.repoMap[name]
	return repo, exists
}

// 全局容器实例
var globalContainer *repoContainer
var once sync.Once

// getRepoContainer 获取全局容器实例
func getRepoContainer() *repoContainer {
	once.Do(func() {
		globalContainer = &repoContainer{
			repoMap: make(map[string]interface{}),
		}
	})
	return globalContainer
}

// RegisterRepository 注册repository
func RegisterRepository[V string, T interface{}](name string, repo T) {
	getRepoContainer().register(name, repo)
}

// GetRepository 获取repository
func GetRepository(name string) (interface{}, bool) {
	return getRepoContainer().get(name)
}
