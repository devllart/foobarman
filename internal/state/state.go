package state

import "devllart/foobarman/internal/products"

// Local
var cmds = map[string]CommandStruct{}

// Global state and variables

// ** Flags

var Run = true        // State game loop
var Mix = false       // Is command mix?
var RandomBuy = false // For seperate logics rand buy and buy of user

// ** Temp State
var TempBool = false // Nothing say
var TempStr = ""     // Nothing say

// ** For manipulation with views
var ProductsIds = []string{} // Save output printers products

// ** For alerts user
var Alerts = []string{}
var Info = ""

// ** State of player
var Scene func()               // Context of scenes
var LastScene string           // Last scene
var Name = ""                  // Barman name
var RawName = ""               // Barman name in raw format
var Status = "Norm"            // Status barman
var Money float64 = 300.33     // Player money
var Bar = []products.Product{} // State of bar
var CountVisitorsServed = 0    // Number of visitors served
var Stage = 0                  // STAGE

// ** State of client

var BarOpen = false
var NotSaler = true
var CoctailReady = false
var Order products.Coctail
var YourCoctail products.Coctail
var CurrentHistory = []string{}

// ** For work with commands
var Command = ""      // Current command
var Args = []string{} // Current args of command
