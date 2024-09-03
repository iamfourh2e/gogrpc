package repo

import (
	"context"
	"fmt"
	"seyhadot/gogrpc/pb"
)

type ProductRepo struct {
	pb.UnimplementedProductServiceServer
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

func (p *ProductRepo) CreateProduct(c context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := &pb.CreateProductRequest{
		Product: &pb.Product{
			Id:          req.Product.Id,
			Name:        req.Product.Name,
			Description: req.Product.Description,
			Price:       req.Product.Price,
		},
	}
	fmt.Printf("Product created: %v\n", product)
	return nil, nil
}
func (p *ProductRepo) GetProduct(c context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	return nil, nil
}
func (p *ProductRepo) UpdateProduct(c context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	return nil, nil
}
func (p *ProductRepo) DeleteProduct(c context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	return nil, nil
}
