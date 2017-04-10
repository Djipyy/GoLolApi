package gololapi

import (
	"encoding/json"

	"strconv"

	cache "github.com/patrickmn/go-cache"
)

type StaticChampionList struct {
	Data    map[string]StaticChampion
	Keys    map[string]string
	Version string
	Type    string
	Format  string
}
type StaticChampion struct {
	//Info      InfoDto
	EnemyTips   []string
	Stats       StatsDto
	Name        string
	Title       string
	Image       ImageDto
	Tags        []string
	Partype     string
	Skins       []SkinDto
	Passive     PassiveDto
	Recommended []RecommendedDto
	AllyTips    []string
	Key         string
	Lore        string
	ID          int
	Blurb       string
	Spells      []ChampionSpellDto
}
type StatsDto struct {
	Armorperlevel        float32
	Hpperlevel           float32
	Attackdamage         float32
	Mpperlevel           float32
	Attackspeedoffset    float32
	Armor                float32
	Hp                   float32
	Hpregenperlevel      float32
	Spellblock           float32
	Attackrange          float32
	Movespeed            float32
	Attackdamageperlevel float32
	Mpregenperlevel      float32
	Mp                   float32
	Spellblockperlevel   float32
	Crit                 float32
	Mpregen              float32
	Attackspeedperlevel  float32
	Hpregen              float32
	Critperlevel         float32
}
type ImageDto struct {
	Full       string
	Group      string
	Sprite     string
	H, W, Y, X int
}
type SkinDto struct {
	Num  int
	Name string
	ID   int
}
type ChampionSpellDto struct {
	CooldownBurn         string
	Ressource            string
	LevelTip             LevelTipDto
	Vars                 []SpellVarsDto
	CostType             string
	Image                ImageDto
	SanitizedDescription string
	SanitizedTooltip     string
	Effect               [][]float32
	Tooltip              string
	MaxRank              int
	CostBurn             string
	RangeBurn            string
	Range                interface{}
	Cooldown             []float32
	Cost                 []int
	Key                  string
	Description          string
	EffectBurn           []string
	AltImages            []ImageDto
	Name                 string
}
type LevelTipDto struct {
	Effect []string
	Label  []string
}

type SpellVarsDto struct {
	RanksWith string
	Dyn       string
	Link      string
	Coeff     []float32
	Key       string
}

type RecommendedDto struct {
	Map      string
	Champion string
	Title    string
	Priority bool
	Mode     string
	Type     string
	blocks   []BlockDto
}
type BlockDto struct {
	Items   []BlockItemDto
	RecMath bool
	Type    string
}
type BlockItemDto struct {
	Count int
	ID    int
}
type PassiveDto struct {
	Image                ImageDto
	Description          string
	SanitizedDescription string
	Name                 string
}

func (api *GoLOLAPI) StaticDataGetChampions(version int, locale string, complete bool) (list StaticChampionList) {
	options := map[string]string{}
	if version != 0 {
		options["version"] = strconv.Itoa(version)
	}
	if locale != "" {
		options["locale"] = locale
	}
	if complete {
		options["champData"] = "all"
	}
	uri, hasParameters := getEndpointURI("/lol/static-data/v3/champions", options)
	response, e := api.RequestStaticData(uri, cache.NoExpiration, hasParameters)
	if e != nil {
		panic(e)
	}
	list = StaticChampionList{}
	err := json.Unmarshal(response, &list)
	if err != nil {
		panic(err)
	}
	return
}

