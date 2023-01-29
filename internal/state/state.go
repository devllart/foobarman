package state

import "devllart/foobarman/internal/drinks"

// Global state and variables

var Run = true
var Info = ""
var Money float64 = 300000000.33
var Name = ""
var Scene func()
var Bar = []drinks.Drink{}

// Flags

var Mix = false

// Temp State
var TempStr = ""
var TempBool = false
var Alerts = []string{}
var RandomBuy = false

var Command = ""
var Args = []string{}
var Arg1 = ""
var Arg2 = ""

var DrinksIds = []string{}
