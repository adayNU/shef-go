package shef

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

type Music struct {
	Artist string `json:"by"`
	Album  string `json:"cd"`
	Title  string `json:"title"`
}

type GetTunedResponse struct {
	StationID     int                   `json:"stationId"`
	ProgramID     string                `json:"programId"`
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
	Offset        int                   `json:"offset"`
	Music         `json:"music"`
	Status        `json:"status"`
}

type KeyResponse struct {
	Hold   keyAction `json:"hold"`
	Key    key       `json:"key"`
	Status `json:"status"`
}

type VersionResponse struct {
	AccessCardID       string `json:"accessCardId"`
	LocationName       string `json:"locationName"`
	ModelID            int    `json:"modelId"`
	ReceiverID         string `json:"receiverId"`
	STBSoftwareVersion string `json:"stbSoftwareVersion"`
	SystemTime         int64  `json:"systemTime"`
	Version            string `json:"version"`
	Status             `json:"status"`
}

type Mode int

const (
	ModeActive Mode = iota
	ModeStandby
)

type ModeResponse struct {
	Mode   `json:"getMode"`
	Status `json:"status"`
}

type Location struct {
	ClientAddress string `json:"clientAddr"`
	LocationName  string `json:"locationName"`
	TunerBond     bool   `json:"tunerBond"`
}

type LocationResponse struct {
	Locations []Location `json:"locations"`
	Status    `json:"status"`
}