type ItemListDto struct {
	Data    map[string]ItemDto
	Tree    []ItemTreeDto
	Version string
	Groups  []GroupDto
	Basic   BasicDataDto
	Type    string
}
type ItemDto struct {
	Gold                 GoldDto
	PlainText            string
	HideFromAll          bool
	InStore              bool
	Into                 []string
	ID                   int
	Stats                BasicDataStatsDto
	Colloq               string
	Maps                 map[string]bool
	SpecialRecipe        int
	Image                ImageDto
	Description          string
	From                 []string
	Group                string
	ConsumeOnFull        bool
	Name                 string
	Consumed             bool
	SanitizedDescription string
	Depth                int
	Rune                 MetaDataDto
	Stacks               int
}
type GoldDto struct {
	Sell        int
	Total       int
	Base        int
	Purchasable bool
}
type BasicDataStatsDto struct {
	RPercentMagicPenetrationModPerLevel float64
	RFlatEnergyModPerLevel              float64
	PercentMagicDamageMod               float64
	PercentCritDamageMod                float64
	PercentSpellBlockMod                float64
	PercentHPRegenMod                   float64
	PercentMovementSpeedMod             float64
	FlatSpellBlockMod                   float64
	PercentHPPoolMod                    float64
	FlatEnergyPoolMod                   float64
	RFlatDodgeMod                       float64
	PercentLifeStealMod                 float64
	RFlatMPRegenModPerLevel             float64
	FlatMPPoolMod                       float64
	RFlatGoldPer10Mod                   float64
	FlatMovementSpeedMod                float64
	RPercentCooldownMod                 float64
	RFlatMPModPerLevel                  float64
	RPercentCooldownModPerLevel         float64
	PercentAttackSpeedMod               float64
	RPercentMagicPenetrationMod         float64
	PercentBlockMod                     float64
	RFlatTimeDeadModPerLevel            float64
	FlatEnergyRegenMod                  float64
	RPercentAttackSpeedModPerLevel      float64
	PercentSpellVampMod                 float64
	FlatMPRegenMod                      float64
	RFlatTimeDeadMod                    float64
	RFlatMagicDamageModPerLevel         float64
	FlatAttackSpeedMod                  float64
	RFlatMagicPenetrationMod            float64
	RFlatCritChanceModPerLevel          float64
	PercentMPRegenMod                   float64
	PercentDodgeMod                     float64
	RFlatHPModPerLevel                  float64
	PercentPhysicalDamageMod            float64
	RFlatDodgeModPerLevel               float64
	RPercentMovementSpeedModPerLevel    float64
	RFlatSpellBlockModPerLevel          float64
	FlatBlockMod                        float64
	PercentMPPoolMod                    float64
	FlatMagicDamageMod                  float64
	RFlatMagicPenetrationModPerLevel    float64
	FlatHPRegenMod                      float64
	RFlatPhysicalDamageModPerLevel      float64
	RFlatEnergyRegenModPerLevel         float64
	FlatPhysicalDamageMod               float64
	RPercentTimeDeadMod                 float64
	FlatCritDamageMod                   float64
	RFlatArmorPenetrationMod            float64
	PercentArmorMod                     float64
	PercentCritChanceMod                float64
	RFlatArmorPenetrationModPerLevel    float64
	RFlatArmorModPerLevel               float64
	RFlatHPRegenModPerLevel             float64
	RPercentTimeDeadModPerLevel         float64
	PercentEXPBonus                     float64
	RFlatCritDamageModPerLevel          float64
	RFlatMovementSpeedModPerLevel       float64
	RPercentArmorPenetrationMod         float64
	RPercentArmorPenetrationModPerLevel float64
	FlatEXPBonus                        float64
	FlatHPPoolMod                       float64
	FlatCritChanceMod                   float64
	FlatArmorMod                        float64
}
type MetaDataDto struct {
	Tier   string
	Type   string
	IsRune bool
}
type ItemTreeDto struct {
	Header string
	Tags   []string
}
type GroupDto struct {
	MaxGroupOwnable string
	Key             string
}
type BasicDataDto struct {
	Gold                 GoldDto
	PlainText            string
	HideFromAll          bool
	InStore              bool
	Into                 []string
	ID                   int
	Stats                BasicDataStatsDto
	Colloq               string
	Maps                 map[string]bool
	SpecialRecipe        int
	Image                ImageDto
	Description          string
	Tags                 []string
	RequiredChampion     string
	From                 []string
	Group                string
	ConsumeOnFull        bool
	Name                 string
	Consumed             bool
	SanitizedDescription string
	Depth                int
	Rune                 MetaDataDto
	Stacks               int
}

