package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/PirosB3/coinswap/incubator/coinswap/internal/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/coinswap/add_liquidity",
		postAddLiquidityHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/coinswap/remove_liquidity",
		postRemoveLiquidityHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/coinswap/swap_order",
		postSwapOrderHandlerFn(cliCtx),
	).Methods("POST")
}

type (
	// AddLiquidityRequest defines the properties of an add liquidity request's body.
	AddLiquidityRequest struct {
		BaseReq       rest.BaseReq `json:"base_req"`
		Deposit       sdk.Coin     `json:"deposit"`
		DepositAmount sdk.Int      `json:"deposit_amount"`
		MinReward     sdk.Int      `json:"min_reward"`
		Deadline      time.Time    `json:"deadline"`
	}

	// RemoveLiquidityRequest defines the properties of a remove liquidity request's body.
	RemoveLiquidityRequest struct {
		BaseReq        rest.BaseReq `json:"base_req"`
		Withdraw       sdk.Coin     `json:"withdraw"`
		WithdrawAmount sdk.Int      `json:"withdraw_amount"`
		MinNative      sdk.Int      `json:"min_native"`
		Deadline       time.Time    `json:"deadline"`
	}

	// SwapOrderRequest defines the properties of a swap order request's body.
	SwapOrderRequest struct {
		BaseReq    rest.BaseReq   `json:"base_req"`
		Input      sdk.Coin       `json:"input"`
		Output     sdk.Coin       `json:"output"`
		Deadline   time.Time      `json:"deadline"`
		Recipient  sdk.AccAddress `json:"recipient"` // in bech32
		IsBuyOrder bool           `json:"is_buy_order"`
	}
)

func postAddLiquidityHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddLiquidityRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		senderAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgAddLiquidity(req.Deposit, req.DepositAmount, req.MinReward, req.Deadline, senderAddr)
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postRemoveLiquidityHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RemoveLiquidityRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		senderAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgRemoveLiquidity(req.Withdraw, req.WithdrawAmount, req.MinNative, req.Deadline, senderAddr)
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postSwapOrderHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SwapOrderRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		senderAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgSwapOrder(req.Input, req.Output, req.Deadline, senderAddr, req.Recipient, req.IsBuyOrder)
		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
