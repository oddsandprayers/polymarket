package condition

type Interface interface {
	// Search is to fetch condition information including all of its associated
	// outcomes. Those onchain condition outcomes map to offchain market outcomes
	// within Polymarket. We need to resolve market conditions in order to answer
	// one question.
	//
	//    How did a market resolve?
	//
	Search(Option) ([]Object, error)
}
