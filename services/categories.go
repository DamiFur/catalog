package services

import (
	"fmt"

	"github.com/damifur/catalog/dao"
	"github.com/damifur/catalog/utils/errors"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, name string) (dao.Category, *errors.Error) {
	category := dao.Category{Name: name}

	if err := db.Create(&category).Error; err != nil {
		return dao.Category{}, errors.InternalServer("Error while creating category", err)
	}

	return category, nil
}

func EditCategory(db *gorm.DB, category dao.Category, newName string) (dao.Category, *errors.Error) {
	err := db.Model(&category).Update("name", newName).Error
	if err != nil {
		return dao.Category{}, errors.InternalServer("Error while editing category name", err)
	}

	category.Name = newName

	return category, nil
}

func DeleteCategory(db *gorm.DB, categoryId int) *errors.Error {
	category := dao.Category{ID: categoryId}
	if err := db.Delete(&category).Error; err != nil {
		return errors.InternalServer("There was an error deleting category: %v", err)
	}
	return nil
}

func GetCategory(db *gorm.DB, id int) (dao.Category, *errors.Error) {
	category := dao.Category{}
	if err := db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.Category{}, errors.NotFound(fmt.Sprintf("Category with id %v was not found", id))
		}
		return dao.Category{}, errors.InternalServer(fmt.Sprintf("There was an error while retrieving category with id %v from DB", id), err)
	}
	return category, nil
}

func GetAllCategories(db *gorm.DB) ([]dao.Category, *errors.Error) {
	ans := make([]dao.Category, 0)

	if err := db.Find(&ans).Error; err != nil {
		return nil, errors.InternalServer("Error while retrieving Categories", err)
	}

	return ans, nil
}
