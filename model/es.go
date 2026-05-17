package model

type SetEsEssayResponse struct {
	Result string `json:"result"`
}

type GetEsEssayResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			ID     string `json:"_id"`
			Source Essay  `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
