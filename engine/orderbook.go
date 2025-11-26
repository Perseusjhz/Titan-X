package engine

import (
	"fmt"

	"github.com/Perseusjhz/Titan-X/models" // 👈 注意：这里引入了刚才写的 models 包
)

// OrderBook: 订单簿
type OrderBook struct {
	BuyOrders  []models.Order
	SellOrders []models.Order
}

// PlaceOrder: 下单
func (ob *OrderBook) PlaceOrder(o models.Order) {
	fmt.Printf("📝 [下单] 用户 %s: %s %s 价格:%s 数量:%s\n",
		o.User.Name, o.Side, o.Symbol, o.Price.String(), o.Quantity.String())

	if o.Side == "BUY" {
		ob.BuyOrders = append(ob.BuyOrders, o)
	} else {
		ob.SellOrders = append(ob.SellOrders, o)
	}
}

// Match: 撮合逻辑
func (ob *OrderBook) Match() {
	fmt.Println("⚙️ [引擎] 开始撮合...")

	if len(ob.BuyOrders) == 0 || len(ob.SellOrders) == 0 {
		fmt.Println("📭 订单簿为空或单边，无法撮合")
		return
	}

	buy := ob.BuyOrders[0]
	sell := ob.SellOrders[0]

	// 核心修改：使用 decimal 的比较方法
	// buy.Price >= sell.Price
	if buy.Price.GreaterThanOrEqual(sell.Price) {
		fmt.Printf("💥 [成交] 买方: %s | 卖方: %s | 价格: %s\n",
			buy.User.Name, sell.User.Name, sell.Price.String())

		// 这里以后会加扣钱逻辑
	} else {
		fmt.Println("💤 价格不匹配")
	}
}
