// Copyright © 2018 Charles Haynes <ceh@ceh.bz>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ratelimiter

import "time"

// RateLimiter limits calles to at most n calls in t time period
type RateLimiter struct {
	t  time.Duration
	ch chan *time.Timer
}

// NewRateLimiter constructs a rate limiter given a count and duration
func New(n int, t time.Duration) (r RateLimiter) {
	r.t = t
	if n > 0 {
		r.ch = make(chan *time.Timer, n)
		for i := 0; i < n; i++ {
			r.ch <- time.NewTimer(0)
		}
	}
	return r
}

// Limit will only return n times in t duration. If called more frequently
// it blocks until the limit is satisfied
func (r RateLimiter) Limit() {
	if r.ch == nil {
		return
	}
	t := <-r.ch
	<-t.C
	r.ch <- time.NewTimer(r.t)
}
