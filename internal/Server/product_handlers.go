package server

import (
	"strconv"

	"github.com/alphacoder5911/EcommerceBackend/internal/dto"
	"github.com/alphacoder5911/EcommerceBackend/internal/utills"
	"github.com/gin-gonic/gin"	
)

func (s *Server) CreateCategory(c *gin.Context){
	var req dto.CreateCategoryRequest
	if err:= c.ShouldBindJSON(&req);err!=nil{
		utills.BadRequestResponse(c,"Invalid json",err)
		return 
	}

	category,err:=s.productService.CreateCategory(&req)
	if err!=nil{
		utills.InternalServerErrorResponse(c,"Failed to create category",err)
		return 
	}

	utills.CreatedResponse(c,"Category created successfully",category)
}

func (s *Server) GetCategories(c *gin.Context){
	cats,err:=s.productService.GetCategories()
	if err!=nil{
		utills.InternalServerErrorResponse(c,"Couldnt fetch categories ",err)
	}

	utills.SuccessResponse(c,"Fetched categorries successfully",cats)
}

func (s *Server) UpdateCategory(c *gin.Context){
	id,err:=strconv.ParseUint(c.Param("id"),10,32)
	if err!=nil{
		utills.BadRequestResponse(c,"Invalid category id ",err)
		return 
	}
	var req dto.UpdateCategoryRequest
	if err:= c.ShouldBindJSON(&req);err!=nil{
		utills.BadRequestResponse(c,"Invalid json",err)
		return 
	}

	updated,err:=s.productService.UpdateCategory(c.GetUint(id),&req)
	if err!=nil{
		utills.InternalServerErrorResponse(c,"Failed to update category",err)
		return 
	}

	utills.SuccessResponse(c,"Updated Category successfully",updated)
}

func (s *Server) DeleteCategory(c *gin.Context){
	id ,err:= strconv.ParseUint(c.Param("Id"),10,32)
	if err!=nil{
		utills.BadRequestResponse(c,"Invalid category request",err)
		return
	}
	err=s.productService.DeleteCategory(c.GetUint(id))
	if err!=nil{
		utills.BadRequestResponse(c,"Failed to delete category ",err)
		return 
	}

	utills.SuccessResponse(c,"Deleted category successfully",nil)
}

func (s *Server) CreateProduct(c *gin.Context) {
	var product dto.CreateProductRequest
	if err:=c.ShouldBindJSON(&product);err!=nil{
		utills.BadRequestResponse(c,"Invalid json",err)
		return
	}

	prdt,err:=s.productService.CreateProduct(&product)
	if err!=nil{
		utills.BadRequestResponse(c,"Failed to create product",err)
		return
	}

	utills.CreatedResponse(c,"Product created ",prdt)
}




func (s *Server) GeetCategories(c *gin.Context){
	
	categories,err:=s.productService.GetCategories()
	if err!=nil{
		utills.BadRequestResponse(c,"Failed to get categories ",err)
		return 
	}

	utills.SuccessResponse(c,"Categories fetched successfully",categories)
}

func (s *Server) getProducts(c *gin.Context){
	page,_:=strconv.Atoi(c.DefaultQuery("page","1"))
	limit,_:=strconv.Atoi(c.DefaultQuery("limit","10"))

	productsn,meta,err:=s.productService.GetProducts(page,limit)
	if err!=nil{
		utills.BadRequestResponse(c,"Failed to get products ",err)
		return
	}

	utills.PaginatedSuccessResponse(c,"Products fetched sucessfully",productsn,*meta)
}

func (s *Server) GetProduct(c *gin.Context){
	id,err:=strconv.ParseUint(c.Param("id"),10,32)
	
	if err!=nil{
		utills.BadRequestResponse(c,"Invalid id presented",err)
		return 
	}

	product,err:=s.productService.GetProduct(uint(id))
	if err!=nil{
		utills.NotFoundResponse(c,"product not found ")
		return 
	}

	utills.SuccessResponse(c,"Fetched the product successfully",product)

}

func(s *Server) UpdateProduct(c *gin.Context){
	id,err:=strconv.ParseUint(c.Param("id"),10,32)
	if err!=nil{
		utills.BadRequestResponse(c,"Wrong id provided ",err)
		return
	}

	var update dto.UpdateProductRequest
	if err:=c.ShouldBindJSON(&update);err!= nil {
		utills.BadRequestResponse(c,"Invalid json",err)
		return
	}

	updatedProd,err:=s.productService.UpdateProduct(uint(id),&update)
	if err!=nil{
		utills.BadRequestResponse(c,"Failed to update ",err)
		return
	}

	utills.SuccessResponse(c,"Updated ",updatedProd)
}

func (s *Server) DeleteProduct(c *gin.Context){
	id,err:=strconv.ParseUint(c.Param("id"),10,32)
	if err!=nil{
		utills.BadRequestResponse(c,"Invalid id ",err)
		return
	}

	err=s.productService.DeleteCategory(uint(id))
	utills.SuccessResponse(c,"Deleted Successfully",nil)
}
