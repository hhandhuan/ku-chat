package ws

type Msg struct {
	ID   uint32 `json:"id"`
	Data string `json:"data"`
}

func (m *Msg) GetData() string {
	return m.Data
}

func (m *Msg) GetID() uint32 {
	return m.ID
}
