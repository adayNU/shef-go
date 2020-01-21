package shef

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/go-check/check"
)

type ClientSuite struct {
	server *httptest.Server
	client *Client
}

func (cs *ClientSuite) SetUpSuite(c *check.C) {
	cs.server = NewMockReceiver()

	var u, err = url.Parse(cs.server.URL)
	c.Assert(err, check.IsNil)

	cs.client = &Client{
		client:   http.DefaultClient,
		receiver: NewReceiver(u.Hostname(), u.Port()),
	}
}

func (cs *ClientSuite) TearDownSuite(c *check.C) {
	cs.server.Close()
}

func (cs *ClientSuite) TestAPICalls(c *check.C) {
	var gtr, err = cs.client.GetTuned()
	c.Check(err, check.IsNil)
	c.Check(gtr, check.DeepEquals, parsedTunedTV)

	gtr, err = cs.client.GetProgramInfoMajor(123)
	c.Check(err, check.IsNil)
	c.Check(gtr, check.DeepEquals, parsedGetProgInfo)

}

var _ = check.Suite(&ClientSuite{})
