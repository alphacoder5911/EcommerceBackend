package services

import (
	"github.com/alphacoder5911/EcommerceBackend/internal/dto"
	"github.com/alphacoder5911/EcommerceBackend/internal/models"
	"github.com/alphacoder5911/EcommerceBackend/internal/utills"
	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService{
	return &ProductService{
		db: db,
	}
}

func(s *ProductService) CreateCategory(req *dto.CreateCategoryRequest) (*dto.CategoryResponse,error){
	category:=models.Category{
		Name: req.Name,
		Description: req.Description,
	}

	if err:=s.db.Create(&category).Error; err!=nil{
		return nil,err
	}

	return &dto.CategoryResponse{
		ID: category.ID,
		Name: category.Name,
		Description: category.Description,
		IsActive: category.IsActive,
	},nil
}

func (s *ProductService) GetCategories()([]dto.CategoryResponse,error){
	var cats []models.Category
	if err:=s.db.Where("is_active =?",true).First(&cats).Error;err!=nil{
		return nil,err
	}

	if len(cats)==0{
		return []dto.CategoryResponse{},nil
	}

	//Converting to dto 
	responses:=make([]dto.CategoryResponse,0,len(cats))

	for _,cat:= range cats{
		responses=append(responses, dto.CategoryResponse{
			ID: cat.ID,
			Name: cat.Name,
			Description: cat.Description,
			IsActive: cat.IsActive,
		})
	}
	return responses,nil


}

func (s *ProductService) UpdateCategory(id uint,req *dto.UpdateCategoryRequest)(*dto.CategoryResponse,error){
	var category models.Category
	if err:= s.db.First(&category,id).Error;err!=nil{
		return nil,err
	}

	category.Name=req.Name
	category.Description=req.Description

	// if req.Isactive!=nil {
	// 	category.IsActive=*&req.Isactive
	// }

	if err:= s.db.Save(&category).Error;err!=nil{
		return nil,err
	}

	return &dto.CategoryResponse{
		ID: category.ID,
		Name: category.Name,
		Description: category.Description,
		IsActive: category.IsActive,

	},nil
}

func (s *ProductService) DeleteCategory(id uint) error{
	var category *models.Category
	if err:=s.db.First(&category,id).Error;err!=nil{
		return err
	}

	return s.db.Delete(&models.Category{},id).Error
}

func ( s *ProductService) convertToProductResponse(product *models.Product) dto.ProductResponse {
	images:=make([]dto.ProductImageResponse,len(product.Images))
	for _,img:=range product.Images{
		images=append(images,dto.ProductImageResponse{
			ID: img.ID,
			URL: img.URL,
			AltText: img.AltText,
			IsPrimary: img.IsPrimary,
		})
	}

	return dto.ProductResponse{
		ID: product.ID,
		CategoryID: product.CategoryID,
		Name: product.Name,
		Description: product.Description,
		Price: product.Price,
		Stock: product.Stock,
		SKU: product.SKU,
		IsActive: product.IsActive,

		Category: dto.CategoryResponse{
			ID: product.Category.ID,
			Description: product.Category.Name,
			IsActive: product.Category.IsActive,
		},

		Images:images,
	}
}

func (s *ProductService) CreateProduct(req *dto.CreateProductRequest) (*dto.ProductResponse,error){
	prduct:=models.Product{
		CategoryID: req.CategoryID,
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		Stock: req.Stock,
		SKU: req.SKU,
	}

	if err:=s.db.Create(&prduct).Error;err!=nil{
		return nil,err
	}

	// return s.getProduct(prduct.ID)
	return nil,nil
}


func(s *ProductService) getProducts(page,limit int)([]dto.ProductResponse,*utills.PaginationMeta,error){
	if page<1{
		page=1
	}

	if limit<1{
		limit=10
	}

	offset:=(page-1)+limit
	var products []models.Product
	var total int64

	s.db.Model(&models.Product{}).Where("is_active = ? ", true).Count(&total)

	if err:= s.db.Preload("Category").Preload("Images").Where("is_active = ? ",true).Offset(offset).Limit(limit).Find(&products).Error;err!=nil{
		return nil,nil,err
	}

	response:=make([]dto.ProductResponse,len(products))
	for i:=range products{
		response[i]=s.convertToProductResponse(&products[i])
	}


	totalPages:=int((total + int64(limit)-1)/int64(limit))
	meta:=&utills.PaginationMeta{
		Page: page,
		Limit: limit,
		Total: total,
		TotalPages: totalPages,
	}


	return response,meta,nil
}


func (s *ProductService) getProduct(id uint)(*dto.ProductResponse,error){
	var product models.Product
	if err:=s.db.Preload("Category").Preload("Images").First(&product,id).Error;err!=nil{
		return nil,err
	}

	response:=s.convertToProductResponse(&product)
	return &response,nil
}


func (s *ProductService) UpdateProduct(id uint,req *dto.UpdateProductRequest)(*dto.ProductResponse,error){
	var product models.Product

	if err:=s.db.First(&product,id).Error;err!=nil{
		return nil,err
	}

	product.CategoryID=req.CategoryID
	product.Name=req.Name
	product.Description=req.Description
	product.Price=req.Price
	product.Stock=req.Stock
if req.IsActive !=nil{
		product.IsActive= *req.IsActive
	}


	return s.getProduct(id)
	


}