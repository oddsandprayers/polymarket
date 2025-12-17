package token

type Interface interface {
	// Search is to fetch price information for any given token ID. With that we
	// can figure out how a market traded over time.
	Search(Option) (Object, error)
}
