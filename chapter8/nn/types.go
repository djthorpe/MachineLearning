package nn

type Config struct {
	InputNeurons  uint
	OutputNeurons uint
	HiddenNeurons uint
	NumEpochs     uint
	LearningRate  float64
}

type net struct {
}

func NewThreeLayer(config Config) *net {
	this := new(net)

}
