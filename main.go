package main

import (
	"fmt"

	"github.com/shopspring/decimal"

	// 👇 确保这里的路径和你 go.mod 第一行写的完全一致
	"github.com/Perseusjhz/Titan-X/engine"
	"github.com/Perseusjhz/Titan-X/models"
)

func main() {
	fmt.Println("🚀 Titan-X Pro版 (多币种结算) 启动...")

	// 1. 初始化引擎
	book := engine.OrderBook{}

	// 2. 创建账户
	// ❌ 错误写法 (旧版): user := &models.Account{..., Balance: ...}
	// ✅ 正确写法 (新版): 初始化时不填钱，因为 Balance 字段已经没了
	user1 := &models.Account{ID: 1, Name: "Perseus"}
	user2 := &models.Account{ID: 2, Name: "MarketMaker"}

	// 3. 初始资金注入 (通过方法充值)
	// Perseus 有 100,000 USDT
	user1.Deposit("USDT", decimal.NewFromFloat(100000.0))

	// MarketMaker 有 10 BTC
	user2.Deposit("BTC", decimal.NewFromFloat(10.0))

	fmt.Println("--- 初始状态 ---")
	// 注意这里访问的是 Balances["USDT"]
	fmt.Printf("用户 %s: USDT=%s\n", user1.Name, user1.Balances["USDT"])
	fmt.Printf("用户 %s: BTC=%s\n", user2.Name, user2.Balances["BTC"])
	fmt.Println("----------------")

	// 4. 创建订单
	// 买单: 我想用 60000 的价格买 0.1 BTC
	buyOrder := models.Order{
		ID: 101, Symbol: "BTC/USDT", Side: "BUY",
		Price:    decimal.NewFromFloat(60000.0),
		Quantity: decimal.NewFromFloat(0.1),
		User:     user1,
	}

	// 卖单: 我愿意 59000 卖 0.1 BTC
	sellOrder := models.Order{
		ID: 201, Symbol: "BTC/USDT", Side: "SELL",
		Price:    decimal.NewFromFloat(59000.0),
		Quantity: decimal.NewFromFloat(0.1),
		User:     user2,
	}

	// 5. 下单并撮合
	book.PlaceOrder(buyOrder)
	book.PlaceOrder(sellOrder)
	book.Match()
}
