Some Go Code
============

goServer1
---------

Accepts multiple telnet connections using a go-routine. Prints out what is received on the server's STDOUT. Kills a connection if "end" is received.

goHttpServer1
-------------

Basic HTTP server that serves out the current directory over an OS chosen port.
Serving a small plain text file does about 3000 reqs/sec on a 27" iMac.
There seem to be some strange stuttering/stalling when ab is run against the server
Going to try serving something from memory in order to eliminate possible IO bottleneck. (goHttpServer2)

goHttpServer2
-------------
Serves out the string "Hello!" over HTTP, from an in memory constant, similar to goHttpServer1 but no file system IO.
Does about 8000 reqs/sec, still stalls sometimes though. (OSX)
File system IO was clearly a limiting factor but isn't causing the stalling.
Seems to be an OSX issue, running ab against the server from an Ubuntu VM doesn't see any stuttering/stalling
Does about 11000 reqs/sec when benchmarked from the Ubuntu VM (server running on OSX host machine)