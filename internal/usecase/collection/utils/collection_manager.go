package collection_utils

type GameToImport struct {
	LudopediaGameId int
	Paid            *int
}

type AddAccountLudopediaGameInput struct {
	AccountId       int
	AccessToken     string
	LudopediaGameId int
	Paid            *int
}

type CollectionManager struct {
	ludopediaGamesIds map[int]bool

	AccountLudopediaGamesMap     map[int][]*GameToImport
	AccessTokenByLudopediaGameId map[int]string // [LudopediaId]: AccessToken
}

func (self *CollectionManager) AddAccountLudopediaGame(i *AddAccountLudopediaGameInput) {
	if self.AccountLudopediaGamesMap[i.AccountId] == nil {
		self.AccountLudopediaGamesMap[i.AccountId] = []*GameToImport{}
	}

	// Sets the access token of the first user that has the game here,
	// so we can get the game information if needed
	// Yes, unfortunately the Ludopedia API requires a user access token to do it
	if self.AccessTokenByLudopediaGameId[i.LudopediaGameId] == "" {
		self.AccessTokenByLudopediaGameId[i.LudopediaGameId] = i.AccessToken
	}

	self.AccountLudopediaGamesMap[i.AccountId] = append(
		self.AccountLudopediaGamesMap[i.AccountId],
		&GameToImport{
			LudopediaGameId: i.LudopediaGameId,
			Paid:            i.Paid,
		},
	)
	self.ludopediaGamesIds[i.LudopediaGameId] = true
}

func (self *CollectionManager) GetLudopediaGamesIds() []int {
	ludopediaGamesIds := make([]int, 0, len(self.ludopediaGamesIds))
	for v := range self.ludopediaGamesIds {
		ludopediaGamesIds = append(ludopediaGamesIds, v)
	}
	return ludopediaGamesIds
}

func NewCollectionManager() *CollectionManager {
	return &CollectionManager{
		ludopediaGamesIds:            map[int]bool{},
		AccountLudopediaGamesMap:     map[int][]*GameToImport{},
		AccessTokenByLudopediaGameId: map[int]string{},
	}
}
