////////////////////////////////////////////////////////////////////////////
// Program: curlurl
// Purpose: curl like url glob handling in Go
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package curlurl

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const (
	URLPlain = iota
	URLSet
	URLRange
)

// The URLPattern holds the elements from url glob patterns.
type URLPattern struct {
	ptype   int // pattern type
	pattern []string
}

type URLGlob struct {
	URL     string
	urlGlob []URLPattern
}

type errorHandling func(errCase string, e error)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	reURLGlob = regexp.MustCompile(`[[{].*?[]}]`)
	reRange   = regexp.MustCompile(`\[(.*?)-(.*?)\]`)
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// NewURLGlob creates a new URLGlob
func NewURLGlob(hosts string) *URLGlob {
	return &URLGlob{URL: hosts}
}

// GetURLs will get URLs from the parsed urlGlob slice
func (g *URLGlob) GetURLs(ndx int) (ret []string) {
	if ndx+1 < len(g.urlGlob) {
		for _, p := range g.urlGlob[ndx].pattern {
			rest := Prefix(g.GetURLs(ndx+1), p)
			ret = append(ret, rest...)
		}
	} else {
		// already at the last item
		ret = g.urlGlob[ndx].pattern
	}
	return
}

func Prefix(in []string, prefix string) (ret []string) {
	for _, p := range in {
		ret = append(ret, prefix+p)
	}
	return
}

// Parse will parse URL into urlGlob
func (g *URLGlob) Parse(abortOn errorHandling) *URLGlob {
	// https://play.golang.org/p/KEiq7__4Ce
	indices := reURLGlob.FindAllStringSubmatchIndex(g.URL, -1)
	if len(indices) == 0 {
		// no glob, return whole string as single entry
		g.urlGlob = []URLPattern{
			{ptype: URLPlain, pattern: []string{g.URL}}}
		return g
	}
	// ii is char pointer into g.URL; jj is index of slice indices
	ii, jj := 0, 0
	for ii < len(g.URL) && jj < len(indices) {
		var up URLPattern
		til := indices[jj][0]
		if ii < til {
			// plain text exist before next match
			up.ptype = URLPlain
			up.pattern = []string{g.URL[ii:til]}
			ii = til
		} else {
			// here ii must == til, i.e., ii pointing to a match
			end := indices[jj][1]
			switch g.URL[til] { // use first char to determine type
			case '{':
				// set expression, from the opening '{'	til the next closing '}'
				// with ','-separated elements in between
				// e.g., site.{one,two,three}.com
				up.ptype = URLSet
				up.pattern = strings.Split(g.URL[til+1:end-1], ",")
			case '[':
				/* range expression, with the opening '[', and
				   - char range: e.g. "a-z]", "B-Q]"
				   - num range: e.g. "0-9]", "17-2000]"
				   - num range with leading zeros: e.g. "001-999]"
				   until the next ']'
				*/
				up.ptype = URLRange
				up.pattern =
					g.GlobRange(reRange.FindStringSubmatch(g.URL[til:end]), abortOn)
			}
			ii = end
			jj++
		}
		g.urlGlob = append(g.urlGlob, up)
	}
	if ii < len(g.URL) {
		// plain text exist after last match
		var up URLPattern
		up.ptype = URLPlain
		up.pattern = []string{g.URL[ii:]}
		g.urlGlob = append(g.urlGlob, up)
	}
	return g
}

// GlobRange handles range expression globing
func (g *URLGlob) GlobRange(submatch []string, abortOn errorHandling) []string {
	if len(submatch) != 3 {
		abortOn("Range expression parsing", errors.New("Wrong range format"))
	}
	start, err := strconv.Atoi(submatch[1])
	if err == nil {
		// OK, it is num range
		end, err := strconv.Atoi(submatch[2])
		abortOn("Range expression num range not match", err)
		var ret []string
		for ii := start; ii <= end; ii++ {
			ret = append(ret, strconv.Itoa(ii))
		}
		return ret
	}
	return nil
}
