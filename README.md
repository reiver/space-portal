# SPACE-PORTAL

**space-portal** is a type of **gateway** server — and in particular, a TLS termination proxy.

**space-portal** sits between the Internet and the **space-base**s
—
traffic from the Internet is received by a **space-portal**, and then the **space-portal** sends it to whichever **space-base** it was addressed to.
Typically, **space-portal** would be running on a separate computer from the **space-base**s.

**space-portal** also hosts secure TLS certificates for HTTP servers on the **space-base**s.

**space-portal** is part of a tool-set for **self-hosting** various software, daemons, and applications.

## See Also

* [space-base](https://github.com/reiver/space-base)
* [space-command](https://github.com/reiver/space-command)
* [space-portal](https://github.com/reiver/space-portal)
