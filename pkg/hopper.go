package pkg

// Hopper defines the structure of the Hopper
type Hopper struct {
	// ID of the hopper
	ID string `json:"id"`
	// Name of the hopper
	Name string `json:"name"`
	// Exchange used by this hopper
	Exchange string `json:"exchange"`
	// HopperID assigned ID
	HopperID string `json:"hopper_id"`
	// SubscriptionID is the user's subscription ID used by this hopper
	SubscriptionID interface{} `json:"subscription_id"`
	// PaymentTerm is the user's payment term for this subscription used
	// by this hopper
	PaymentTerm string `json:"payment_term"`
	// PaymentMethodID is the payment method used by this hopper
	PaymentMethodID string `json:"payment_method_id"`
	// StartTime is the start time for this subscription
	StartTime string `json:"start_time"`
	// EndTime is the stop time for this subscription
	EndTime string `json:"end_time"`
	// SubscriptionStatus is the status of the user's subscription
	// used by this hopper
	SubscriptionStatus string `json:"subscription_status"`
	// AutoRenewal determines is auto renew is set
	AutoRenewal string `json:"auto_renewal"`
	// Subscription type used by this hopper
	Subscription string `json:"subscription"`
	// PlanName is the name of the subscription
	PlanName string `json:"plan_name"`
	// PlanDescription contains the features for this subscription plan
	PlanDescription string `json:"plan_description"`
	// ProductID Unknown
	ProductID string `json:"product_id"`
	// LastLoadedConfig ?
	LastLoadedConfig interface{} `json:"last_loaded_config"`
	// Image is the image set for this hopper
	Image string `json:"image"`
	// BaseCurrency set for this hopper
	BaseCurrency string `json:"base_currency"`
	// IsBuyingEnabled indicates if the hopper is allowed to buy
	IsBuyingEnabled interface{} `json:"buying_enabled"`
	// IsSellingEnabled indicates if the hopper is allowed to sell
	IsSellingEnabled interface{} `json:"selling_enabled"`
	// Enabled indicates if this hopper is enabled
	Enabled interface{} `json:"enabled"`
	// SetDefault indicates if this hopper is the defaul
	SetDefault string `json:"set_default"`
	// ErrorMessage contains the last error message
	ErrorMessage interface{} `json:"error_message"`
	// LastSignal contains the last received signal
	LastSignal string `json:"last_signal"`
	// LastSignalEncoding specifies the encoding used for the signal information
	LastSignalEncoding string `json:"last_signal_encoding"`
	// TotalCur is the current total assets in base currency in exchange
	CurrentTotal string `json:"total_cur"`
	// ConfigError contains config errors
	ConfigError string `json:"config_error"`
	// Created is the date and time this hopper was created at
	Created string `json:"created"`
	// BotType indicates if this is a trading bot, market maker or arbitrage
	BotType string `json:"bot_type"`
	// UserID is the owned ID of this hopper
	UserID string `json:"user_id"`
	// AllowedCoins is the current coin list used by this hopper
	AllowedCoins []string `json:"allowed_coins"`
	// PaperTrading indicates if this is a paper hopper
	PaperTrading interface{} `json:"paper_trading"`
	// StartBalance holds the value of assets in the base
	// currency when first started
	StartBalance string `json:"start_balance"`
}

