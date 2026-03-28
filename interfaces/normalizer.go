package interfaces

type Normalizer interface {
	NormalizeInputs()
}

func NormalizeInput(n Normalizer) {
	n.NormalizeInputs()
}
