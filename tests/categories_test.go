package tests

import (
	"testing"

	"github.com/damifur/catalog/dao"
	"github.com/damifur/catalog/services"
)

func TestCanCreateEditAndDeleteCategory(t *testing.T) {
	dao.CreateDatabase("test")
	defer dao.CloseDB()

	db, err := dao.GetSession()
	if err != nil {
		t.Fatal("Failed to connect to database")
	}

	name := "TestCategory"
	category, err1 := services.CreateCategory(db, name)

	if err1 != nil || category.Name != name {
		t.Fatalf("Category name was not saved properly: %v", err1)
	}

	category, err1 = services.GetCategory(db, category.Id)
	if err1 != nil {
		t.Fatalf("There was an error retrieving category with id %v: %v", category.Id, err1)
	}

	if category.Name != name {
		t.Fatal("Category name is not correct")
	}

	newName := "TestCategoryNew"
	category, err1 = services.EditCategory(db, category, newName)
	if err1 != nil {
		t.Fatalf("Category name was not properly changed: %v", err1)
	}

	category, err1 = services.GetCategory(db, category.Id)
	if err1 != nil {
		t.Fatalf("There was an error retrieving category with id %v: %v", category.Id, err1)
	}

	if category.Name != newName {
		t.Fatal("Category name is not correct")
	}

	if err1 = services.DeleteCategory(db, category.Id); err1 != nil {
		t.Fatalf("Category was not correctly deleted: %v", err1)
	}

	_, err1 = services.GetCategory(db, category.Id)
	if err1 == nil || err1.Status != 404 {
		t.Fatalf("Should have gotten a 404 but instead got %v", err1)
	}

}