// HopperBaseconfig is the structure of the baseconfig for a hopper
// TODO: I'd love to add comments to all these fields, but I actually
// don't know what they all mean
type HopperBaseconfig struct {
	Name                            string      `json:"name"`
	Live                            string      `json:"live"`
	LastLoadedConfig                string      `json:"last_loaded_config"`
	LastLoadedConfigDate            string      `json:"last_loaded_config_date"`
	PaperTradingAccount             string      `json:"paper_trading_account"`
	Exchange                        string      `json:"exchange"`
	APIKey                          string      `json:"api_key"`
	APISecret                       string      `json:"api_secret"`
	APIPassphrase                   string      `json:"api_passphrase"`
	ExtraAPIKey                     string      `json:"extra_api_key"`
	ExtraAPISecret                  string      `json:"extra_api_secret"`
	TickerType                      string      `json:"ticker_type"`
	AllowedCoins                    []string    `json:"allowed_coins"`
	AllowAllCoins                   string      `json:"allow_all_coins"`
	PercBuyAmount                   string      `json:"perc_buy_amount"`
	MinBuyAmount                    string      `json:"min_buy_amount"`
	MaxAmountAllocated              string      `json:"max_amount_allocated"`
	Strategy                        string      `json:"strategy"`
	NumTargetsPerBuy                string      `json:"num_targets_per_buy"`
	AdvancedTaCandleSize            interface{} `json:"advanced_ta_candle_size"`
	AdvancedTaStochFastK            interface{} `json:"advanced_ta_stoch_fast_k"`
	AdvancedTaStochSlowK            interface{} `json:"advanced_ta_stoch_slow_k"`
	AdvancedTaStochSlowKMatype      interface{} `json:"advanced_ta_stoch_slow_k_matype"`
	AdvancedTaStochSlowD            interface{} `json:"advanced_ta_stoch_slow_d"`
	AdvancedTaStochSlowDMatype      interface{} `json:"advanced_ta_stoch_slow_d_matype"`
	AdvancedTaStochOversold         interface{} `json:"advanced_ta_stoch_oversold"`
	AdvancedTaStochOverbought       interface{} `json:"advanced_ta_stoch_overbought"`
	AdvancedTaStochrsiPeriod        interface{} `json:"advanced_ta_stochrsi_period"`
	AdvancedTaStochrsiFastK         interface{} `json:"advanced_ta_stochrsi_fast_k"`
	AdvancedTaStochrsiFastD         interface{} `json:"advanced_ta_stochrsi_fast_d"`
	AdvancedTaStochrsiFastDMatype   interface{} `json:"advanced_ta_stochrsi_fast_d_matype"`
	AdvancedTaStochrsiOversold      interface{} `json:"advanced_ta_stochrsi_oversold"`
	AdvancedTaStochrsiOverbought    interface{} `json:"advanced_ta_stochrsi_overbought"`
	AdvancedTaRsiPeriod             interface{} `json:"advanced_ta_rsi_period"`
	AdvancedTaRsiOversold           interface{} `json:"advanced_ta_rsi_oversold"`
	AdvancedTaRsiOverbought         interface{} `json:"advanced_ta_rsi_overbought"`
	AdvancedTaMacdFastPeriod        interface{} `json:"advanced_ta_macd_fast_period"`
	AdvancedTaMacdSlowPeriod        interface{} `json:"advanced_ta_macd_slow_period"`
	AdvancedTaMacdSignalPeriod      interface{} `json:"advanced_ta_macd_signal_period"`
	AdvancedTaEmaShortPeriod        interface{} `json:"advanced_ta_ema_short_period"`
	AdvancedTaEmaLongPeriod         interface{} `json:"advanced_ta_ema_long_period"`
	AdvancedTaSmaShortPeriod        interface{} `json:"advanced_ta_sma_short_period"`
	AdvancedTaSmaLongPeriod         interface{} `json:"advanced_ta_sma_long_period"`
	AdvancedTaKamaShortPeriod       interface{} `json:"advanced_ta_kama_short_period"`
	AdvancedTaKamaLongPeriod        interface{} `json:"advanced_ta_kama_long_period"`
	AdvancedTaBbandsPeriod          interface{} `json:"advanced_ta_bbands_period"`
	AdvancedTaBbandsDevUp           interface{} `json:"advanced_ta_bbands_dev_up"`
	AdvancedTaBbandsDevLow          interface{} `json:"advanced_ta_bbands_dev_low"`
	AdvancedTaBbandsMatype          interface{} `json:"advanced_ta_bbands_matype"`
	BuyScoreCorrected               interface{} `json:"buy_score_corrected"`
	MinBuyScore                     interface{} `json:"min_buy_score"`
	SellScoreCorrected              interface{} `json:"sell_score_corrected"`
	MinSellScore                    interface{} `json:"min_sell_score"`
	RsiCandleSize                   interface{} `json:"rsi_candle_size"`
	RsiPeriod                       interface{} `json:"rsi_period"`
	RsiOversold                     interface{} `json:"rsi_oversold"`
	BbandsCandleSize                interface{} `json:"bbands_candle_size"`
	BbandsDays                      interface{} `json:"bbands_days"`
	BbandsDeviation                 interface{} `json:"bbands_deviation"`
	BuyOrderType                    interface{} `json:"buy_order_type"`
	BidPercentage                   interface{} `json:"bid_percentage"`
	BidPercentageType               interface{} `json:"bid_percentage_type"`
	MaxOpenTimeBuy                  interface{} `json:"max_open_time_buy"`
	MaxOpenPositions                string      `json:"max_open_positions"`
	MaxOpenPositionsPerCoin         string      `json:"max_open_positions_per_coin"`
	Cooldown                        interface{} `json:"cooldown"`
	CooldownWhen                    interface{} `json:"cooldown_when"`
	CooldownCount                   interface{} `json:"cooldown_count"`
	CooldownVal                     interface{} `json:"cooldown_val"`
	OneOpenOrder                    interface{} `json:"one_open_order"`
	OnlyWhenPositiveTime            interface{} `json:"only_when_positive_time"`
	Arbitrage                       interface{} `json:"arbitrage"`
	StopLossPercentage              interface{} `json:"stop_loss_percentage"`
	StopLossTrailing                interface{} `json:"stop_loss_trailing"`
	StopLossTrailingPercentage      interface{} `json:"stop_loss_trailing_percentage"`
	StopLossTrailingArm             interface{} `json:"stop_loss_trailing_arm"`
	StopLossTrailingOnly            interface{} `json:"stop_loss_trailing_only"`
	TrailingStopLossProfit          interface{} `json:"trailing_stop_loss_profit"`
	TrailingBuyPercentage           interface{} `json:"trailing_buy_percentage"`
	MaxOpenShortPositions           interface{} `json:"max_open_short_positions"`
	ShortPercentageProfit           interface{} `json:"short_percentage_profit"`
	ShortStopLossTrailingPercentage interface{} `json:"short_stop_loss_trailing_percentage"`
	ShortStopLossTrailingArm        interface{} `json:"short_stop_loss_trailing_arm"`
	ShortAutoClosePositionsTime     interface{} `json:"short_auto_close_positions_time"`
	ShortAutoRemovePositionsTime    interface{} `json:"short_auto_remove_positions_time"`
	ShortRemoveOnLoss               interface{} `json:"short_remove_on_loss"`
	HoldAssets                      interface{} `json:"hold_assets"`
	AutoClosePositionsTime          interface{} `json:"auto_close_positions_time"`
	DcaOrderType                    string      `json:"dca_order_type"`
	AutoDcaTime                     string      `json:"auto_dca_time"`
	AutoDcaMax                      string      `json:"auto_dca_max"`
	AutoDcaPercentage               string      `json:"auto_dca_percentage"`
	AutoDcaSize                     string      `json:"auto_dca_size"`
	AutoDcaSizeCustom               string      `json:"auto_dca_size_custom"`
	SetPercentage                   interface{} `json:"set_percentage"`
	SellOrderType                   interface{} `json:"sell_order_type"`
	MaxOpenTime                     interface{} `json:"max_open_time"`
	AskPercentage                   interface{} `json:"ask_percentage"`
	AskPercentageType               interface{} `json:"ask_percentage_type"`
	Submit                          interface{} `json:"submit"`
	StopBuying                      interface{} `json:"stop_buying"`
	StopSelling                     interface{} `json:"stop_selling"`
	WalletscrubberLeftovers         interface{} `json:"walletscrubber_leftovers"`
	WalletscrubberAuto              interface{} `json:"walletscrubber_auto"`
	WalletscrubberNoopen            interface{} `json:"walletscrubber_noopen"`
	CollectCurrency                 string      `json:"collect_currency"`
	OutputLiveFeed                  string      `json:"output_live_feed"`
	OutputErrorsOnly                string      `json:"output_errors_only"`
	Autosync                        interface{} `json:"autosync"`
	AutosyncAllCoins                interface{} `json:"autosync_all_coins"`
	AutosyncAllowedCoins            interface{} `json:"autosync_allowed_coins"`
	SendTradeEmail                  interface{} `json:"send_trade_email"`
	SendTradeErrorEmail             interface{} `json:"send_trade_error_email"`
	SendCancelledEmail              interface{} `json:"send_cancelled_email"`
	BuyBtcAmount                    string      `json:"buy_btc_amount"`
}

