package kangaroo

type utils interface {
	IsFloatInt(floatValue float64) bool
}

// NewKangarooFabric return fabric for kangaroo
func NewKangarooFabric(utils utils) func(x, v int) *Kangaroo {
	return func(X, V int) *Kangaroo {
		return &Kangaroo{
			utils: utils,
			x:     X,
			v:     V,
		}
	}
}
