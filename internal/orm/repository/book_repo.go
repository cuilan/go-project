package repository

// BookRepositoryName 用户repository名称
const BookRepositoryName = "book_repository"

// GetBookRepository 获取用户repository（便捷方法）
func GetBookRepository() IBookRepository {
	repo, exists := GetRepository(BookRepositoryName)
	if !exists {
		panic("book repository not registered")
	}
	return repo.(IBookRepository)
}

type IBookRepository interface{}
