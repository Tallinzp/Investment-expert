package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var err error
	var tradeType int
	var tradePrice float64
	var tradeNum int
	var market int
	for id, arg := range os.Args {
		if id == 1 {
			if strings.Compare(arg, "h") == 0 {
				fmt.Println(
					`format: ./tax [12]+ price stock_number [12]+
第一个参数的填入规则：1 - 买，2 - 卖
第二个参数的填入规则：price - 股票交易价格
第三个参数的填入规则：stock_number - 股票交易数量
第四个参数的填入规则：1 - 泸市, 2 - 深市
			`)
			}
			tradeType, err = strconv.Atoi(arg)
			if err != nil || tradeType > 2 || tradeType < 1 {
				fmt.Println("第一个参数错误，买入填1，卖出填2")
			}
		}

		if id == 2 {
			tradePrice, err = strconv.ParseFloat(arg, 64)
			if err != nil || tradePrice < 0 {
				fmt.Println("第二个参数错误，请填入股票交易价格")
			}
		}

		if id == 3 {
			tradeNum, err = strconv.Atoi(arg)
			if err != nil || tradeNum < 0 {
				fmt.Println("第三个参数错误，请填入股票交易数量")
			}
		}

		if id == 4 {
			market, err = strconv.Atoi(arg)
			if err != nil || market > 2 || market < 1 {
				fmt.Println("第四个参数错误，请填入交易所，泸市填1，深市填2")
			}
		}
	}

	tradeAmount := tradePrice * float64(tradeNum)
	fee := 0.0

	fmt.Println("交易金额：", int(tradeAmount))
	fmt.Println("交易费用：")

	// 佣金｜收费方：券商
	commissionRate := 0.00025
	commission := tradeAmount * commissionRate
	if commission < 5 {
		commission = 5
	}
	fee += commission
	fmt.Println("\t券商佣金：", commission)

	// 印花税(仅向卖方收取)｜收费方：税务局
	if tradeType == 2 {
		stampDutyRate := 0.001
		stampDuty := tradeAmount * stampDutyRate
		fee += stampDuty
		fmt.Println("\t印花税", stampDuty)
	}

	// 过户费(仅泸市收取，深市不受)｜收费方：上海交易所
	if market == 1 {
		transferFeeRate := 0.00002
		transferFee := tradeAmount * transferFeeRate
		fee += transferFee
		fmt.Println("\t泸市过户费", transferFee)
	}

	fmt.Println("交易费用汇总：", fee)

	// 买入
	if tradeType == 1 {
		fmt.Println("实际成交价格：", (tradeAmount+fee)/float64(tradeNum))
	}

	if tradeType == 2 {
		fmt.Println("实际卖出所得：", tradeAmount-fee)
	}
}
