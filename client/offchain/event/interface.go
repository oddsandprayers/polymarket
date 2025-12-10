package event

type Interface interface {
	// Search is to fetch event information including all of its associated
	// markets. One event has potentially many markets within Polymarket. E.g. a
	// single soccer game may have three markets in order to answer the following
	// questions.
	//
	//    1. Will team A win or lose?
	//    2. Will team B win or lose?
	//    3. Will the game end in a draw?
	//
	Search(Option) ([]Object, error)
}
