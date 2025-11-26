package main // 1. package: 告诉计算机这是一个可执行程序，而不是一个库

import (
	"fmt" // 2. import: 引入标准工具箱，fmt 用于在屏幕上打印文字
)

// --- 数据结构定义 (相当于金融里的“会计科目”定义) ---

// Account: 定义一个“账户”长什么样
type Account struct {
	ID      int     // 用户ID (整数，比如 1, 2, 3)
	Name    string  // 用户名 (字符串，比如 "Jiao")
	Balance float64 // 余额 (浮点数，比如 1000.50)
	// 在正式项目中，这里会加上 "锁" 来防止并发冲突
}

// Order: 定义一张“订单”长什么样
type Order struct {
	ID        int     // 订单号
	Symbol    string  // 交易对，比如 "BTC/USDT"
	Side      string  // 方向: "BUY"(买) 或 "SELL"(卖)
	Price     float64 // 价格
	Quantity  float64 // 数量
	User      *Account // 归属: 这张单子是谁下的？(* 代表这是一个指向Account的指针)
}

// OrderBook: 定义“订单簿”，用来存放所有挂单
type OrderBook struct {
	BuyOrders  []Order // 买单列表 ([] 代表数组/切片，可以放很多个Order)
	SellOrders []Order // 卖单列表
}

// --- 核心功能 (相当于业务部门的操作) ---

// Deposit: 充值功能
// (u *Account) 意思是：这个功能是专门给 Account 类型用的
// * 号很关键：代表我们要修改“真身”的数据，而不是修改复印件
func (u *Account) Deposit(amount float64) {
	u.Balance = u.Balance + amount
	fmt.Printf("✅ [系统] 用户 %s 充值成功! 当前余额: %.2f\n", u.Name, u.Balance)
}

// PlaceOrder: 下单功能
// ob *OrderBook: 把订单放入订单簿
func (ob *OrderBook) PlaceOrder(o Order) {
	fmt.Printf("📝 [下单] 用户 %s 提交了订单: %s %s 价格:%.2f 数量:%.2f\n", 
		o.User.Name, o.Side, o.Symbol, o.Price, o.Quantity)

	// 简单的逻辑判断
	if o.Side == "BUY" {
		// append 意思是把新订单追加到列表末尾
		ob.BuyOrders = append(ob.BuyOrders, o)
	} else if o.Side == "SELL" {
		ob.SellOrders = append(ob.SellOrders, o)
	}
}

// Match: 撮合引擎 (核心中的核心)
// 这里写一个最最简单的逻辑：只要有买单和卖单，就看价格能不能成交
func (ob *OrderBook) Match() {
	fmt.Println("⚙️ [引擎] 开始尝试撮合交易...")

	// 循环检查买单和卖单
	// len() 用来获取列表长度
	if len(ob.BuyOrders) > 0 && len(ob.SellOrders) > 0 {
		buy := ob.BuyOrders[0]   // 取出第一个买单
		sell := ob.SellOrders[0] // 取出第一个卖单

		// 如果 买单价格 >= 卖单价格，说明可以成交
		if buy.Price >= sell.Price {
			fmt.Printf("💥 [成交] 撮合成功! %s 买入 BTC, 卖家是 %s, 成交价: %.2f\n", 
				buy.User.Name, sell.User.Name, sell.Price)
			
			// 这里未来要写：扣钱、加币的逻辑
		} else {
			fmt.Println("💤 [引擎] 价格不匹配，无法成交。")
		}
	} else {
		fmt.Println("📭 [引擎] 订单簿为空，等待更多订单...")
	}
}

// --- 主程序入口 (一切从这里开始执行) ---
func main() {
	fmt.Println("🚀 Titan-X 交易所系统启动中...")

	// 1. 初始化一个空的订单簿
	book := OrderBook{}

	// 2. 创建两个用户 (结构体实例化)
	user1 := &Account{ID: 1, Name: "Perseus", Balance: 0}
	user2 := &Account{ID: 2, Name: "MarketMaker", Balance: 0}

	// 3. 模拟资金流动
	user1.Deposit(10000.00) // 你的账户充值
	user2.Deposit(5.00)     // 做市商充值 BTC (这里简化演示)

	// 4. 用户1 下一个买单：我想用 60000块 买 0.1个BTC
	buyOrder := Order{
		ID: 101, Symbol: "BTC/USDT", Side: "BUY", 
		Price: 60000, Quantity: 0.1, User: user1,
	}
	book.PlaceOrder(buyOrder)

	// 5. 用户2 下一个卖单：我愿意 59000块 卖 0.1个BTC
	sellOrder := Order{
		ID: 201, Symbol: "BTC/USDT", Side: "SELL", 
		Price: 59000, Quantity: 0.1, User: user2,
	}
	book.PlaceOrder(sellOrder)

	// 6. 触发撮合引擎
	book.Match()
}