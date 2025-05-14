package adapters

type LudopediaCollectionItem struct {
	Id           int     `json:"id_jogo"`
	Name         string  `json:"nm_jogo"`
	ImageUrl     *string `json:"thumb"`
	LudopediaUrl *string `json:"link"`
	Paid         *int    `json:"vl_custo"`
}

type GetCollectionOutput struct {
	Collection []*LudopediaCollectionItem `json:"colecao"`
}

type GetCollectionInput struct {
	AccessToken string
	Page        string
}

type GetGameOutput struct {
	Id                 int    `json:"id_jogo"`
	Name               string `json:"nm_jogo"`
	ImageUrl           string `json:"thumb"`
	LudopediaUrl       string `json:"link"`
	MinAmountOfPlayers int    `json:"qt_jogadores_min"`
	MaxAmountOfPlayers int    `json:"qt_jogadores_max"`
	AverageDuration    int    `json:"vl_tempo_jogo"`
	MinAge             int    `json:"idade_minima"`
}

type GetGameInput struct {
	AccessToken string
	LudopediaId int
}

type Ludopedia interface {
	GetCollection(i *GetCollectionInput) (*GetCollectionOutput, error)
	GetGame(i *GetGameInput) (*GetGameOutput, error)
}
