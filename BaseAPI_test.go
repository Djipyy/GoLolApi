package gololapi_test

import (
	"testing"

	gololapi "romaghi.net/golol-api"
)

func TestMinifySummonerName(t *testing.T) {
	input_1 := "Beng Waltsunny"
	expected_1 := "bengwaltsunny"
	output_1 := gololapi.MinifySummonerName(input_1)
	if output_1 != expected_1 {
		t.Errorf("Output was incorrect, got %s, wanted %s", output_1, expected_1)
	}
}

func TestGetEndpointURI(t *testing.T) {
	endpoint := "/lol/static-data/v3/summoner-spells"
	expected_1 := "/lol/static-data/v3/summoner-spells?locale=fr_FR"
	expected_2 := "/lol/static-data/v3/summoner-spells?locale=fr_FR&spellData=all"
	expected_3 := "/lol/static-data/v3/summoner-spells"
	options_1 := map[string]string{"locale": "fr_FR"}
	options_2 := map[string]string{"locale": "fr_FR", "spellData": "all"}
	options_3 := map[string]string{}
	output_1, hasParameters_1 := gololapi.GetEndpointURI(endpoint, options_1)
	if output_1 != expected_1 || !hasParameters_1 {
		t.Errorf("Output was incorrect, got %s, wanted %s", output_1, expected_1)
	}
	output_2, hasParameters_2 := gololapi.GetEndpointURI(endpoint, options_2)
	if output_2 != expected_2 || !hasParameters_2 {
		t.Errorf("Output was incorrect, got %s, wanted %s", output_2, expected_2)
	}
	output_3, hasParameters_3 := gololapi.GetEndpointURI(endpoint, options_3)
	if output_3 != expected_3 || hasParameters_3 {
		t.Errorf("Output was incorrect, got %s, wanted %s", output_3, expected_3)
	}
}
