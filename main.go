package main

import "fmt"

type ServerMachine struct {
	ID        string
	CPU       int
	RAM       int
	BasePrice int
}

func FullPrice(servers []ServerMachine) {
	for _, server := range servers {
		allowance := func(s ServerMachine) int {
			if s.CPU >= 8 || s.RAM >= 32 {
				return 30
			} else if s.CPU >= 4 || s.RAM >= 16 {
				return 15
			}
			return 0
		}(server)

		calculatorLoyalty := func(base int) int {
			if base > 10_000 {
				return 1_000
			}
			return 0
		}
		loyalty := calculatorLoyalty(server.BasePrice)
		surchargeAmount := server.BasePrice * allowance / 100
		finalPrice := server.BasePrice + surchargeAmount - loyalty
		fmt.Printf("Надбавка %d%% = %d руб. Скидка %d руб. Итого: %d + %d - %d = %d руб\n",
			allowance, surchargeAmount, loyalty, server.BasePrice, surchargeAmount, loyalty, finalPrice)

	}
}

func main() {
	servers := []ServerMachine{
		{ID: "SRV-HIGH", CPU: 8, RAM: 32, BasePrice: 12_000},
		{ID: "SRV-MID", CPU: 4, RAM: 16, BasePrice: 6_000},
		{ID: "SRV-MINI", CPU: 2, RAM: 4, BasePrice: 2_000},
	}
	FullPrice(servers)
}
