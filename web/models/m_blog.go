package models

import "github.com/jinzhu/gorm"

type (
	Blog struct {
		gorm.Model
		Title	string
		Content	string
	}
	BlogFilter struct {
		Paging
		Ids []int
		IdGreater int
	}
	Paging struct {

	}
)

func (b *Blog) Create() error {
	return db.Save(b).Error
}

func (b *Blog) GetById() error {
	return db.First(b, b.ID).Error
}

func (b *Blog) Update() error {
	return db.Model(&Blog{}).Update(b).Error
}

func (f *BlogFilter) GetList() (blogs []Blog, err error) {
	err = db.Scopes(f.scope()...).Find(&blogs).Error
	return
}

func (b *Blog) Delete() error {
	return db.Delete(b).Error
}

func (f *BlogFilter) scope() (funcs []func(db2 *gorm.DB) *gorm.DB) {
	if len(f.Ids) > 0 {
		funcs = append(funcs, func(db2 *gorm.DB) *gorm.DB {
			return db2.Where("id in (?)", f.Ids)
		})
	}
	return
}