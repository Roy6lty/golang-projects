package database

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrCantFindProduct    = errors.New("product not found")
	ErrCantDecodeProducts = errors.New("cant decode product")
	ErrUserIdIsInvalid    = errors.New("Invalid User ID")
	ErrCantUpdateUser     = errors.New("can't Add product to cart")
	ErrCantRemoveItemCart = errors.New("cannot remove item from cart")
	ErrCantGetItem        = errors.New(" ws unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot purchase item")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}

func BuyfromCart() gin.HandlerFunc {

}

func InstantBuy() gin.HandlerFunc {

}
