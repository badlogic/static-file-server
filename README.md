# static-file-server
A Go based static file server that serves a directory on a given port and also opens the browser to view `http://localhost:<port>`. See the [release page](https://github.com/badlogic/static-file-server/releases/tag/1.0.0) for binaries for Windows, Linux, and macOS.

Usage: `static-server <directory> <port> <relative-url>?`

E.g.

`static-server build/ 8000`

`static-server build/ 8000 test.html`
