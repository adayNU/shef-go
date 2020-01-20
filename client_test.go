package shef

import (
	"net/http"
	"net/http/httptest"

	"github.com/go-check/check"
	"github.com/gorilla/mux"
)

var resp = `{"callsign":"FOODHD","date":"20070324","duration":1791,"episodeTitle":"SpaghettiandClamSauce","expiration":"0","expiryTime":0,"isOffAir":false,"isPartial":false,"isPclocked":1,"isPpv":false,"isRecording":false,"isViewed":true,"isVod":false,"keepUntilFull":true,"major":231,"minor":65535,"offset":263,"programId":"4405732","rating":"NoRating","recType":3,"startTime":1278342008,"stationId":3900976,"status":{"code":200,"commandResult":0,"msg":"OK.","query":"/tv/getTuned"},"title":"Tyler'sUltimate","uniqueId":"6728716739474078694"}`

type ClientSuite struct {
	server *httptest.Server
}

func (c *ClientSuite) SetUpSuite(_ *check.C) {
	var r = mux.NewRouter()
	r.HandleFunc(getTuned, getTunedHandler)
	c.server = httptest.NewServer(r)
}

func getTunedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(resp))
}

var _ = check.Suite(&ClientSuite{})
