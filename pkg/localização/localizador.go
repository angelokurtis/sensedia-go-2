package localização

type Localizador interface {
	OndeEstou() (*Local, error)
}
