package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const ApiUrl = "https://wow.kostya.sexy"

type (
	RaidPlayer struct {
		Name string `json:"name"`
		Spec string `json:"spec"`
		Alt  bool   `json:"alt"`
	}

	Party struct {
		Players [5]RaidPlayer `json:"players"`
	}

	RaidParty struct {
		Id             int      `json:"id"`
		Timing         string   `json:"timing"`
		RaidLeaderName string   `json:"raidLeaderName"`
		Parties        [8]Party `json:"parties"`
	}

	LootPriority struct {
		ItemID   int    `json:"itemID"`
		Priority string `json:"priority"`
	}

	News struct {
		Author string `json:"author"`
		Text   string `json:"text"`
		Time   int    `json:"time"`
	}

	Miniwarlock struct {
		OriginalName string `json:"originalName"`
		WarlockName  string `json:"warlockName"`
		Ready        bool   `json:"ready"`
	}

	WarlocksZone struct {
		Location string        `json:"location"`
		Warlocks []Miniwarlock `json:"warlocks"`
	}
)

func getRaidParties() (*[]RaidParty, error) {
	var raids []RaidParty
	err := callAPI(ApiUrl+"/raids", func(reader io.ReadCloser) error {
		return json.NewDecoder(reader).Decode(&raids)
	})
	if err != nil {
		return nil, err
	}
	return &raids, nil
}

func getLootPriorities() (*[]LootPriority, error) {
	var priorities []LootPriority
	err := callAPI(ApiUrl+"/loot", func(reader io.ReadCloser) error {
		return json.NewDecoder(reader).Decode(&priorities)
	})
	if err != nil {
		return nil, err
	}
	return &priorities, nil
}

func getWarlocks() (*[]WarlocksZone, error) {
	var zones []WarlocksZone
	err := callAPI(ApiUrl+"/warlocks", func(reader io.ReadCloser) error {
		return json.NewDecoder(reader).Decode(&zones)
	})
	if err != nil {
		return nil, err
	}
	return &zones, nil
}

func getNews() (*[]News, error) {
	var news []News
	err := callAPI(ApiUrl+"/news", func(reader io.ReadCloser) error {
		return json.NewDecoder(reader).Decode(&news)
	})
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func callAPI(url string, consumer func(reader io.ReadCloser) error) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return consumer(resp.Body)
}
