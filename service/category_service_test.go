package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	categorySuccess := entity.Category{
		Id:   "2",
		Name: "Djalal",
	}

	categoryRepository.Mock.On("FindById", "2").Return(categorySuccess)
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, categorySuccess.Id, result.Id)
	assert.Equal(t, categorySuccess.Name, result.Name)
}
