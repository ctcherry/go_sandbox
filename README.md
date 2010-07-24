Some Go Code
============

goServer1
---------

Accepts multiple telnet connections using a go-routine. Prints out what is received on the server's STDOUT. Kills a connection if "end" is received.

goHttpServer1
-------------

Basic HTTP server that serves out the current directory over an OS chosen port.
Serving a small plain text file does about 3000 reqs/sec on a 27" iMac.
There seem to be some strange concurrency limits that get hit?
Going to try serving something from memory in order to eliminate possible IO bottleneck.