// GetAllHoppersResponse holds an array of HopperResponses
type GetAllHoppersResponse struct {
	// Data contains the response information
	Data struct {
		// Hoppers contain the list of hoppers
		Hoppers []Hopper `json:"hoppers"`
	} `json:"data"`
}

// GetHopperResponse holds the HopperResponse
type GetHopperResponse struct {
	// Data contains the response information
	Data struct {
		// Hopper contain the hopper information
		Hopper Hopper `json:"hopper"`
	} `json:"data"`
}

// GetHopperConfigResponse holds the baseconfig response
type GetHopperConfigResponse struct {
	// Data contains the response information
	Data HopperBaseconfig `json:"data"`
}

type GetHopperOpenPositionsResponse struct {
	Data []Position `json:"data"`
}

type GetHopperReservedPositionsResponse struct {
	Data []Position `json:"data"`
}

type GetHopperShortPositionsResponse struct {
	Data []Position `json:"data"`
}

type TaValues struct {
	Strategy     interface{} `json:"strategy"`
	Coin         interface{} `json:"coin"`
	StrategyType interface{} `json:"strategy_type"`
	StrategyID   interface{} `json:"strategy_id"`
	Values       interface{} `json:"values"`
	Signals      interface{} `json:"signals"`
}
type Position struct {
	ID                    string      `json:"id"`
	Pair                  string      `json:"pair"`
	Currency              string      `json:"currency"`
	Amount                string      `json:"amount"`
	Rate                  string      `json:"rate"`
	BuyID                 string      `json:"buy_id"`
	OrderID               string      `json:"order_id"`
	SellID                string      `json:"sell_id"`
	BuyTrigger            string      `json:"buy_trigger"`
	SellTrigger           string      `json:"sell_trigger"`
	TaValues              interface{} `json:"ta_values"`
	PercentageProfit      string      `json:"percentage_profit"`
	StopLoss              string      `json:"stop_loss"`
	StopLossPercentage    string      `json:"stop_loss_percentage"`
	SellRate              string      `json:"sell_rate"`
	TrailingPercentage    string      `json:"trailing_percentage"`
	TrailingArm           string      `json:"trailing_arm"`
	TrailingArmPercentage string      `json:"trailing_arm_percentage"`
	BuyTime               string      `json:"buy_time"`
	SellTime              string      `json:"sell_time"`
	AutoCloseTime         string      `json:"auto_close_time"`
	Hold                  string      `json:"hold"`
	Sold                  string      `json:"sold"`
	Cost                  string      `json:"cost"`
	ReservedCost          string      `json:"val"`
	ReservedBaseCurrency  string      `json:"base"`
	ShortStartProfit      string      `json:"start_profit"`
	ShortStartRate        string      `json:"start_rate"`
}

