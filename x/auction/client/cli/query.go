package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/spf13/cobra"

	"github.com/tharsis/ethermint/x/auction/types"
)

func GetQueryCmd() *cobra.Command {
	auctionQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	auctionQueryCmd.AddCommand(
		GetCmdList(),
		GetCmdGetAuction(),
		GetCmdGetBid(),
		GetCmdGetBids(),
		GetCmdListByBidder(),
		GetCmdQueryParams(),
		GetCmdBalance(),
	)

	return auctionQueryCmd
}

// GetCmdList queries all auctions.
func GetCmdList() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List auctions.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.List(cmd.Context(), &types.ListRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdGetBid queries an auction bid.
func GetCmdGetBid() *cobra.Command {
	return &cobra.Command{
		Use:   "get-bid [auction-id] [bidder]",
		Short: "Get auction bid.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id := args[0]
			bidder := args[1]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetBid(cmd.Context(), &types.BidRequest{AuctionId: id, Bidder: bidder})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdGetBids queries all auction bids.
func GetCmdGetBids() *cobra.Command {
	return &cobra.Command{
		Use:   "get-bids [auction-id]",
		Short: "Get all auction bids.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetBids(cmd.Context(), &types.BidsRequest{AuctionId: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdGetAuction queries an auction.
func GetCmdGetAuction() *cobra.Command {
	return &cobra.Command{
		Use:   "get [ID]",
		Short: "Get auction.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetAuction(cmd.Context(), &types.AuctionRequest{Id: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdListByBidder queries auctions by bidder.
func GetCmdListByBidder() *cobra.Command {
	return &cobra.Command{
		Use:   "query-by-owner [address]",
		Short: "Query auctions by owner/creator.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			address := args[0]

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ListByBidder(cmd.Context(), &types.ListByBidderRequest{BidderAddress: address})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current auction parameters information.",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as auction parameters.
				Example:
				$ %s query auction params
				`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.QueryParams(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdBalance queries the auction module account balance.
func GetCmdBalance() *cobra.Command {
	return &cobra.Command{
		Use:   "balance",
		Short: "Get auction module account balance.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Balance(cmd.Context(), &types.BalanceRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}
