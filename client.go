package shef

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	NoMinor = 65535
)

const (
	protocol         = "http://"
	getTuned         = "/tv/getTuned"
	getProgramInfo   = "/tv/getProgInfo"
	tune             = "/tv/tune"
	remoteProcessKey = "/remote/processKey"
	getVersion       = "/info/getVersion"
	getOptions       = "/info/getOptions"
	getMode          = "/info/getMode"
	getLocations     = "/info/getLocations"
)

type Client struct {
	client   *http.Client
	receiver *Receiver
}

// NewClient returns a |*Client|. If a nil *http.Client is passed,
// the |*Client| uses an |http.DefaultClient|.
func NewClient(c *http.Client, r *Receiver) *Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &Client{
		client:   c,
		receiver: r,
	}
}

// GetTuned returns a |*GetTunedResponse| or an error on error.
func (c *Client) GetTuned() (*GetTunedResponse, error) {
	var gtr = &GetTunedResponse{}
	var err = c.get(getTuned, "", gtr)

	return gtr, err
}

// GetProgramInfoMajor re
func (c *Client) GetProgramInfoMajor(major int) (*GetTunedResponse, error) {
	return c.GetProgramInfoAtTime(major, NoMinor, nil)
}

func (c *Client) GetProgramInfo(major, minor int) (*GetTunedResponse, error) {
	return c.GetProgramInfoAtTime(major, minor, nil)
}

func (c *Client) GetProgramInfoAtTime(major, minor int, time *time.Time) (*GetTunedResponse, error) {
	var q = make(url.Values)
	q.Add("major", string(major))
	q.Add("minor", string(minor))
	if time != nil {
		q.Add("time", string(time.Unix()))
	}

	var gtr = &GetTunedResponse{}
	var err = c.get(getProgramInfo, q.Encode(), gtr)

	return gtr, err
}

func (c *Client) TuneMajor(major int) (*Status, error) {
	return c.Tune(major, NoMinor)
}

func (c *Client) Tune(major, minor int) (*Status, error) {
	var q = make(url.Values)
	q.Add("major", string(major))
	q.Add("minor", string(minor))

	var s = &Status{}
	var err = c.get(tune, q.Encode(), s)

	return s, err
}

func (c *Client) Press(k key, a keyAction) (*KeyResponse, error) {
	var q = make(url.Values)
	q.Add("key", string(k))
	q.Add("hold", string(a))

	var kr = &KeyResponse{}
	var err = c.get(remoteProcessKey, q.Encode(), kr)

	return kr, err
}

func (c *Client) GetVersion() (*VersionResponse, error) {
	var v = &VersionResponse{}
	var err = c.get(getVersion, "", v)

	return v, err
}

func (c *Client) GetMode() (*ModeResponse, error) {
	var m = &ModeResponse{}
	var err = c.get(getMode, "", m)

	return m, err
}

func (c *Client) GetLocations() (*LocationResponse, error) {
	var lr = &LocationResponse{}
	var err = c.get(getLocations, "", lr)

	return lr, err
}

var (
	ErrNotModified        = errors.New("not modified")
	ErrConflict           = errors.New("conflict")
	ErrServiceUnavailable = errors.New("service unavailable")
)

func (c *Client) get(path, query string, s interface{}) error {
	var resp, err = c.client.Get(protocol + c.receiver.Address() + path + "?" + query)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		// No-op.
	case http.StatusNotModified:
		return ErrNotModified
	case http.StatusBadRequest, http.StatusForbidden:
		return errors.New("bad / forbidden request")
	case http.StatusConflict:
		return ErrConflict
	case http.StatusInternalServerError, http.StatusHTTPVersionNotSupported:
		return errors.New("unfulfillable request")
	case http.StatusServiceUnavailable:
		return ErrServiceUnavailable
	}

	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, s)
}
