package interfaces

import (
	"github.com/sathwikm17/Ecom_Project/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProduct interface {
	AddProduct(p *entities.Product) (string, error)
	GetProductById(id primitive.ObjectID) (*entities.Product, error)
	SearchProducts(name string) (*[]entities.Product, error)
}