type GetHopperTradesResponse struct {
	Data TradeList `json:"data"`
}

type TradeList struct {
	Trades  []Trade     `json:"trades"`
	Offset  string      `json:"offset"`
	Count   interface{} `json:"count"`
	Results interface{} `json:"results"`
}

type StrategyResult struct {
	Strategy     string `json:"strategy"`
	Coin         string `json:"coin"`
	StrategyType string `json:"strategy_type"`
	StrategyID   string `json:"strategy_id"`
	Values       string `json:"values"`
	Signals      string `json:"signals"`
}

type Trade struct {
	ID              string      `json:"id"`
	Exchange        string      `json:"exchange"`
	Currency        string      `json:"currency"`
	Pair            string      `json:"pair"`
	Type            string      `json:"type"`
	OrderCur        string      `json:"order_cur"`
	OrderID         string      `json:"order_id"`
	Date            string      `json:"date"`
	Amount          string      `json:"amount"`
	Rate            string      `json:"rate"`
	Total           string      `json:"total"`
	Fee             string      `json:"fee"`
	Result          string      `json:"result"`
	TriggerStrategy string      `json:"trigger_strategy"`
	StrategyResult  interface{} `json:"strategy_result"`
	BuyID           string      `json:"buy_id"`
	IsShort         string      `json:"is_short"`
	ResultShort     string      `json:"result_short"`
}

