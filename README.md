Destination-Isolating i2p HTTP Proxy(SAM Application)
=====================================================

*one eepSite, one tunnel.*

This is an i2p SAM application which presents an HTTP proxy(on port 4443 by
default) that acts as an intermediate between your browser and the i2p network.
Then it uses the SAM library to create a unique destination for each i2p site
that you visit. This way, your unique destination couldn't be used to track you
with a network of colluding sites. I doubt it's a substantial problem right now
but it might be someday. Facebook has an onion site, and i2p should have
destination isolation before there is a facebook.i2p.

Excitingly, after like a year of not being able to devote the time I should
have to this, I've finally made it work. I am successfully using this proxy to
browse eepSites pretty normally. There are still a few bugs to chase down, but
if somebody else wanted to try it out, I might not *totally* embarass myself.

What works so far:
------------------

### The http proxy

Again, *slightly less experimental*, but currently it is possible to set
your web browser's HTTP proxy to localhost:4443 and use it to browse eepSites.
I've mostly managed to stop all the bugs and hangs that would keep someone from
using it successfully.

This is pretty exciting though, when a site contains a resource which it pulls
from another site, the proxy requests that down it's own tunnel, and not via the
tunnel associated with the site that contains the resource. If a tunnel doesn't
exist, it's created. This can be used to identify that you are using this
version of the proxy though, but not much else.

#### Examples

##### curl

        curl -x 127.0.0.1:4443 http://i2p-projekt.i2p

##### surf

        export http_proxy="http://127.0.0.1:4443" surf http://i2p-projekt.i2p

#### Current Concerns:

If it wasn't super, super obvious to everyone, it's really, really easy to tell
the difference between this proxy and the default i2p/i2pd http proxies and I
don't think there's anything I can do about that.

I haven't been able to observe any DNS leaks yet, but that doesn't mean they
aren't there. My plan is to implement some kind of proper URL validation for it.

Before version 0.21, a framework for generating service tunnels ad-hoc will also
be in place. This will be used for fuzz-testing the http proxy and the pipe
proxy. Almost everything will be improved by the availability of this.

I haven't implemented addresshelpers/jump host interaction yet, but I have a
good idea how to now.

### The pipes

Moved to [misc/docs/PIPES.md](https://github.com/eyedeekay/si-i2p-plugin/tree/master/misc/docs/PIPES.md)

What I'm doing right now:
-------------------------

Implementing pipe-controlled service tunnels.

What the final version should do:
---------------------------------

The final version should use the parent pipe and the aggregating pipe to send
and recieve requests as an http proxy in the familiar way.

Version Roadmap:

  * ~~0.17 - Named Pipes work for top-level i2p domains and can retrieve~~
   ~~directories under a site~~
  * ~~0.18 - Named Pipes for i2p domains and can retrieve subdirectories,~~
   ~~which it caches in clearly-named folders as normal files(Containing HTML)~~
  * ~~0.19 - Expose an http proxy that hooks up to the existing infrastructure~~
   ~~for destination isolation~~
  * 0.20 - ~~Ready for more mainstream testing~~, ~~should successfully isolate~~
   ~~requests for resources embedded in the retrieved web pages~~ and should be
   able to generate services on the fly by talking to the SAM bridge.
  * 0.21 - Addresshelper. First worthwhile release for people who aren't shell
  enthusiasts.

Silly Questions I'm asking myself about how I want it to work:
--------------------------------------------------------------

Definitely not going to do any more filtering than I already do. Instead I'm
going to do build a different proxy to do that and demonstrate how to connect
the two.

Installation and Usage:
=======================

Moved to [misc/docs/INSTALL.md](https://github.com/eyedeekay/si-i2p-plugin/tree/master/misc/docs/INSTALL.md)
