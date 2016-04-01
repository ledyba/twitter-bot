package main

import "github.com/ChimeraCoder/anaconda"

// ".*?"([^\n])
func register() []*Account {
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	return []*Account{
		NewAccount("", "", newEaaS()),
		NewAccount("", "", newMii()),
		NewAccount("", "", newSeizonSenryaku()),
	}
}
