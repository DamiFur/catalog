package controllers

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/damifur/catalog/dao"
	"github.com/damifur/catalog/services"
	"github.com/damifur/catalog/utils/errors"
	"github.com/damifur/catalog/utils/validators"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoryInput struct {
	Name string
}

type ItemInput struct {
	Name        string
	Price       float64
	Category    int
	Description string
}

type CategoriesOut struct {
	Categories []dao.Category
	Items      []dao.Item
}

func GetAdminPanel(c *gin.Context) {

	db, err := GetDB(c)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	categories, err := services.GetAllCategories(db)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	items, err := services.GetAllItems(db)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	for _, item := range items {
		fmt.Println("Category ", item.CategoryID, item.Category)
	}

	c.HTML(http.StatusOK, "admin-panel.html", CategoriesOut{categories, items})
	return
}

func CreateCategory(c *gin.Context) {

	// TODO: check admin credentials

	v := validators.New()

	input := CategoryInput{}
	v.SafeBody(c, &input)
	if err := v.Error(); err != nil {
		errors.Respond(c, err)
		return
	}

	db, err := GetDB(c)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	cat, err := services.CreateCategory(db, input.Name)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, cat)
	return
}

func CreateItem(c *gin.Context) {

	req := c.Request

	fmt.Println("DEBUG 1")

	// Max File size: 10MB
	if err0 := req.ParseMultipartForm(10 << 20); err0 != nil {
		fmt.Println(err0)
		errors.Respond(c, errors.InternalServer("There was an error uploading file", err0))
		return
	}


	fmt.Println("DEBUG 2")

	file, handler, err0 := req.FormFile("image")
	if err0 != nil {
		fmt.Println(err0)
		errors.Respond(c, errors.InternalServer("There was an error uploading file", err0))
		return
	}

	fmt.Println("DEBUG 3")

	defer file.Close()
	fmt.Println("Uploaded File: %+v\n", handler.Filename)
	fmt.Println("File Size: %+v\n", handler.Size)
	fmt.Println("MIME Header: %+v\n", handler.Header)

	buf := make([]byte, handler.Size)
	imageReader := bufio.NewReader(file)
	imageReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	//fmt.Println(imgBase64Str)
	name := req.FormValue("name")
	desc := req.FormValue("description")
	price, err0 := strconv.ParseFloat(req.FormValue("price"), 64)
	if err0 != nil {
		errors.Respond(c, errors.BadRequest("Price must be a number"))
		return
	}
	catId, err0 := strconv.Atoi(req.FormValue("category"))
	if err0 != nil {
		errors.Respond(c, errors.BadRequest("There is a problem with category ID"))
		return
	}
	// TODO: check admin credentials

	db, err := GetDB(c)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	item, err := services.CreateItem(db, name, imgBase64Str, desc, catId, price)
	if err != nil {
		errors.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, item)
	return
}

func UploadImage(c *gin.Context) {
	req := c.Request

	// Max File size: 20MB
	req.ParseMultipartForm(10 << 20)

	file, handler, err := req.FormFile("tmp-img")
	if err != nil {
		errors.Respond(c, errors.InternalServer("There was an error uploading file", err))
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	buf := make([]byte, handler.Size)
	imageReader := bufio.NewReader(file)
	imageReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	fmt.Println(imgBase64Str)

	//tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer tempFile.Close()

	//fileBytes, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//tempFile.Write(fileBytes)

}
