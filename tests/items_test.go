package tests

import (
	"testing"

	"github.com/damifur/catalog/dao"
	"github.com/damifur/catalog/services"
)

func TestCanCreateEditAndDeleteItem(t *testing.T) {
	dao.CreateDatabase("test")
	defer dao.CloseDB()

	db, err := dao.GetSession()
	if err != nil {
		t.Fatal("Failed to connect to database")
	}

	category, err1 := services.CreateCategory(db, "TestCategory")
	if err1 != nil {
		t.Fatalf("Error while creating category: %v", err1)
	}

	name, image, description, price := "item1", "imageencodedexampletest", "This is a test Item", 6.4
	item, err1 := services.CreateItem(db, name, image, description, category.Id, price)
	if err1 != nil {
		t.Fatalf("Error creating item: %v", err)
	}

	if item.Name != name || item.Description != description || item.Image != image || item.Price != price {
		t.Fatalf("Item returned by create function is not well formed: %v", item)
	}

	item, err1 = services.GetItem(db, item.Id)
	if err1 != nil {
		t.Fatalf("Error while retrieving item: %v", err)
	}

	if item.Name != name || item.Description != description || item.Image != image || item.Price != price {
		t.Fatalf("Item returned by create function is not well formed: %v", item)
	}

	newName, newImage, newDescription, newPrice := "item1 CHANGED", "imageencodedexampletestCHANGED", "This is a changed test Item", 8.4
	item.Name, item.Image, item.Description, item.Price = newName, newImage, newDescription, newPrice

	item, err1 = services.UpdateItem(db, item.Id, item)
	if err1 != nil {
		t.Fatalf("Error while updating item with id %d: %v", item.Id, err)
	}

	item, err1 = services.GetItem(db, item.Id)
	if err1 != nil {
		t.Fatalf("Error while retrieving item: %v", err)
	}

	item, err1 = services.GetItem(db, item.Id)
	if err1 != nil {
		t.Fatalf("Error while retrieving item: %v", err)
	}

	if item.Name != newName || item.Description != newDescription || item.Image != newImage || item.Price != newPrice {
		t.Fatalf("Item returned by create function is not well formed: %v", item)
	}

	if err := services.DeleteItem(db, item.Id); err != nil {
		t.Fatalf("There was an error while deleting item with id %d: %v", item.Id, err)
	}

	_, err1 = services.GetItem(db, item.Id)
	if err1 == nil || err1.Status != 404 {
		t.Fatalf("Should have gotten a 404 but instead got %v", err1)
	}

}

func TestCantCreateItemIfCategoryDoesntExist(t *testing.T) {
	dao.CreateDatabase("test")
	defer dao.CloseDB()

	db, err := dao.GetSession()
	if err != nil {
		t.Fatal("Failed to connect to database")
	}

	name, image, description, price := "item1", "imageencodedexampletest", "This is a test Item", 6.4
	_, err1 := services.CreateItem(db, name, image, description, 12345, price)
	if err1 == nil || err1.Status != 404 {
		t.Fatalf("We should have got a 500 error but instead we got: %v", err1)
	}

}
