/*
 * Copyright (c) 2019 shawn1m. All rights reserved.
 * Use of this source code is governed by The MIT License (MIT) that can be
 * found in the LICENSE file..
 */

// Package outbound implements multiple dns client and dispatcher for outbound connection.
package clients

import (
	"math/rand"
	"time"

	"github.com/miekg/dns"

	"github.com/shawn1m/overture/core/common"
	"github.com/shawn1m/overture/core/rule"
)

type RuleClient struct {
	responseMessage *dns.Msg
	questionMessage *dns.Msg

	minimumTTL   int
	domainTTLMap map[string]uint32

	defaultIP string
	rawName   string
	rules     []rule.Rule
}

func NewRuleClient(q *dns.Msg, defaultIP string, rules []rule.Rule, minimumTTL int, domainTTLMap map[string]uint32) *RuleClient {
	c := &RuleClient{
		questionMessage: q.Copy(),
		defaultIP:       defaultIP,
		minimumTTL:      minimumTTL,
		domainTTLMap:    domainTTLMap,
		rules:           rules,
	}
	c.rawName = c.questionMessage.Question[0].Name
	return c
}

func (c *RuleClient) Exchange() *dns.Msg {
	name := c.rawName[:len(c.rawName)-1]

	for _, rule := range c.rules {
		if rule.Match(name) {
			if c.exchangeFromIP() {
				if c.responseMessage != nil {
					common.SetMinimumTTL(c.responseMessage, uint32(c.minimumTTL))
					common.SetTTLByMap(c.responseMessage, c.domainTTLMap)
				}
				return c.responseMessage
			}
		}
	}

	return nil
}

func (c *RuleClient) exchangeFromIP() bool {
	a, _ := dns.NewRR(c.rawName + " IN A " + c.defaultIP)
	c.setLocalResponseMessage([]dns.RR{a})
	return true
}

func (c *RuleClient) setLocalResponseMessage(rrl []dns.RR) {
	shuffleRRList := func(rrl []dns.RR) {
		rand.Seed(time.Now().UnixNano())
		for i := range rrl {
			j := rand.Intn(i + 1)
			rrl[i], rrl[j] = rrl[j], rrl[i]
		}
	}

	c.responseMessage = new(dns.Msg)
	for _, rr := range rrl {
		c.responseMessage.Answer = append(c.responseMessage.Answer, rr)
	}
	shuffleRRList(c.responseMessage.Answer)
	c.responseMessage.SetReply(c.questionMessage)
	c.responseMessage.RecursionAvailable = true
}
