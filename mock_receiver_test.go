package shef

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

func NewMockReceiver() *httptest.Server {
	var r = mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(handler)
	return httptest.NewServer(r)
}

const (
	tunedTV    = `{"callsign":"FOODHD","date":"20070324","duration":1791,"episodeTitle":"Spaghetti and Clam Sauce","expiration":"0","expiryTime":0,"isOffAir":false,"isPartial":false,"isPclocked":1,"isPpv":false,"isRecording":true,"isViewed":true,"isVod":false,"keepUntilFull":true,"major":231,"minor":65535,"offset":263,"programId":"4405732","rating":"No Rating","recType":3,"startTime":1278342008,"stationId":3900976,"status":{"code":200,"commandResult":0,"msg":"OK.","query":"/tv/getTuned"},"title":"Tyler's Ultimate","uniqueId":"6728716739474078694"}`
	tunedMusic = `{"callsign":"MUSICTV","date":"20070324","duration":1791,"expiration":"0","expiryTime":0,"isOffAir":false,"isPartial":false,"isPclocked":3,"isPpv":false,"isRecording":false,"isVod":false,"major":231,"minor":65535,"offset":263,"programId":"4405732","rating":"No Rating","startTime":1278342008,"stationId":3900976,"status":{"code":200,"commandResult":0,"msg":"OK.","query":"/tv/getProgInfo"},"title":"90s Hits","uniqueId":"6728716739474078694","music":{"by":"Baha Men","cd":"Who Let The Dogs Out","title":"Who Let The Dogs Out"}}`
	tuned      = `{"status":{"code":200,"commandResult":0,"msg": "OK","query":"/tv/tune?major=508"}}`
	keyPress   = `{"hold":"keyPress","key":"exit","status":{"code":200,"commandResult":0,"msg":"OK.","query":"/remote/processKey?key=info&hold=keyPress"}}`
	version    = `{"accessCardId":"0021-1495-6572","receiverId":"0288 7745 5858","status":{"code":200,"commandResult":0,"msg":"OK","query":"/info/getVersion"},"stbSoftwareVersion":"0x4ed7","systemTime":1281625203,"version":"1.2"}`
	mode       = `{"mode":0"status":{"code":200,"commandResult":0,"msg":"OK","query":"/info/mode"}}`
	locations  = `{"locations":[{"clientAddr":"0","locationName":"A"}],"status":{"code":200,"commandResult":0,"msg":"OK.","query":"/info/getLocations?callback=jsonp"}}`
)

var (
	parsedTunedTV = &GetTunedResponse{
		StationID:     3900976,
		ProgramID:     "4405732",
		StartTime:     1278342008,
		Duration:      1791,
		Major:         231,
		Minor:         NoMinor,
		CallSign:      "FOODHD",
		IsOffAir:      false,
		IsVOD:         false,
		IsPPV:         false,
		IsRecording:   true,
		Rating:        "No Rating",
		IsPCLocked:    Locked,
		Date:          "20070324",
		Title:         "Tyler's Ultimate",
		EpisodeTitle:  "Spaghetti and Clam Sauce",
		UniqueID:      "6728716739474078694",
		KeepUntilFull: true,
		IsViewed:      true,
		Expiration:    "0",
		ExpiryTime:    0,
		IsPartial:     false,
		Offset:        263,
		Status: Status{
			Code:          http.StatusOK,
			CommandResult: 0,
			Msg:           "OK.",
			Query:         getTuned,
		},
	}
	parsedGetProgInfo = &GetTunedResponse{
		StationID:   3900976,
		ProgramID:   "4405732",
		StartTime:   1278342008,
		Duration:    1791,
		Major:       231,
		Minor:       NoMinor,
		CallSign:    "MUSICTV",
		IsOffAir:    false,
		IsVOD:       false,
		IsPPV:       false,
		IsRecording: false,
		Rating:      "No Rating",
		IsPCLocked:  Unlocked,
		Date:        "20070324",
		Title:       "90s Hits",
		UniqueID:    "6728716739474078694",
		Expiration:  "0",
		ExpiryTime:  0,
		IsPartial:   false,
		Offset:      263,
		Music: Music{
			Artist: "Baha Men",
			Album:  "Who Let The Dogs Out",
			Title:  "Who Let The Dogs Out",
		},
		Status: Status{
			Code:          http.StatusOK,
			CommandResult: 0,
			Msg:           "OK.",
			Query:         getProgramInfo,
		},
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	var body string

	switch r.URL.Path {
	case getTuned:
		body = tunedTV
	case getProgramInfo:
		body = tunedMusic
	case tune:
		body = tuned
	case remoteProcessKey:
		body = keyPress
	case getVersion:
		body = version
	case getOptions:
		body = ""
	case getMode:
		body = mode
	case getLocations:
		body = locations
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write([]byte(body))
}