type HopperStats struct {
	DateRange                 interface{} `json:"date_range"`
	StartTime                 interface{} `json:"start_time"`
	EndTime                   interface{} `json:"end_time"`
	Buysandsells              interface{} `json:"buysandsells"`
	Mosttraded                interface{} `json:"mosttraded"`
	AverageHoldingTime        interface{} `json:"average_holding_time"`
	AverageProfit             interface{} `json:"average_profit"`
	ProfitsSellTriggers       interface{} `json:"profits_sell_triggers"`
	ProfitsSellTriggersTotal  interface{} `json:"profits_sell_triggers_total"`
	ProfitsSellTriggersCounts interface{} `json:"profits_sell_triggers_counts"`
	ProfitsBuyTriggers        interface{} `json:"profits_buy_triggers"`
	ProfitsBuyTriggersTotal   interface{} `json:"profits_buy_triggers_total"`
	ProfitsBuyTriggersCounts  interface{} `json:"profits_buy_triggers_counts"`
	DailyProfits              interface{} `json:"daily_profits"`
	AllocationOfFunds         struct {
		Available       interface{} `json:"available"`
		InOpenPositions interface{} `json:"in open positions"`
		Reserved        interface{} `json:"reserved"`
	} `json:"allocation_of_funds"`
	Currentinassets         [][]interface{} `json:"currentinassets"`
	Currentusdinassets      [][]interface{} `json:"currentusdinassets"`
	AverageProfitTotal      interface{}     `json:"average_profit_total"`
	AverageHoldingTimeTotal interface{}     `json:"average_holding_time_total"`
	Totaltraded             interface{}     `json:"totaltraded"`
	Returns                 interface{}     `json:"returns"`
	TotalReturns            interface{}     `json:"total_returns"`
	Invested                interface{}     `json:"invested"`
	TotalInvested           interface{}     `json:"total_invested"`
	Positions               interface{}     `json:"positions"`
	Profit                  interface{}     `json:"profit"`
	Invest                  interface{}     `json:"invest"`
	Fees                    interface{}     `json:"fees"`
	Profitstats             interface{}     `json:"profitstats"`
	StatsOpenPositions      interface{}     `json:"stats_open_positions"`
	Assets                  interface{}     `json:"assets"`
}

type GetHopperStatsResponse struct {
	Data HopperStats `json:"data"`
}

type GetHopperAssetsResponse struct {
	Data map[string]string `json:"data"`
}

type HopperLog struct {
	LogType   string      `json:"log_type"`
	LogOutput string      `json:"log_output"`
	LogTime   interface{} `json:"log_time"`
}

type GetHopperLogsResponse struct {
	Data []HopperLog `json:"data"`
}

type GetHopperPanicSellResponse struct {
	Data string `json:"data"`
}
