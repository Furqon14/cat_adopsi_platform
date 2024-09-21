package controller

import (
	"cat_adoption_platform/config"
	"cat_adoption_platform/model/dto"
	"cat_adoption_platform/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veritrans/go-midtrans"
)

type MidtransController struct {
	service service.MidtransService
	rg      *gin.RouterGroup
}

func (c *MidtransController) PaymentHandler(ctx *gin.Context) {
	payload := dto.ChargeMidtrans{}
	// payload := strings.NewReader("{\"payment_type\":\"gopay\",\"transaction_details\":{\"order_id\":\"order-id-123\",\"gross_amount\":100000},\"customer_details\":{\"first_name\":\"Budi\",\"last_name\":\"Utomo\",\"email\":\"budi.utomo@midtrans.com\",\"phone\":\"081223323423\",\"customer_details_required_fields\":[\"email\",\"first_name\",\"phone\"]},\"custom_field1\":\"custom field 1 content\",\"custom_field2\":\"custom field 2 content\",\"custom_field3\":\"custom field 3 content\",\"custom_expiry\":{\"expiry_duration\":60,\"unit\":\"minute\"},\"metadata\":{\"you\":\"can\",\"put\":\"any\",\"parameter\":\"you like\"}}")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	midclient := config.InitMidtrans()

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	var chargeReq *midtrans.ChargeReq

	// Check the payment method from the payload
	switch payload.PaymentType {
	case "credit_card":
		chargeReq = &midtrans.ChargeReq{
			PaymentType: midtrans.SourceCreditCard,
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  payload.TransactionDetails.OrderID,
				GrossAmt: payload.TransactionDetails.GrossAmount,
			},
			CustomerDetail: &midtrans.CustDetail{
				FName: payload.CustomerDetails.FirstName,
				LName: payload.CustomerDetails.LastName,
				Email: payload.CustomerDetails.Email,
				Phone: payload.CustomerDetails.Phone,
			},
		}
	case "gopay":
		chargeReq = &midtrans.ChargeReq{
			PaymentType: midtrans.SourceGopay,
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  payload.TransactionDetails.OrderID,
				GrossAmt: payload.TransactionDetails.GrossAmount,
			},
			CustomerDetail: &midtrans.CustDetail{
				FName: payload.CustomerDetails.FirstName,
				LName: payload.CustomerDetails.LastName,
				Email: payload.CustomerDetails.Email,
				Phone: payload.CustomerDetails.Phone,
			},
		}
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported payment method"})
		return
	}

	resp, err := coreGateway.Charge(chargeReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (mts *MidtransController) Route() {
	// Your implementation here
	router := mts.rg.Group("/charge")
	router.POST("", mts.PaymentHandler)
	// Add more routes here as needed...
}

func NewMidtransController(service service.MidtransService, rg *gin.RouterGroup) *MidtransController {
	return &MidtransController{
		service: service,
		rg:      rg,
	}
}
