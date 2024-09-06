package valueobjects

type OrderStatusType string

const (
	TypeReceived        OrderStatusType = "Received"
	TypeAwaitingPayment OrderStatusType = "Awaiting Payment"
	TypeConfirmed       OrderStatusType = "Confirmed"
	TypePreparing       OrderStatusType = "Preparing"
	TypeReady           OrderStatusType = "Ready"
	TypeFinished        OrderStatusType = "Finished"
	TypeCanceled        OrderStatusType = "Canceled"
)
