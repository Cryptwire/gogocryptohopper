package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cryptwire/gogocryptohopper"
	"github.com/cryptwire/gogocryptohopper/pkg"
)

func main() {
	fmt.Println("Example: Fetch hoppers")

	ch, err := gogocryptohopper.New(
		"https://www.cryptohopper.com",
		"https://api.cryptohopper.com/v1",
		accessToken)
	if err != nil {
		panic(err)
	}

	// fetchAllHoppers(ch)
	// fetchSingleHopper("CH_HOPPER_ID", ch)
	// fetchHopperConfig("CH_HOPPER_ID", ch)
	// fetchOpenPositions("CH_HOPPER_ID", ch)
	// fetchReservedPositions("CH_HOPPER_ID", ch)
	// fetchShortPositions("CH_HOPPER_ID", ch)
	// fetchTradeHistory("CH_HOPPER_ID", ch)

	// Here you can fetch trades from a specific timeframe
	// from, err := time.Parse("2006-01-02 15:04:05", "2021-08-01 00:00:00")
	// if err != nil {
	// 	panic(err)
	// }
	// to := time.Now().Add(time.Hour * 24)
	// fetchTradeHistoryWithTimeframe("CH_HOPPER_ID", from, to, ch)

	// fetchHopperStats("CH_HOPPER_ID", ch)
	// fetchHopperAssets("CH_HOPPER_ID", ch)

	// fetchHopperLogs("CH_HOPPER_ID", ch)

	hopperPanicSell("CH_HOPPER_ID", ch)

}

func fetchAllHoppers(ch *gogocryptohopper.GoGoCryptohopper) {
	result, err := ch.Hopper.GetAll()
	if err != nil {
		panic(err)
	}

	for _, hopper := range result {
		fmt.Println("========HOPPER==========")
		fmt.Println("ID: ", hopper.ID)
		fmt.Println("Name: ", hopper.Name)
		fmt.Println("Exchange: ", hopper.Exchange)
		fmt.Println("HopperID: ", hopper.HopperID)
		fmt.Println("SubscriptionID: ", hopper.SubscriptionID)
		fmt.Println("PaymentTerm: ", hopper.PaymentTerm)
		fmt.Println("PaymentMethodID: ", hopper.PaymentMethodID)
		fmt.Println("StartTime: ", hopper.StartTime)
		fmt.Println("EndTime: ", hopper.EndTime)
		fmt.Println("SubscriptionStatus: ", hopper.SubscriptionStatus)
		fmt.Println("AutoRenewal: ", hopper.AutoRenewal)
		fmt.Println("Subscription: ", hopper.Subscription)
		fmt.Println("PlanName: ", hopper.PlanName)
		fmt.Println("PlanDescription: ", hopper.PlanDescription)
		fmt.Println("ProductID: ", hopper.ProductID)
		fmt.Println("LastLoadedConfig: ", hopper.LastLoadedConfig)
		fmt.Println("Image: ", hopper.Image)
		fmt.Println("BaseCurrency: ", hopper.BaseCurrency)
		fmt.Println("IsBuyingEnabled: ", hopper.IsBuyingEnabled)
		fmt.Println("IsSellingEnabled: ", hopper.IsSellingEnabled)
		fmt.Println("Enabled: ", hopper.Enabled)
		fmt.Println("SetDefault: ", hopper.SetDefault)
		fmt.Println("ErrorMessage: ", hopper.ErrorMessage)
		fmt.Println("LastSignal: ", hopper.LastSignal)
		fmt.Println("LastSignalEncoding: ", hopper.LastSignalEncoding)
		fmt.Println("TotalCur: ", hopper.CurrentTotal)
		fmt.Println("ConfigError: ", hopper.ConfigError)
		fmt.Println("Created: ", hopper.Created)
		fmt.Println("BotType: ", hopper.BotType)
		fmt.Println("UserID: ", hopper.UserID)
		fmt.Println("AllowedCoins: ", hopper.AllowedCoins)
		fmt.Println("PaperTrading: ", hopper.PaperTrading)
		fmt.Println("StartBalance: ", hopper.StartBalance)
	}
}

