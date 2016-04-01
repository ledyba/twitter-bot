package main

import "github.com/ChimeraCoder/anaconda"

func register() []*Account {
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	return []*Account{
		NewAccount("", "", newEaaS()),
	}
}
