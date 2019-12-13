package main

type SiteFaqData struct {
	SelectedID string
}

func (data *SiteFaqData) HandleClick(id string) {
	if data.SelectedID == id {
		data.SelectedID = "" // toggle
		return
	}
	data.SelectedID = id
}
