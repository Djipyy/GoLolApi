package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

//MatchList This object contains a list of matches.
type MatchList struct {
	Matches    []MatchReference
	TotalGames int
	StartIndex int
	EndIndex   int
}

//MatchReference This object contains a reference to a match.
type MatchReference struct {
	Lane       string
	ID         float64 `json:"gameId"`
	Champion   int
	PlatformID string
	Season     int
	Queue      int
	Role       string
	Timestamp  float64
}

//MatchData This object contains the data of a match.
type MatchData struct {
	SeasonID              int
	QueueID               int
	GameID                float64
	ParticipantIdentities []ParticipantIdentity
	GameVersion           string
	PlatformID            string
	GameMode              string
	MapID                 int
	GameType              string
	Teams                 []TeamStats
	Participants          []Participant
	GameDuration          float64
	GameCreation          float64
}

//ParticipantIdentity This object contains the identity of a participant. Only works in ranked matches.
type ParticipantIdentity struct {
	Player Player
	ID     int `json:"participantId"`
}

//Player This object contains player information.
type Player struct {
	CurrentPlatformID string
	SummonerName      string
	MatchHistoryURI   string
	PlatformID        string
	CurrentAccountID  float64
	ProfileIcon       int
	SummonerID        float64
	AccountID         float64
}

//TeamStats This object contains statistics of a match.
type TeamStats struct {
	FirstDragon          bool
	FirstInhibitor       bool
	Bans                 []TeamBans
	BaronKills           int
	FirstRiftHerald      bool
	FirstBaron           bool
	RiftHeraldKills      int
	FirstBlood           bool
	TeamID               int
	FirstTower           bool
	VilemawKills         int
	InhibitorKills       int
	TowerKills           int
	DominionVictoryScore int
	Win                  string
	DragonKills          int
}

//TeamBans This object contains the bans of a team.
type TeamBans struct {
	PickTurn   int
	ChampionID int
}

//Participant This object contains participant data.
type Participant struct {
	Stats                     ParticipantStats
	ParticipantID             int
	Runes                     []RuneDto
	Timeline                  ParticipantTimeline
	TeamID                    int
	Spell2Id                  int
	Masteries                 []MasteryDto
	HighestAchievedSeasonTier string
	Spell1ID                  int
	ChampionID                int
}

//ParticipantStats This object contains the statistics of a Participant.
type ParticipantStats struct {
	PhysicalDamageDealt             float64
	NeutralMinionsKilledTeamJungle  int
	MagicDamageDealt                float64
	TotalPlayerScore                int
	Deaths                          int
	Win                             bool
	NeutralMinionsKilledEnemyJungle int
	AltarsCaptured                  int
	LargestCriticalStrike           int
	TotalDamageDealt                float64
	MagicDamageDealtToChampions     float64
	MagicalDamageTaken              float64
	VisionWardsBoughtInGame         int
	LargestMultiKill                int
	LargestKillingSpree             int
	Item1                           int
	QuadraKills                     int
	TeamObjective                   int
	TotalTimeCrowdControlDealt      int
	Float64estTimeSpentLiving       int
	WasAfk                          bool
	WardsKilled                     int
	FirstTowerAssist                bool
	EarlySurrenderAccomplice        bool
	Leaver                          bool
	FirstTowerKill                  bool
	Item2                           int
	Item3                           int
	CausedEarlySurrender            bool
	FirstBloodAssist                bool
	Item6                           int
	WardsPlaced                     int
	Item4                           int
	Item5                           int
	TurretKills                     int
	TripleKills                     int
	ChampLevel                      int
	NodeNeutralizeAssist            int
	FirstInhibitorKill              bool
	GoldEarned                      int
	TeamEarlySurrendered            bool
	UnrealKills                     int
	Kills                           int
	DoubleKills                     int
	NodeCaptureAssist               int
	TrueDamageTaken                 float64
	NodeNeutralize                  int
	FirstInhibitorAssist            bool
	Assists                         int
	TotalScoreRank                  int
	NeutralMinionsKilled            int
	ObjectivePlayerScore            int
	CombatPlayerScore               int
	AltarsNeutralized               int
	PhysicalDamageDealtToChampions  float64
	GoldSpent                       int
	TrueDamageDealt                 float64
	GameEndedInSurrender            bool
	TrueDamageDealtToChampions      float64
	ParticipantID                   int
	PentaKills                      int
	GameEndedInEarlySurrender       bool
	TotalHeal                       float64
	TotalMinionsKilled              int
	FirstBloodKill                  bool
	NodeCapture                     int
	SightWardsBoughtInGame          int
	TotalDamageDealtToChampions     float64
	TotalUnitsHealed                int
	InhibitorKills                  int
	TotalDamageTaken                float64
	KillingSprees                   int
	Item0                           int
	PhysicalDamageTaken             float64
}