func (api *GoLOLAPI) StaticDataGetItems(version int, locale string, complete bool) (list ItemListDto) {
	options := map[string]string{}
	if version != 0 {
		options["version"] = strconv.Itoa(version)
	}
	if locale != "" {
		options["locale"] = locale
	}
	if complete {
		options["itemListData"] = "all"
	}
	uri, hasParameters := getEndpointURI("/lol/static-data/v3/items", options)
	response, e := api.RequestStaticData(uri, cache.NoExpiration, hasParameters)
	if e != nil {
		panic(e)
	}
	list = ItemListDto{}
	err := json.Unmarshal(response, &list)
	if err != nil {
		panic(err)
	}
	return
}

type LanguageStringsDto struct {
	Data    map[string]string
	Version string
	Type    string
}

func (api *GoLOLAPI) StaticDataGetLanguageStrings(version string, locale string) (strings LanguageStringsDto) {
	options := map[string]string{}
	if version != "" {
		options["version"] = version
	}
	if locale != "" {
		options["locale"] = locale
	}
	uri, hasParameters := getEndpointURI("/lol/static-data/v3/language-strings", options)
	response, e := api.RequestStaticData(uri, cache.NoExpiration, hasParameters)
	if e != nil {
		panic(e)
	}
	strings = LanguageStringsDto{}
	err := json.Unmarshal(response, &strings)
	if err != nil {
		panic(err)
	}
	return
}

func (api *GoLOLAPI) StaticDataGetLanguages() (list []string) {
	response, e := api.RequestStaticData("/lol/static-data/v3/languages", cache.NoExpiration, false)
	if e != nil {
		panic(e)
	}
	list = []string{}
	err := json.Unmarshal(response, &list)
	if err != nil {
		panic(err)
	}
	return
}

type MapDataDto struct {
	Data    map[string]MapDetailsDto
	Version string
	Type    string
}
type MapDetailsDto struct {
	MapName               string
	Image                 ImageDto
	ID                    float64 `json:"mapId"`
	UnpurchasableItemList []float64
}

func (api *GoLOLAPI) StaticDataGetMaps(version string, locale string) (maps MapDataDto) {
	options := map[string]string{}
	if version != "" {
		options["version"] = version
	}
	if locale != "" {
		options["locale"] = locale
	}
	uri, hasParameters := getEndpointURI("/lol/static-data/v3/maps", options)
	response, e := api.RequestStaticData(uri, cache.NoExpiration, hasParameters)
	if e != nil {
		panic(e)
	}
	maps = MapDataDto{}
	err := json.Unmarshal(response, &maps)
	if err != nil {
		panic(err)
	}
	return
}

type MasteryListDto struct {
	Data    map[string]MasteryDto
	Version string
	Tree    MasteryTreeDto
	Type    string
}
type MasteryTreeDto struct {
	Resolve  []MasteryTreeListDto
	Ferocity []MasteryTreeListDto
	Cunning  []MasteryTreeListDto
}
type MasteryTreeListDto struct {
	MasteryTreeItems []MasteryTreeItemDto
}
type MasteryTreeItemDto struct {
	ID     int `json:"masteryId"`
	Prereq string
}
type MasteryDto struct {
	Prereq               string
	MasteryTree          string
	Name                 string
	Ranks                int
	Image                ImageDto
	SanitizedDescription []string
	ID                   int
	Description          []string
}

func (api *GoLOLAPI) StaticDataGetMasteries(version int, locale string, complete bool) (list MasteryListDto) {
	options := map[string]string{}
	if version != 0 {
		options["version"] = strconv.Itoa(version)
	}
	if locale != "" {
		options["locale"] = locale
	}
	if complete {
		options["masteryListData"] = "all"
	}
	uri, hasParameters := getEndpointURI("/lol/static-data/v3/items", options)
	response, e := api.RequestStaticData(uri, cache.NoExpiration, hasParameters)
	if e != nil {
		panic(e)
	}
	list = MasteryListDto{}
	err := json.Unmarshal(response, &list)
	if err != nil {
		panic(err)
	}
	return
}

