package valueobjects

import "fmt"

type OrderStatusType = string

const (
	TypeReceived        OrderStatusType = "Received"
	TypeAwaitingPayment OrderStatusType = "Awaiting Payment"
	TypeConfirmed       OrderStatusType = "Confirmed"
	TypePreparing       OrderStatusType = "Preparing"
	TypeReady           OrderStatusType = "Ready"
	TypeFinished        OrderStatusType = "Finished"
	TypeCanceled        OrderStatusType = "Canceled"
)

var validOrderStatuses = map[OrderStatusType]OrderStatusType{
	TypeReceived:        TypeReceived,
	TypeAwaitingPayment: TypeAwaitingPayment,
	TypeConfirmed:       TypeConfirmed,
	TypePreparing:       TypePreparing,
	TypeReady:           TypeReady,
	TypeFinished:        TypeFinished,
	TypeCanceled:        TypeCanceled,
}

func ParseToOrderStatusType(status string) (OrderStatusType, error) {
	orderStatus, ok := validOrderStatuses[status]
	if !ok {
		return "", fmt.Errorf("order status '%s' is not valid", status)
	}

	return orderStatus, nil
}
