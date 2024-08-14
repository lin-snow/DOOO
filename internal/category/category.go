package category

import (
	"github.com/gin-gonic/gin"
	"github.com/lin-snow/dooo/internal/model"
	"github.com/lin-snow/dooo/pkg/auth"
	"gorm.io/gorm"
)

// Create A Category
func AddCategory(ctx *gin.Context, db *gorm.DB) {
	// Get Current User
	user := auth.GetCurrentUser(ctx, db)

	// Get Category Name
	newCategoryName := ctx.Query("newCategoryName")

	// Check if the category name is empty
	if newCategoryName == "" {
		ctx.JSON(200, gin.H{
			"Message": "Category Name is Blank",
			"Code":    model.ERR_CATEGORYNAME_BLANK,
			"Data":    gin.H{},
		})
		return
	}

	// Check if the category name is existed
	var tempCategory model.Category
	db.Where("name = ? AND user_id = ?", newCategoryName, user.ID).First(&tempCategory)
	if tempCategory.ID != 0 {
		ctx.JSON(200, gin.H{
			"Message": "Category Name is Existed!",
			"Code":    model.ERR_CATEGORY_EXIST,
			"Data":    tempCategory.Name,
		})
		return
	}

	// Create A New Category
	newCategory := model.Category{
		Name:   newCategoryName,
		UserID: user.ID,
	}
	db.Create(&newCategory)
	ctx.JSON(200, gin.H{
		"Message": "Create Category Successfully",
		"Code":    model.SUCCESS,
		"Data":    newCategory,
	})
}

// Query All Categories
func QueryCategories(ctx *gin.Context, db *gorm.DB) {
	// Get Current User
	user := auth.GetCurrentUser(ctx, db)

	// Query All Categories
	var categories []model.Category
	db.Where("user_id = ?", user.ID).Find(&categories)

	// Check if the categories are empty
	if len(categories) == 0 {
		ctx.JSON(200, gin.H{
			"Message": "Categories Not Found",
			"Code":    model.ERR_CATEGORY_NOT_FOUND,
			"Data":    gin.H{},
		})
		return
	}

	// Return Categories
	ctx.JSON(200, gin.H{
		"Message": "Query Categories Successfully",
		"Code":    model.SUCCESS,
		"Data":    categories,
	})
}

// Update Category
func UpdateCategory(ctx *gin.Context, db *gorm.DB) {
	// Get Current User
	user := auth.GetCurrentUser(ctx, db)

	// Get Category ID
	categoryID := ctx.Query("categoryID")

	// Get New Category Name
	newCategoryName := ctx.Query("newCategoryName")

	// Check if the category name is empty
	if newCategoryName == "" {
		ctx.JSON(200, gin.H{
			"Message": "Category Name is Blank",
			"Code":    model.ERR_CATEGORYNAME_BLANK,
			"Data":    gin.H{},
		})
		return
	}

	// Check if the category name is existed
	var tempCategory model.Category
	db.Where("name = ? AND user_id = ?", newCategoryName, user.ID).First(&tempCategory)
	if tempCategory.ID != 0 {
		ctx.JSON(200, gin.H{
			"Message": "Category Name is Existed!",
			"Code":    model.ERR_CATEGORY_EXIST,
			"Data":    tempCategory.Name,
		})
		return
	}

	// Update Category
	var category model.Category
	db.Where("id = ? AND user_id = ?", categoryID, user.ID).First(&category)
	if category.ID == 0 {
		ctx.JSON(200, gin.H{
			"Message": "Category Not Found",
			"Code":    model.ERR_CATEGORY_NOT_FOUND,
			"Data":    gin.H{},
		})
		return
	} else {
		category.Name = newCategoryName
		db.Save(&category)
		ctx.JSON(200, gin.H{
			"Message": "Update Category Successfully",
			"Code":    model.SUCCESS,
			"Data":    category,
		})
	}
}

// Delete Category
func DeleteCategory(ctx *gin.Context, db *gorm.DB) {
	// Get Current User
	user := auth.GetCurrentUser(ctx, db)

	// Get Category ID
	categoryID := ctx.Query("categoryID")

	// Delete All Todos in the Category
	var delallincate bool
	delallincateStr := ctx.Query("delallincate")
	if delallincateStr == "true" {
		delallincate = true
	} else {
		delallincate = false
	}

	// Delete Category
	var category model.Category
	db.Where("id = ? AND user_id = ?", categoryID, user.ID).First(&category)
	if category.ID == 0 {
		ctx.JSON(200, gin.H{
			"Message": "Category Not Found",
			"Code":    model.ERR_CATEGORY_NOT_FOUND,
			"Data":    gin.H{},
		})
		return
	} else {
		// Delete All Todos in the Category
		if delallincate {
			var todos []model.Todo
			db.Where("category_id = ?", category.ID).Find(&todos)
			for _, todo := range todos {
				db.Delete(&todo)
			}
		}

		// Delete Category
		db.Delete(&category)
		ctx.JSON(200, gin.H{
			"Message": "Delete Category Successfully",
			"Code":    model.SUCCESS,
			"Data":    gin.H{},
		})
	}
}
