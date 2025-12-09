package client

const (
	// PolymarketApiUrl is the endpoint format of the underlying Polymarket APIs
	// being used. Also see the documentation page for this default option.
	//
	//     https://docs.polymarket.com/quickstart/introduction/endpoints
	//
	PolymarketApiUrl url = "https://%s.polymarket.com/%s"

	// SubgraphApiUrl is the endpoint of the underlying subgraph API being used.
	// Also see the explorer page for this default option.
	//
	//     https://thegraph.com/explorer/subgraphs/81Dm16JjuFSrqz813HysXoUPvzTwE7fsfPk2RTf66nyC
	//
	SubgraphApiUrl url = "https://gateway.thegraph.com/api/subgraphs/id/81Dm16JjuFSrqz813HysXoUPvzTwE7fsfPk2RTf66nyC"
)

type url string
