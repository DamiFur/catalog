package services

import (
	"fmt"

	"github.com/damifur/catalog/dao"
	"github.com/damifur/catalog/utils/errors"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB, name, image, description string, category int, price float64) (dao.Item, *errors.Error) {
	cat, err := GetCategory(db, category)
	if err != nil {
		return dao.Item{}, err
	}
	item := dao.Item{Name: name, Image: image, Description: description, CategoryID: cat.ID, Category: cat, Price: price}

	if err := db.Create(&item).Error; err != nil {
		return dao.Item{}, errors.InternalServer("Error while creating item", err)
	}

	return item, nil
}

func GetItem(db *gorm.DB, id int) (dao.Item, *errors.Error) {
	item := dao.Item{}
	if err := db.First(&item, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.Item{}, errors.NotFound(fmt.Sprintf("Category with id %v was not found", id))
		}
		return dao.Item{}, errors.InternalServer(fmt.Sprintf("There was an error while retrieving category with id %v from DB", id), err)
	}
	return item, nil
}

func UpdateItem(db *gorm.DB, id int, newItem dao.Item) (dao.Item, *errors.Error) {
	newItem.Id = id
	if err := db.Save(newItem).Error; err != nil {
		return dao.Item{}, errors.InternalServer(fmt.Sprintf("There was an error updating item with id %d", id), err)
	}

	return newItem, nil
}

func DeleteItem(db *gorm.DB, id int) *errors.Error {
	item := dao.Item{Id: id}
	if err := db.Delete(&item).Error; err != nil {
		return errors.InternalServer(fmt.Sprintf("Error while deleting item with id %v", id), err)
	}
	return nil
}

func GetAllItems(db *gorm.DB) ([]dao.Item, *errors.Error) {
	ans := make([]dao.Item, 0)

	if err := db.Preload("Category").Find(&ans).Error; err != nil {
		return nil, errors.InternalServer("Error while retrieving items", err)
	}

	return ans, nil
}
