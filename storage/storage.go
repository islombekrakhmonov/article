package storage

import "bootcamp/article/models"

type StorageI interface {
	Article() ArticleRepoI
	Author() AuthorRepoI
}

type ArticleRepoI interface {
	Create(entity models.ArticleCreateModel) (err error)
	GetList(query models.Query) (resp []models.ArticleListItem, err error)
	GetByID(ID string) (resp models.Article, err error)
	Update(entity models.ArticleUpdateModel) (int64, error)
	Delete(ID string) (effectedRowsNum int, err error)
}

type AuthorRepoI interface {
	Create(entity models.PersonCreateModel) (err error)
	GetList(query models.Query) (resp []models.Person, err error)
	GetByID(ID string) (resp models.Person, err error)
	Update(entity models.PersonUpdateModel) (error)
	Delete(ID string) (effectedRowsNum int, err error)
}
