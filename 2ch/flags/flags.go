package main

import (
	"flag"
	"fmt"
)

// Config struct use for saving flags value
type Config struct {
	subject string
	isAwesome bool
	howAwesome int
	countTheWays CountTheWays
}

// Setup function initailize sent config value from flags value
func (c *Config) Setup() {
	// As following, it can set flag directly 
	// var something = flag.String("subject", "Awesome", "Awesome")
	// But It is better to input value in structor
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	
	// abbreviation
	flag.StringVar(&c.subject,"s", "", "subject is a string, it defaults to empty")
	flag.BoolVar(&c.isAwesome, "isAwesome", false, "is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howAwesome", 0, "how awesome out of 10?")

	// custom variables type
	flag.Var(&c.countTheWays,"c","comma separated list of integers")

}

// GetMessage function use all inner config variables, and return sentence

func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
        msg += " is awesome"
	}else {
		msg += " is not awesome"
	}
	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())
	return msg
}
