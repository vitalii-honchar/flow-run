package provider

type GetProviderHandler struct {
}

func (g *GetProviderHandler) Group() string {
	return groupProviderV1
}
