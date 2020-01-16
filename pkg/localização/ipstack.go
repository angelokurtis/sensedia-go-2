package localização

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/angelokurtis/sensedia-go-2/internal/log"
	"github.com/pkg/errors"
)

type IPStack struct{}

func NewIPStack() *IPStack {
	return &IPStack{}
}

func (i *IPStack) OndeEstou() (*Local, error) {
	log.Info("buscando minha localização atual")

	ak := os.Getenv("IPSTACK_API_ACCESS_KEY")
	r, err := http.Get("http://api.ipstack.com/check?access_key=" + ak)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	b := r.Body
	defer b.Close()
	l := &Local{}
	err = json.NewDecoder(b).Decode(l)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if l.Cidade == "" {
		return nil, errors.New("localização não encontrada")
	}
	log.Info("localização encontrada")
	log.Infof("estamos em %s/%s", l.Cidade, l.CódigoRegional)
	return l, nil
}
