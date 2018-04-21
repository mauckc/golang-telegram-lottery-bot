You must create bot_token.go file, which include TOKEN variable in global package scope
i.e.
*** in file bot_token.go

package main

var TOKEN = "YOUR_TOKEN_HERE"

***

You also must make sure you have a db.sqlite3 file in this directory


You will need to "go get" shell command the packages that will be imported

I converted the russian words to english in the number to word funciton and non-admin response

