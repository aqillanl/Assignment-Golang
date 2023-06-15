package dto

import (
	"assignment-2/entity"
	"time"
)

type NewOrderRequest struct {
	OrderedAt    time.Time        
	CustomerName string           
	Items        []NewItemRequest 
}

func (o *NewOrderRequest) OrderRequestToEntity() *entity.Order {
	return &entity.Order{
		CustomerName: o.CustomerName,
		OrderedAt:    o.OrderedAt,
	}
}

type NewOrderResponse struct {
	StatusCode int             
	Message    string          
	Data       NewOrderRequest 
}

type GetAllOrdersResponse struct {
	StatusCode int         
	Message    string      
	Data       []OrderData 
}

type GetOrderByIDResponse struct {
	StatusCode int       
	Message    string    
	Data       OrderData 
}

type DeleteOrderByIDResponse struct {
	StatusCode int    
	Message    string 
}

type OrderData struct {
	ID           uint       
	CreatedAt    time.Time  
	UpdatedAt    time.Time  
	CustomerName string     
	Items        []ItemData 
}
