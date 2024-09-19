package valueobjects

import "fmt"

type ProductCategoryType = string

const (
	TypeLanche         ProductCategoryType = "Lanche"
	TypeAcompanhamento ProductCategoryType = "Acompanhamento"
	TypeBebida         ProductCategoryType = "Bebida"
	TypeSobremesa      ProductCategoryType = "Sobremesa"
)

var validProductCategories = map[OrderStatusType]OrderStatusType{
	TypeLanche:         TypeLanche,
	TypeAcompanhamento: TypeAcompanhamento,
	TypeBebida:         TypeBebida,
	TypeSobremesa:      TypeSobremesa,
}

func ParseToProductCategoryType(category string) (OrderStatusType, error) {
	categoryType, ok := validProductCategories[category]
	if !ok {
		return "", fmt.Errorf("category '%s' is not valid", category)
	}

	return categoryType, nil
}
