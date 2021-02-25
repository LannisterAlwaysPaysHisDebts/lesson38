package handler

import (
	"context"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/04category/common"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/04category/domain/model"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/04category/domain/service"

	category "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/04category/proto/category"
)

type Category struct {
	Srv service.ICategoryDataService
}

func (c *Category) CreateCategory(ctx context.Context, in *category.CategoryRequest, out *category.CreateCategoryResponse) error {
	categoryModel := &model.Category{}
	err := common.SwapTo(in, categoryModel)
	if err != nil {
		return err
	}
	categoryId, err := c.Srv.AddCategory(categoryModel)
	if err != nil {
		return err
	}
	out.Message = "添加成功"
	out.CategoryId = categoryId
	return nil
}

func (c *Category) UpdateCategory(ctx context.Context, in *category.CategoryRequest, out *category.UpdateCategoryResponse) error {
	category2 := &model.Category{}
	err := common.SwapTo(in, category2)
	if err != nil {
		return err
	}
	err = c.Srv.UpdateCategory(category2)
	if err != nil {
		return err
	}
	out.Message = "分类更新成功"
	return nil
}

func (c *Category) DeleteCategory(ctx context.Context, in *category.DeleteCategoryRequest, out *category.DeleteCategoryResponse) error {
	err := c.Srv.DeleteCategory(in.CategoryId)
	if err != nil {
		return nil
	}
	out.Message = "删除成功"
	return nil
}

func (c *Category) FindCategoryByName(ctx context.Context, in *category.FindByNameRequest, out *category.CategoryResponse) error {
	category2, err := c.Srv.FindCategoryByName(in.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category2, out)
}

func (c *Category) FindCategoryById(ctx context.Context, in *category.FindByIdRequest, out *category.CategoryResponse) error {
	category2, err := c.Srv.FindCategoryByID(in.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category2, out)
}

func (c *Category) FindCategoryByLevel(ctx context.Context, in *category.FindByLevelRequest, out *category.FindAllResponse) error {
	categorySlice, err := c.Srv.FindCategoryByLevel(in.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, out)
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context, in *category.FindByParentRequest, out *category.FindAllResponse) error {
	categorySlice, err := c.Srv.FindCategoryByParent(in.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, out)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, in *category.FindAllRequest, out *category.FindAllResponse) error {
	categorySlice, err := c.Srv.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, out)
	return nil
}

func categoryToResponse(categorySlice []model.Category, out *category.FindAllResponse) {
	for _, cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		out.Category = append(out.Category, cr)
	}
}
