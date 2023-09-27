package services

import (
	"context"
	"time"

	"github.com/sathwikm17/Ecom_Project/entities"
	"github.com/sathwikm17/Ecom_Project/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection
}

func InitProductService(collection *mongo.Collection) interfaces.IProduct {
	return &ProductService{Product: collection}
}

func (prod *ProductService) AddProduct(p *entities.Product) (string, error) {
	p.Id = primitive.NewObjectID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = p.CreatedAt
	_, err := prod.Product.InsertOne(context.Background(), p)
	if err != nil {
		return "", err
	} else {
		return "Record Inserted Successfully", nil
	}
}

func (prod *ProductService) GetProductById(id primitive.ObjectID) (*entities.Product, error) {

	ctx := context.Background()
	var product entities.Product
	err := prod.Product.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (prod *ProductService) SearchProducts(name string) (*[]entities.Product, error) {
	return nil, nil
}
