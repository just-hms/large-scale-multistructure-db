package usecase

type TokenUsecase struct {
	api TokenApi
}

func NewTokenUsecase(a TokenApi) *TokenUsecase {
	return &TokenUsecase{
		api: a,
	}
}

func (uc *TokenUsecase) CreateToken(userID string) (string, error) {
	return uc.api.CreateToken(userID)
}

func (uc *TokenUsecase) ExtractTokenID(tokenString string) (string, error) {
	return uc.api.ExtractTokenID(tokenString)
}
