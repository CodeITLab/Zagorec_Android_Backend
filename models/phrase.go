package models

type phrase struct {
	id int `json:"id,omitempty"`
	phrase string `json:"phrase,omitempty"`
	explanation string `jsong:"phrase,omitempty"`
}

func NewPhrase() phrase {
	return phrase{
		id:          0,
		phrase:      "",
		explanation: "",
	}
}

func (r *phrase) Save(db interface{}) error {
	return nil
}

func (r *phrase) Get(db interface{}) error {
	return nil
}

func (r *phrase) Update(db interface{}) error {
	return nil
}

func (r *phrase) Delete(db interface{}) error {
	return nil
}