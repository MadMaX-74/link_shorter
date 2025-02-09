package link

import (
	"go_dev/pkg/db"
)

type LinkRepository struct {
	Databse *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Databse: database,
	}
}

func (repository *LinkRepository) Create(link *Link) (*Link, error) {
	res := repository.Databse.Create(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}
func (repository *LinkRepository) Read(id uint) (*Link, error) {
	var link Link
	res := repository.Databse.First(&link, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &link, nil
}
func (repository *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	res := repository.Databse.First(&link, "hash = ?", hash)
	if res.Error != nil {
		return nil, res.Error
	}
	return &link, nil
}
func (repository *LinkRepository) Delete(id uint) error {
	res := repository.Databse.Delete(&Link{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (repository *LinkRepository) Update(link *Link) (*Link, error) {
	res := repository.Databse.Updates(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}
func (repository *LinkRepository) Count() int64 {
	var count int64
	repository.Databse.
		Table("links").
		Where("deleted_at IS NULL").
		Count(&count)
	return count
}

func (repository *LinkRepository) GetAll(limit, offset int) []*Link {
	var links []*Link
	repository.Databse.
		Table("links").
		Where("deleted_at IS NULL").
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Scan(&links)
	return links
}
