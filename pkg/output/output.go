package output

import (
	"github.com/infosec-cyber/ffuf/v2/pkg/ffuf"
)

func NewOutputProviderByName(name string, conf *ffuf.Config) ffuf.OutputProvider {
	//We have only one outputprovider at the moment
	return NewStdoutput(conf)
}
