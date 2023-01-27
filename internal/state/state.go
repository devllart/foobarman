package state

import "devllart/foobarman/internal/drinks"

// Global state and variables

var Run = true
var Info = ""
var Money = 300.33
var Name = ""
var Scene func()
var Bar = []drinks.Drink{}

// Temp State
var TempStr = ""

var Command = ""
var Args = []string{}
var Arg1 = ""
var Arg2 = ""

var DrinksIds = []string{}
