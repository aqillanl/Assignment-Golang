package dto

import (
	"assignment-2/entity"
	"time"
)

type NewItemRequest struct {
	ItemCode    string 
	Description string 
	Quantity    uint   
}

func (i *NewItemRequest) ItemRequestToEntity() *entity.Item {
	return &entity.Item{
		ItemCode:    i.ItemCode,
		Description: i.Description,
		Quantity:    i.Quantity,
	}
}

type ItemData struct {
	ID          uint      
	CreatedAt   time.Time 
	UpdatedAt   time.Time 
	ItemCode    string    
	Description string    
	Quantity    uint      
	OrderID     uint      
}

type UpdateItemResponse struct {
	StatusCode int      
	Message    string   
	Data       ItemData 
}
