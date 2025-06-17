# SPACE-PORTAL

**space-portal** is a type of **gateway** server — and in particular, a TLS termination proxy.

**space-portal** is part of a tool-set for **self-hosting** various software, daemons, and applications.

**space-portal** sits between the Internet and the **space-base**s
—
traffic from the Internet is received by a **space-portal**, and then the **space-portal** sends it to whichever **space-base** it was addressed to.

**space-portal** also hosts secure TLS certificates for HTTP servers on the **space-base**s.

Typically, **space-portal** would be running on a separate computer from the **space-base**s.
And, although you can run **space-portal** on a computer with a single Ethernet-port, you should be able to get better performance if the computer running **space-portal** has 2 Ethernet-ports
—
where one Ethernet-port (directly or indirectly) connects to the Internet, and the other Ethernet-port connects to the network that the **space-base**s are (also) connected to (ideally connected via an Ethernet-switch rather than an Ethernet-hub).

## See Also

* [space-base](https://github.com/reiver/space-base)
* [space-command](https://github.com/reiver/space-command)
* [space-portal](https://github.com/reiver/space-portal)
