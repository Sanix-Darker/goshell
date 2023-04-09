package main

// Fix the space issue (when we are writing an input)
// concatenate multiple line when we're executing them
// keep an history of what we are executing
// capture the arrows keys to put to the stdin the appropriage value
const checkTrueThis = true
const VERSION = "0.0.1-alpha"
const WATERMARK = `GoShell ` + VERSION + ` by dk
Type "help" for more information.
Hit "reset" to clean all instructions
Hit "resetimport" to clean all imports`
const TEMP_FILE = "/tmp/golang-shell-tmpcode.go"
const HIST_FILE = "/tmp/golang-shell-tmphistory.txt"

// we are going to store all instructions here
var CODE_INPUT []string
var VAR_INPUT []string
var IMPORTS_INPUT []string
