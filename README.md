# Golang HTTP_TOOL
- A tool which makes http requests and prints the address of the request along with the
MD5 hash of the response.
- This tool able to perform the requests in parallel so that the tool can complete sooner. The
order in which addresses are printed is not important.
- This tool is able to limit the number of parallel requests, to prevent exhausting local resources.
This accepts a flag to indicate this limit, and By default, parallel requests would be 10 if the flag is not provided.


# SETUP
Before you begin you must have Go installed and configured properly for your computer. Please see
https://golang.org/doc/install


# Getting Started
Clone the git repository in your system and then cd into project root directory
https://github.com/Deepak1124/myhttp_tool.git

Build your tool by executing the following steps
$ cd myhttp_tool
$ go build



# EXAMPLES
$ go run .\myhttptool.go http://www.adjust.com http://google.com
http://www.adjust.com d1b40e2a2ba488a054186e4ed0733f9752f66949
http://google.com 9d8ec921bdd275fb2a605176582e08758eb60641


$ go run .\myhttptool.go adjust.com
http://adjust.com d1b40e2a2ba488a054186e4ed0733f9752f66949


$ go run .\myhttptool.go -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
http://google.com 8ff1c478ccca08cca025b028f68b352f
http://adjust.com 6b2560b9a5262571258cc173248b7492
http://yandex.com 4baab01ff9ff0f793bf423aeef539c9d
http://facebook.com ccae5ffa91c4936aef3efd5091a43f3e
http://twitter.com 857efe81a54c8b5c2241846ac4f08e66
http://reddit.com/r/funny ff3b2b7dcd9e716ca0adcbd208061c9a
http://reddit.com/r/notfunny ff3b2b7dcd9e716ca0adcbd208061c9a
http://yahoo.com e2d50a30b7bfbfda097d72e32578c6a6
http://baroquemusiclibrary.com 8e5138a0111364f08b10d37ed3371b11




