package cryptowatch

import (
	"net/url"
	"strconv"
	"strings"
)

type TradesParams struct {
	Limit int64
	Since int64
}

type urlEncodable interface {
	toValues() url.Values
}

func urlEncode(p urlEncodable) string {
	v := p.toValues()
	return v.Encode()
}

func (p *TradesParams) toValues() url.Values {
	v := url.Values{}
	if p.Limit > 0 {
		v.Add("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Since > 0 {
		v.Add("since", strconv.Itoa(int(p.Since)))
	}
	return v
}

type OHLCParams struct {
	Before  int64
	After   int64
	Periods []int64
}

func (p *OHLCParams) toValues() url.Values {
	v := url.Values{}
	if p.Before > 0 {
		v.Add("before", strconv.Itoa(int(p.Before)))
	}
	if p.After > 0 {
		v.Add("after", strconv.Itoa(int(p.After)))
	}
	if len(p.Periods) > 0 {
		s := make([]string, len(p.Periods))
		for i, j := range p.Periods {
			s[i] = strconv.Itoa(int(j))
		}
		v.Add("periods", strings.Join(s, ","))
	}
	return v
}