//ParticipantTimeline This object contains the timeline of a Participant.
type ParticipantTimeline struct {
	Lane                        string
	ParticipantID               int
	CsDiffPerMinDeltas          map[string]float64
	GoldPerMinDeltas            map[string]float64
	XpDiffPerMinDeltas          map[string]float64
	CreepsPerMinDeltas          map[string]float64
	XpPerMinDeltas              map[string]float64
	Role                        string
	DamageTakenDiffPerMinDeltas map[string]float64
	DamageTakenPerMinDeltas     map[string]float64
}

//MatchTimeline This object contains the timeline of a match.
type MatchTimeline struct {
	Frames        []MatchFrame
	FrameInterval float64
}

//MatchFrame This object contains a frame of a timeline.
type MatchFrame struct {
	Timestamp float64
	Frames    map[int]MatchParticipantFrame `json:"participantFrames"`
	Events    []MatchEvent
}

//MatchParticipantFrame This object contains a participant data in a MatchFrame.
type MatchParticipantFrame struct {
	TotalGold           int
	TeamScore           int
	ParticipantID       int
	Level               int
	CurrentGold         int
	MinionsKilled       int
	DominionScore       int
	Position            MatchPosition
	Xp                  int
	JungleMinionsKilled int
}

//MatchPosition This object contains the coordinates of a participant.
type MatchPosition struct {
	X, Y int
}

//MatchEvent This object contains event data.
type MatchEvent struct {
	Timestamp float64
	Type      string
}

//GetRecentMatches Get matchlist for last 20 matches played of given summoner
func (s *Summoner) GetRecentMatches() (list MatchList) {
	response, e := s.API.RequestEndpoint("/lol/match/v3/matchlists/by-account/"+strconv.FormatFloat(s.AccountID, 'f', -1, 64)+"/recent", time.Minute)
	if e != nil {
		panic(e)
	}
	list = MatchList{}
	err := json.Unmarshal(response, &list)
	if err != nil {
		panic(err)
	}
	return
}

//GetMatchData Get match by match ID
func (api *GoLOLAPI) GetMatchData(id float64) (match MatchData) {
	response, e := api.RequestEndpoint("/lol/match/v3/matches/"+strconv.FormatFloat(id, 'f', -1, 64), 24*time.Hour)
	if e != nil {
		panic(e)
	}
	match = MatchData{}
	err := json.Unmarshal(response, &match)
	if err != nil {
		panic(err)
	}
	return
}

//GetMatchTimeline Get match timeline by match ID.
func (api *GoLOLAPI) GetMatchTimeline(id float64) (timeline MatchTimeline) {
	response, e := api.RequestEndpoint("/lol/match/v3/timelines/by-match/"+strconv.FormatFloat(id, 'f', -1, 64), 24*time.Hour)
	if e != nil {
		panic(e)
	}
	timeline = MatchTimeline{}
	err := json.Unmarshal(response, &timeline)
	if err != nil {
		panic(err)
	}
	return
}