func fetchSingleHopper(id string, ch *gogocryptohopper.GoGoCryptohopper) {

	hopper, err := ch.Hopper.Get(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("==================")
	fmt.Println("ID: ", hopper.ID)
	fmt.Println("Name: ", hopper.Name)
	fmt.Println("Exchange: ", hopper.Exchange)
	fmt.Println("HopperID: ", hopper.HopperID)
	fmt.Println("SubscriptionID: ", hopper.SubscriptionID)
	fmt.Println("PaymentTerm: ", hopper.PaymentTerm)
	fmt.Println("PaymentMethodID: ", hopper.PaymentMethodID)
	fmt.Println("StartTime: ", hopper.StartTime)
	fmt.Println("EndTime: ", hopper.EndTime)
	fmt.Println("SubscriptionStatus: ", hopper.SubscriptionStatus)
	fmt.Println("AutoRenewal: ", hopper.AutoRenewal)
	fmt.Println("Subscription: ", hopper.Subscription)
	fmt.Println("PlanName: ", hopper.PlanName)
	fmt.Println("PlanDescription: ", hopper.PlanDescription)
	fmt.Println("ProductID: ", hopper.ProductID)
	fmt.Println("LastLoadedConfig: ", hopper.LastLoadedConfig)
	fmt.Println("Image: ", hopper.Image)
	fmt.Println("BaseCurrency: ", hopper.BaseCurrency)
	fmt.Println("IsBuyingEnabled: ", hopper.IsBuyingEnabled)
	fmt.Println("IsSellingEnabled: ", hopper.IsSellingEnabled)
	fmt.Println("Enabled: ", hopper.Enabled)
	fmt.Println("SetDefault: ", hopper.SetDefault)
	fmt.Println("ErrorMessage: ", hopper.ErrorMessage)
	fmt.Println("LastSignal: ", hopper.LastSignal)
	fmt.Println("LastSignalEncoding: ", hopper.LastSignalEncoding)
	fmt.Println("TotalCur: ", hopper.CurrentTotal)
	fmt.Println("ConfigError: ", hopper.ConfigError)
	fmt.Println("Created: ", hopper.Created)
	fmt.Println("BotType: ", hopper.BotType)
	fmt.Println("UserID: ", hopper.UserID)
	fmt.Println("AllowedCoins: ", hopper.AllowedCoins)
	fmt.Println("PaperTrading: ", hopper.PaperTrading)
	fmt.Println("StartBalance: ", hopper.StartBalance)
}

func fetchHopperConfig(id string, ch *gogocryptohopper.GoGoCryptohopper) {
	config, err := ch.Hopper.GetBaseconfig(id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name: ", config.Name)
	fmt.Println("Live: ", config.Live)
	fmt.Println("LastLoadedConfig: ", config.LastLoadedConfig)
	fmt.Println("LastLoadedConfigDate: ", config.LastLoadedConfigDate)
	fmt.Println("PaperTradingAccount: ", config.PaperTradingAccount)
	fmt.Println("Exchange: ", config.Exchange)
	fmt.Println("APIKey: ", config.APIKey)
	fmt.Println("APISecret: ", config.APISecret)
	fmt.Println("APIPassphrase: ", config.APIPassphrase)
	fmt.Println("ExtraAPIKey: ", config.ExtraAPIKey)
	fmt.Println("ExtraAPISecret: ", config.ExtraAPISecret)
	fmt.Println("TickerType: ", config.TickerType)
	fmt.Println("AllowedCoins: ", config.AllowedCoins)
	fmt.Println("AllowAllCoins: ", config.AllowAllCoins)
	fmt.Println("PercBuyAmount: ", config.PercBuyAmount)
	fmt.Println("MinBuyAmount: ", config.MinBuyAmount)
	fmt.Println("MaxAmountAllocated: ", config.MaxAmountAllocated)
	fmt.Println("Strategy: ", config.Strategy)
	fmt.Println("NumTargetsPerBuy: ", config.NumTargetsPerBuy)
	fmt.Println("AdvancedTaCandleSize: ", config.AdvancedTaCandleSize)
	fmt.Println("AdvancedTaStochFastK: ", config.AdvancedTaStochFastK)
	fmt.Println("AdvancedTaStochSlowK: ", config.AdvancedTaStochSlowK)
	fmt.Println("AdvancedTaStochSlowKMatype: ", config.AdvancedTaStochSlowKMatype)
	fmt.Println("AdvancedTaStochSlowD: ", config.AdvancedTaStochSlowD)
	fmt.Println("AdvancedTaStochSlowDMatype: ", config.AdvancedTaStochSlowDMatype)
	fmt.Println("AdvancedTaStochOversold: ", config.AdvancedTaStochOversold)
	fmt.Println("AdvancedTaStochOverbought: ", config.AdvancedTaStochOverbought)
	fmt.Println("AdvancedTaStochrsiPeriod: ", config.AdvancedTaStochrsiPeriod)
	fmt.Println("AdvancedTaStochrsiFastK: ", config.AdvancedTaStochrsiFastK)
	fmt.Println("AdvancedTaStochrsiFastD: ", config.AdvancedTaStochrsiFastD)
	fmt.Println("AdvancedTaStochrsiFastDMatype: ", config.AdvancedTaStochrsiFastDMatype)
	fmt.Println("AdvancedTaStochrsiOversold: ", config.AdvancedTaStochrsiOversold)
	fmt.Println("AdvancedTaStochrsiOverbought: ", config.AdvancedTaStochrsiOverbought)
	fmt.Println("AdvancedTaRsiPeriod: ", config.AdvancedTaRsiPeriod)
	fmt.Println("AdvancedTaRsiOversold: ", config.AdvancedTaRsiOversold)
	fmt.Println("AdvancedTaRsiOverbought: ", config.AdvancedTaRsiOverbought)
	fmt.Println("AdvancedTaMacdFastPeriod: ", config.AdvancedTaMacdFastPeriod)
	fmt.Println("AdvancedTaMacdSlowPeriod: ", config.AdvancedTaMacdSlowPeriod)
	fmt.Println("AdvancedTaMacdSignalPeriod: ", config.AdvancedTaMacdSignalPeriod)
	fmt.Println("AdvancedTaEmaShortPeriod: ", config.AdvancedTaEmaShortPeriod)
	fmt.Println("AdvancedTaEmaLongPeriod: ", config.AdvancedTaEmaLongPeriod)
	fmt.Println("AdvancedTaSmaShortPeriod: ", config.AdvancedTaSmaShortPeriod)
	fmt.Println("AdvancedTaSmaLongPeriod: ", config.AdvancedTaSmaLongPeriod)
	fmt.Println("AdvancedTaKamaShortPeriod: ", config.AdvancedTaKamaShortPeriod)
	fmt.Println("AdvancedTaKamaLongPeriod: ", config.AdvancedTaKamaLongPeriod)
	fmt.Println("AdvancedTaBbandsPeriod: ", config.AdvancedTaBbandsPeriod)
	fmt.Println("AdvancedTaBbandsDevUp: ", config.AdvancedTaBbandsDevUp)
	fmt.Println("AdvancedTaBbandsDevLow: ", config.AdvancedTaBbandsDevLow)
	fmt.Println("AdvancedTaBbandsMatype: ", config.AdvancedTaBbandsMatype)
	fmt.Println("BuyScoreCorrected: ", config.BuyScoreCorrected)
	fmt.Println("MinBuyScore: ", config.MinBuyScore)
	fmt.Println("SellScoreCorrected: ", config.SellScoreCorrected)
	fmt.Println("MinSellScore: ", config.MinSellScore)
	fmt.Println("RsiCandleSize: ", config.RsiCandleSize)
	fmt.Println("RsiPeriod: ", config.RsiPeriod)
	fmt.Println("RsiOversold: ", config.RsiOversold)
	fmt.Println("BbandsCandleSize: ", config.BbandsCandleSize)
	fmt.Println("BbandsDays: ", config.BbandsDays)
	fmt.Println("BbandsDeviation: ", config.BbandsDeviation)
	fmt.Println("BuyOrderType: ", config.BuyOrderType)
	fmt.Println("BidPercentage: ", config.BidPercentage)
	fmt.Println("BidPercentageType: ", config.BidPercentageType)
	fmt.Println("MaxOpenTimeBuy: ", config.MaxOpenTimeBuy)
	fmt.Println("MaxOpenPositions: ", config.MaxOpenPositions)
	fmt.Println("MaxOpenPositionsPerCoin: ", config.MaxOpenPositionsPerCoin)
	fmt.Println("Cooldown: ", config.Cooldown)
	fmt.Println("CooldownWhen: ", config.CooldownWhen)
	fmt.Println("CooldownCount: ", config.CooldownCount)
	fmt.Println("CooldownVal: ", config.CooldownVal)
	fmt.Println("OneOpenOrder: ", config.OneOpenOrder)
	fmt.Println("OnlyWhenPositiveTime: ", config.OnlyWhenPositiveTime)
	fmt.Println("Arbitrage: ", config.Arbitrage)
	fmt.Println("StopLossPercentage: ", config.StopLossPercentage)
	fmt.Println("StopLossTrailing: ", config.StopLossTrailing)
	fmt.Println("StopLossTrailingPercentage: ", config.StopLossTrailingPercentage)
	fmt.Println("StopLossTrailingArm: ", config.StopLossTrailingArm)
	fmt.Println("StopLossTrailingOnly: ", config.StopLossTrailingOnly)
	fmt.Println("TrailingStopLossProfit: ", config.TrailingStopLossProfit)
	fmt.Println("TrailingBuyPercentage: ", config.TrailingBuyPercentage)
	fmt.Println("MaxOpenShortPositions: ", config.MaxOpenShortPositions)
	fmt.Println("ShortPercentageProfit: ", config.ShortPercentageProfit)
	fmt.Println("ShortStopLossTrailingPercentage: ", config.ShortStopLossTrailingPercentage)
	fmt.Println("ShortStopLossTrailingArm: ", config.ShortStopLossTrailingArm)
	fmt.Println("ShortAutoClosePositionsTime: ", config.ShortAutoClosePositionsTime)
	fmt.Println("ShortAutoRemovePositionsTime: ", config.ShortAutoRemovePositionsTime)
	fmt.Println("ShortRemoveOnLoss: ", config.ShortRemoveOnLoss)
	fmt.Println("HoldAssets: ", config.HoldAssets)
	fmt.Println("AutoClosePositionsTime: ", config.AutoClosePositionsTime)
	fmt.Println("DcaOrderType: ", config.DcaOrderType)
	fmt.Println("AutoDcaTime: ", config.AutoDcaTime)
	fmt.Println("AutoDcaMax: ", config.AutoDcaMax)
	fmt.Println("AutoDcaPercentage: ", config.AutoDcaPercentage)
	fmt.Println("AutoDcaSize: ", config.AutoDcaSize)
	fmt.Println("AutoDcaSizeCustom: ", config.AutoDcaSizeCustom)
	fmt.Println("SetPercentage: ", config.SetPercentage)
	fmt.Println("SellOrderType: ", config.SellOrderType)
	fmt.Println("MaxOpenTime: ", config.MaxOpenTime)
	fmt.Println("AskPercentage: ", config.AskPercentage)
	fmt.Println("AskPercentageType: ", config.AskPercentageType)
	fmt.Println("Submit: ", config.Submit)
	fmt.Println("StopBuying: ", config.StopBuying)
	fmt.Println("StopSelling: ", config.StopSelling)
	fmt.Println("WalletscrubberLeftovers: ", config.WalletscrubberLeftovers)
	fmt.Println("WalletscrubberAuto: ", config.WalletscrubberAuto)
	fmt.Println("WalletscrubberNoopen: ", config.WalletscrubberNoopen)
	fmt.Println("CollectCurrency: ", config.CollectCurrency)
	fmt.Println("OutputLiveFeed: ", config.OutputLiveFeed)
	fmt.Println("OutputErrorsOnly: ", config.OutputErrorsOnly)
	fmt.Println("Autosync: ", config.Autosync)
	fmt.Println("AutosyncAllCoins: ", config.AutosyncAllCoins)
	fmt.Println("AutosyncAllowedCoins: ", config.AutosyncAllowedCoins)
	fmt.Println("SendTradeEmail: ", config.SendTradeEmail)
	fmt.Println("SendTradeErrorEmail: ", config.SendTradeErrorEmail)
	fmt.Println("SendCancelledEmail: ", config.SendCancelledEmail)
	fmt.Println("BuyBtcAmount: ", config.BuyBtcAmount)
}

func fetchOpenPositions(id string, ch *gogocryptohopper.GoGoCryptohopper) {
	result, err := ch.Hopper.GetOpenPositions(id)
	if err != nil {
		panic(err)
	}

	for _, position := range result {
		fmt.Println("Open Position: ", position.ID)
		fmt.Println(" - Rate", position.Rate)
		fmt.Println(" - Cost", position.Cost)
		fmt.Println(" - Currency", position.Currency)
		fmt.Println(" - Amount", position.Amount)
		fmt.Println(" - Bought at", position.BuyTime)
		fmt.Println(" - Because", position.BuyTrigger)
	}
	fmt.Println("Total positions", len(result))
}

func fetchReservedPositions(id string, ch *gogocryptohopper.GoGoCryptohopper) {
	result, err := ch.Hopper.GetReservedPositions(id)
	if err != nil {
		panic(err)
	}

	for _, position := range result {
		fmt.Println("Reserved Position: ", position.ID)
		fmt.Println(" - Pair", position.Pair)
		fmt.Println(" - Rate", position.Rate)
		fmt.Println(" - Cost", position.ReservedCost)
		fmt.Println(" - Currency", position.Currency)
		fmt.Println(" - Base currency", position.ReservedBaseCurrency)
		fmt.Println(" - Amount", position.Amount)
		fmt.Println(" - Bought at", position.BuyTime)
		fmt.Println(" - Because", position.BuyTrigger)
	}
	fmt.Println("Total positions", len(result))
}

func fetchShortPositions(id string, ch *gogocryptohopper.GoGoCryptohopper) {
	result, err := ch.Hopper.GetShortPositions(id)
	if err != nil {
		panic(err)
	}

	for _, position := range result {
		fmt.Println("Short Position: ", position.ID)
		fmt.Println(" - Pair", position.Pair)
		fmt.Println(" - Rate", position.Rate)
		fmt.Println(" - Cost", position.Cost)
		fmt.Println(" - Currency", position.Currency)
		fmt.Println(" - Amount", position.Amount)
		fmt.Println(" - Start rate", position.ShortStartRate)
		fmt.Println(" - Short profit", position.ShortStartProfit)
		fmt.Println(" - Short at", position.SellTime)
		fmt.Println(" - Because", position.SellTrigger)
	}
	fmt.Println("Total positions", len(result))
}

func fetchTradeHistory(id string, ch *gogocryptohopper.GoGoCryptohopper) {
	result, err := ch.Hopper.GetTradeHistory(id)
	if err != nil {
		panic(err)
	}

	for _, trade := range result {
		fmt.Println("TradeID: ", trade.ID)
	}
	fmt.Println("Total positions", len(result))
}

func fetchTradeHistoryWithTimeframe(
	id string,
	dateFrom time.Time,
	dateTo time.Time,
	ch *gogocryptohopper.GoGoCryptohopper) {

	offset := 0
	results := 500

	totalTradesReturned := 0
	var totalTrades []pkg.Trade

	for results >= 500 {
		historyInfo, err := ch.Hopper.GetTradeHistoryWithTimeframes(id, dateFrom, dateTo, offset)
		if err != nil {
			panic(err)
		}

		for _, trade := range historyInfo.Trades {
			fmt.Printf("TradeID: %s, Date: %s\n", trade.OrderID, trade.Date)
			// break
		}

		switch v := historyInfo.Results.(type) {
		case string:
			resultCount, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				resultCount = 0
			}
			results = int(resultCount)
		case int:
		case int32:
		case int64:
			results = int(v)
		}

		offset = offset + results
		totalTradesReturned = totalTradesReturned + results

		fmt.Printf("\nTrades received\n")
		fmt.Printf("Start: %s\n", dateFrom)
		fmt.Printf("  End: %s\n", dateTo)
		fmt.Printf("Count: %d\n", historyInfo.Count)
		fmt.Printf("Offset: %v\n", historyInfo.Offset)
		fmt.Printf("Results: %d\n", historyInfo.Results)

		totalTrades = append(totalTrades, historyInfo.Trades...)
	}

	fmt.Printf("\nAll trades\n")

	for _, trade := range totalTrades {
		fmt.Printf("TradeID: %s, Date: %s\n", trade.OrderID, trade.Date)
	}

	fmt.Printf("\n\nTotal Trades: %d\n\n", totalTradesReturned)
}

func fetchHopperStats(
	id string,
	ch *gogocryptohopper.GoGoCryptohopper) {

	result, err := ch.Hopper.GetStats(id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hopper Stats")
	fmt.Println("DateRange")
	fmt.Println(result.DateRange)
	fmt.Println("Assets")
	fmt.Println(result.Assets)
	fmt.Println("Allocation")
	fmt.Println(result.AllocationOfFunds)
	fmt.Println("Fees")
	fmt.Println(result.Profitstats)

}

func fetchHopperAssets(
	id string,
	ch *gogocryptohopper.GoGoCryptohopper) {

	assets, err := ch.Hopper.GetAssets(id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hopper assets")
	fmt.Println(assets)

}

func fetchHopperLogs(
	id string,
	ch *gogocryptohopper.GoGoCryptohopper) {

	logs, err := ch.Hopper.GetLogs(id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hopper logs")
	for _, log := range logs {
		fmt.Println(log.LogType, log.LogOutput)
	}

}

func hopperPanicSell(id string,
	ch *gogocryptohopper.GoGoCryptohopper) {

	err := ch.Hopper.PanicSell(id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Panic mode triggered")
}
