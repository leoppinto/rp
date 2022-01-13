package service

import (
	"context"
	"rp/grpc/pb"
	"rp/model"
)

type ProductGrpcServer struct {
	pb.UnimplementedProductServiceServer
	Producs *model.Products
}

func (p *ProductGrpcServer) CreateProduct(ctx context.Context, in *pb.Product) (*pb.ProductResult, error) {

	product := model.NewProduct()
	product.Name = in.Name
	product.ID = "XXXX-XXXX-XXXX-XXXX"
	p.Producs.Add(product)

	return &pb.ProductResult{
		Id:   product.ID,
		Name: product.Name,
	}, nil
}

func NewProductGrpcServer(producs *model.Products) *ProductGrpcServer {
	return &ProductGrpcServer{
		Producs: producs,
	}

}
