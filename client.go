package shef

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Receiver struct {
	ip   string
	port int
}

func (r *Receiver) Address() string {
	return r.ip + ":" + string(r.port)
}

const (
	protocol         = "http://"
	getTuned         = "/tv/getTuned"
	getProgInfo      = "/tv/getProgInfo"
	tune             = "/tv/tune"
	remoteProcessKey = "/remote/processKey"
	getVersion       = "/info/getVersion"
	getOptions       = "/info/getOptions"
	mode             = "/info/mode"
	getLocations     = "/info/getLocations"
)

type Client struct {
	client   *http.Client
	receiver Receiver
}

type Status struct {
	Code          int    `json:"code"`
	CommandResult int    `json:"commandResult"`
	Msg           string `json:"msg"`
	Query         string `json:"query"`
}

type ParentalControlStatus int

const (
	PCInvalid ParentalControlStatus = iota
	Locked
	TemporarilyUnlocked
	Unlocked
)

type RecordingType int

const (
	RTInvalid RecordingType = iota
	Manual
	FindBy
	Regular
	Recurring
)

type GetTunedResponse struct {
	StationID     int                   `json:"stationId"`
	ProgramID     int                   `json:"programId"`
	MaterialID    int                   `json:"materialId"`
	StartTime     int64                 `json:"startTime"`
	Duration      int                   `json:"duration"`
	Major         int                   `json:"major"`
	Minor         int                   `json:"minor"`
	CallSign      string                `json:"callsign"`
	IsOffAir      bool                  `json:"isOffAir"`
	IsVOD         bool                  `json:"isVod"`
	IsPPV         bool                  `json:"isPpv"`
	IsPurchased   bool                  `json:"isPurchased"`
	IsRecording   bool                  `json:"isRecording"`
	Rating        string                `json:"rating"`
	IsPCLocked    ParentalControlStatus `json:"isPclocked"`
	Date          string                `json:"date"`
	Title         string                `json:"title"`
	EpisodeTitle  string                `json:"episodeTitle"`
	UniqueID      string                `json:"uniqueId"`
	KeepUntilFull bool                  `json:"keepUntilFull"`
	IsViewed      bool                  `json:"isViewed"`
	Expiration    string                `json:"expiration"`
	ExpiryTime    int64                 `json:"expiryTime"`
	RecordingType RecordingType         `json:"recordingType"`
	FindByWord    string                `json:"findByWord"`
	IsPartial     bool                  `json:"isPartial"`
	Priority      string                `json:"priority"`
	Offset        string                `json:"offset"`
	Music         struct {
		Artist string `json:"by"`
		Album  string `json:"cd"`
		Title  string `json:"title"`
	}
	Status `json:"status"`
}

func (c *Client) GetTuned() (*GetTunedResponse, error) {
	var resp, err = c.client.Get(protocol + c.receiver.Address() + getTuned)
	if err != nil {
		return nil, err
	}

	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gtr = &GetTunedResponse{}
	err = json.Unmarshal(b, gtr)

	return gtr, err
}

type InfoOpts struct {
	Minor      int
	Time       time.Time
	ClientAddr string
}

func (c *Client) GetProgramInfo(major int) (*GetTunedResponse, error) {

}

var (
	c *Client
)
