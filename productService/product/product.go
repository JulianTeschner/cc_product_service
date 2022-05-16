package product

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID    primitive.ObjectID `bson:"_id, omitempty" json:"_id"`
	UUID  primitive.ObjectID `bson:"uuid, omitempty" json:"uuid"`
	Name  string             `bson:"name" json:"name"`
	Price float64            `bson:"price" json:"price"`
}

// UnmarshalJSON is a custom json unmarshaler for User
func (u *Product) UnmarshalJSON(data []byte) error {
	// Unmarshal into a temporary type where the "ends" field is a string.
	decoded := new(struct {
		ID    primitive.ObjectID `json:"_id"`
		UUID  primitive.ObjectID `json:"uuid"`
		Name  string             `json:"name"`
		Price float64            `json:"price"`
	})

	if err := json.Unmarshal(data, decoded); err != nil {
		return err
	}

	if decoded.ID == primitive.NilObjectID {
		u.ID = primitive.NewObjectID()
	} else {
		u.ID = decoded.ID
	}

	if decoded.UUID == primitive.NilObjectID {
		u.UUID = primitive.NewObjectID()
	} else {
		u.UUID = decoded.ID
	}

	u.Name = decoded.Name
	u.Price = decoded.Price

	return nil
}

// func (product Product) String() string {
// 	return fmt.Sprintf("%s %s\n%s\n%s\n%s\n%s\n", product.FirstName, product.LastName, product.Email, product.DateOfBirth.Format("01-02-2006"), product.Address, product.Balance)
// }