type RealmDto struct {
	Lg             string
	Dd             string
	L              string
	N              map[string]string
	ProfileIconMax int
	Store          string
	V              string
	Cdn            string
	Css            string
}

func (api *GoLOLAPI) StaticDataGetRealm() (realm RealmDto) {
	response, e := api.RequestStaticData("/lol/static-data/v3/realms", cache.NoExpiration, false)
	if e != nil {
		panic(e)
	}
	realm = RealmDto{}
	err := json.Unmarshal(response, &realm)
	if err != nil {
		panic(err)
	}
	return
}

type RuneListDto struct {
	Data    map[string]RuneDto
	Version string
	Type    string
	Basic   BasicDataDto
}
type RuneDto struct {
	PlainText            string
	HideFromAll          bool
	InStore              bool
	Into                 []string
	ID                   int
	Stats                BasicDataStatsDto
	Colloq               string
	Maps                 map[string]bool
	SpecialRecipe        int
	Image                ImageDto
	Description          string
	Tags                 []string
	RequiredChampion     string
	From                 []string
	Group                string
	ConsumeOnFull        bool
	Name                 string
	Consumed             bool
	SanitizedDescription string
	Depth                int
	Rune                 MetaDataDto
	Stacks               int
}

func (api *GoLOLAPI) StaticDataGetRunes(version string, locale string, complete bool) (list RuneListDto) {
	options := map[string]string{}
	if version != "" {
		options["version"] = version
	}
	if locale != "" {
		options["locale"] = locale
	}
	if complete {
		options["runeListData"] = "all"
	}
	uri, hasParameters := getEndpointURI("/lol/static-data/v3/runes", options)
	response, e := api.RequestStaticData(uri, cache.NoExpiration, hasParameters)
	if e != nil {
		panic(e)
	}
	list = RuneListDto{}
	err := json.Unmarshal(response, &list)
	if err != nil {
		panic(err)
	}
	return
}

type SummonerSpellListDto struct {
	Data    map[string]SummonerSpellDto
	Version string
	Type    string
}
type SummonerSpellDto struct {
	Vars                 []SpellVarsDto
	Image                ImageDto
	CostBurn             string
	Cooldown             []float64
	EffectBurn           []string
	ID                   int
	CooldownBurn         string
	Tooltip              string
	MaxRank              int
	RangeBurn            string
	Description          string
	Effect               [][]float64
	Key                  string
	LevelTip             LevelTipDto
	Modes                []string
	Ressource            string
	Name                 string
	CostType             string
	SanitizedDescription string
	SanitizedTooltip     string
	Range                []int
	Cost                 []int
	SummonerLevel        int
}

func (api *GoLOLAPI) StaticDataGetSummonerSpells(version string, locale string, complete bool) (list SummonerSpellListDto) {
	options := map[string]string{}
	if version != "" {
		options["version"] = version
	}
	if locale != "" {
		options["locale"] = locale
	}
	if complete {
		options["spellData"] = "all"
	}
	uri, hasParameters := getEndpointURI("/lol/static-data/v3/summoner-spells", options)
	response, e := api.RequestStaticData(uri, cache.NoExpiration, hasParameters)
	if e != nil {
		panic(e)
	}
	list = SummonerSpellListDto{}
	err := json.Unmarshal(response, &list)
	if err != nil {
		panic(err)
	}
	return
}
func (api *GoLOLAPI) StaticDataGetVersions() (versions []string) {
	response, e := api.RequestStaticData("/lol/static-data/v3/summoner-spells", cache.NoExpiration, false)
	if e != nil {
		panic(e)
	}
	versions = []string{}
	err := json.Unmarshal(response, &versions)
	if err != nil {
		panic(err)
	}
	return
}
