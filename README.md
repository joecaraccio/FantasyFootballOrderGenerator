# Fantasy Order Generate

Goal to generate a completely Fair Fantasy Football Order. Completely Transparent and fully open sourced. 



## Procedure

1. The source of this executable, generated simply from 'main.go' is publically available in this repository. This source will be available and unchanged, unless the hardcoded list of names must be changed. It uses all public, open source modules (link in comments) that allow reading. No other sources. 

1. The SHA256 chechsum shall be printed at the end of this file. The SHA256 is `64dd014a38a282448f4fc44f212b8bcceb55b0ad3ecd2915f6c8ba5b355c3c25`. This SHA256 will be printed at the end of the output and must be used to confirm the results. If this SHA256 does not match the end of the generate file. The results are invalid.

1. On pick order night, once > 7 League Members (exclusing the League Commisioner Joe) and at least 1 Non League Member are in the Discord, and can confirm that they can see the PICKERS (Jon K)'s Discord Stream, the process may begin.

1. PICKER (Jon K) must navigate to this Repo and download it. This will download a .zip of the whole repo, then simply extract it.  [Link to Repo Download Here](https://github.com/joecaraccio/FantasyFootballOrderGenerator/archive/refs/heads/main.zip). The top right hamburger menu when clicked, provides a download link.

1. Once succesfully clicked, the PICKER (Jon K) must show the downloaded executable has an accurate download timestamp to the download time. Must wait for the entire discord call to confirm the download process before proceeding.

1. Once downloaded, the RANDOMIZER (Teigue) will request a number between 3 and 9 (including 3 and 9). This is to represent another layer of randomization.

1. Once the PICKER (Jon K) confirms the number he must run the executable the requested number of times. If the RANDOMIZER (Teigue) requests 3 iterations, the 3rd run on the executable will represent the order. If the RANDOMIZER (Teigue) requests 8, the 8th run of the executable will represent the oder.

1. After each iteration, the RANDOMIZER (Teigue) and 1 additional member must confirm the iteration is complete.

1. On the final iteration is complete, the order is determined, the SHA256 checksum (printed at the end) must match the value listed prior in this document. This is officially the picking order and will be used to determine picking for Draft positions. This order should be screenshot and pasted into discord.




## Technical Details 

Use GoLang 1.24.5
Verify by this by running `go version`

My output is: `go version go1.24.5 windows/amd64`

To simply test: `go run main.go`

To generate an executable: `go build -o generate_fantasy_order.exe main.go`
