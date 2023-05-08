# Welcome to **Etherniti**

> Etherniti - Infura Like Open Source Services

## Usage

```bash
docker run \
    -it \
    -d etherniti/proxy-oss:latest
```

## Prerequisites

### Minimum hardware requirements

* 256 MB of RAM
* 1 GB of drive space (although 10 GB is a recommended minimum if running Jenkins as a Docker container)
   
### Recommended hardware configuration for small DAPPs

* 1 GB+ of RAM
* 50 GB+ of drive space

## Installation platforms

This section describes how to install/run Jenkins on different platforms and operating systems.
Docker

### Docker

Docker is a platform for running applications in an isolated environment called a "container" (or Docker container). Applications like Etherniti can be downloaded as read-only "images" (or Docker images), each of which is run in Docker as a container. A Docker container is in effect a "running instance" of a Docker image. From this perspective, an image is stored permanently more or less (i.e. insofar as image updates are published), whereas containers are stored temporarily. Read more about these concepts in the Docker documentation’s Getting Started, Part 1: Orientation and setup page.

Docker’s fundamental platform and container design means that a single Docker image (for any given application like Etherniti) can be run on any supported operating system (macOS, Linux and Windows) or cloud service (AWS and Azure) which is also running Docker.

### Unikernels

> cooming Soon

## Deployment

You can configure your server to use a RT kernel via:

```bash
sudo apt-get install linux-headers-lowlatency
sudo apt-get install linux-lowlatency
sudo update-grub
```

After that, run:

```bash
docker run \
    -it \
    -d etherniti/proxy-oss:latest
```

Not: add as many environment variables as needed in order to customize your deployment.

## Development

### Using Intellij GoLand

In order to execute builds and run them from IDE, you will need to pass all required env variables.
You can copy and paste them to workspace.xml file or add them manually one by one

```xml
<env name="X_ETHERNITI_LOG_LEVEL" value="debug" />
<env name="X_ETHERNITI_SSL_CERT_FILE" value="/etc/letsencrypt/live/etherniti.org/fullchain.pem" />
<env name="X_ETHERNITI_SSL_KEY_FILE" value="/etc/letsencrypt/live/etherniti.org/privkey.pem" />
<env name="X_ETHERNITI_SECURE_LISTENING_PORT" value="4430" />
<env name="X_ETHERNITI_DEBUG_SERVER" value="true" />
<env name="X_ETHERNITI_HIDE_SERVER_DATA_IN_CONSOLE" value="true" />
<env name="X_ETHERNITI_TOKEN_SECRET" value="your-favourite-jwt-token-secret" />
<env name="X_ETHERNITI_ENABLE_HTTPS_REDIRECT" value="false" />
<env name="X_ETHERNITI_ENABLE_LOGGING" value="true" />
<env name="X_ETHERNITI_ENABLE_PROFILER" value="true" />
<env name="X_ETHERNITI_USE_UNIQUE_REQUEST_ID" value="false" />
<env name="X_ETHERNITI_ENABLE_SECURITY" value="true" />
<env name="X_ETHERNITI_ENABLE_ANALYTICS" value="true" />
<env name="X_ETHERNITI_ENABLE_INTERNAL_ANALYTICS" value="true" />
<env name="X_ETHERNITI_ENABLE_CORS" value="true" />
<env name="X_ETHERNITI_ENABLE_CACHE" value="true" />
<env name="X_ETHERNITI_ENABLE_RATE_LIMIT" value="false" />
<env name="X_ETHERNITI_BLOCK_TOR_CONNECTIONS" value="false" />
<env name="X_ETHERNITI_LISTENING_PORT" value="8080" />
<env name="X_ETHERNITI_LISTENING_INTERFACE" value="127.0.0.1" />
<env name="X_ETHERNITI_LISTENING_ADDRESS" value="127.0.0.1" />
<env name="X_ETHERNITI_SWAGGER_ADDRESS" value="127.0.0.1" />
<env name="X_ETHERNITI_TOKEN_EXPIRATION" value="6000" />
<env name="X_ETHERNITI_RATELIMIT" value="10" />
<env name="X_ETHERNITI_RATE_LIMIT_UNITS" value="10" />
<env name="X_ETHERNITI_INFURA_TOKEN" value="$YOUR_INFURA_KEY" />
<env name="X_ETHERNITI_LISTENING_MODE" value="http" />
```

### Dependencies by size

Following, **etherniti** dependencies are listed, orderer by impact on final executable size:

```bash
  3.4 MB net/http
  3.0 MB runtime
  2.0 MB github.com/zerjioang/etherniti/vendor/github.com/json-iterator/go
  1.7 MB net
  1.6 MB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/ssh
  1.6 MB github.com/zerjioang/etherniti/vendor/github.com/golang/protobuf/proto
  1.4 MB reflect
  1.3 MB github.com/zerjioang/etherniti/vendor/golang.org/x/sys/unix
  1.2 MB github.com/zerjioang/etherniti/vendor/github.com/modern-go/reflect2
  1.2 MB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4
  1.2 MB github.com/zerjioang/etherniti/vendor/github.com/prometheus/client_golang/prometheus/promhttp
  1.1 MB github.com/zerjioang/etherniti/vendor/github.com/dgraph-io/badger
  1.0 MB github.com/zerjioang/etherniti/vendor/github.com/prometheus/client_golang/prometheus
  932 kB crypto/tls
  871 kB math/big
  766 kB github.com/zerjioang/go-hpc/lib/eth/fixtures/crypto/secp256k1
  753 kB encoding/gob
  713 kB syscall
  631 kB crypto/x509
  626 kB text/template
  596 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/object
  594 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/openpgp/packet
  570 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/format/packfile
  546 kB encoding/json
  535 kB github.com/zerjioang/go-hpc/thirdparty/echo
  510 kB html/template
  509 kB text/template/parse
  508 kB github.com/zerjioang/etherniti/vendor/github.com/gorilla/websocket
  495 kB github.com/zerjioang/etherniti/core/controllers
  447 kB vendor/golang_org/x/text/unicode/norm
  424 kB time
  421 kB github.com/zerjioang/go-hpc/lib/eth/rpc
  415 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/protocol/packp
  410 kB github.com/zerjioang/etherniti/vendor/github.com/prometheus/procfs
  403 kB regexp/syntax
  396 kB vendor/golang_org/x/net/dns/dnsmessage
  389 kB fmt
  360 kB github.com/zerjioang/etherniti/vendor/github.com/prometheus/common/model
  351 kB regexp
  340 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/storage/filesystem
  338 kB image
  336 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/storage/filesystem/dotgit
  330 kB github.com/zerjioang/etherniti/core/controllers/network
  328 kB github.com/zerjioang/etherniti/vendor/github.com/sergi/go-diff/diffmatchpatch
  328 kB compress/flate
  322 kB github.com/zerjioang/go-hpc/lib/eth/fixtures/abi
  321 kB github.com/zerjioang/etherniti/vendor/github.com/prometheus/common/expfmt
  320 kB github.com/zerjioang/etherniti/vendor/golang.org/x/net/trace
  319 kB os
  312 kB archive/zip
  307 kB runtime/pprof
  283 kB crypto/elliptic
  276 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/ssh/agent
  271 kB encoding/asn1
  269 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/openpgp
  265 kB github.com/zerjioang/etherniti/vendor/github.com/kevinburke/ssh_config
  263 kB github.com/zerjioang/go-hpc/thirdparty/decimal
  260 kB vendor/golang_org/x/crypto/cryptobyte
  247 kB github.com/zerjioang/etherniti/vendor/github.com/tidwall/gjson
  246 kB strconv
  241 kB image/jpeg
  240 kB vendor/golang_org/x/text/unicode/bidi
  230 kB strings
  225 kB unicode
  222 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/transport/server
  219 kB os/exec
  209 kB github.com/zerjioang/go-hpc/thirdparty/jwt-go
  204 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/format/idxfile
  204 kB math
  201 kB github.com/zerjioang/etherniti/vendor/github.com/dgraph-io/badger/table
  200 kB vendor/golang_org/x/net/idna
  191 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/storage/memory
  189 kB flag
  186 kB mime
  184 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/transport/http
  181 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/transport/ssh
  179 kB html
  179 kB internal/poll
  177 kB bytes
  176 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/format/index
  174 kB github.com/zerjioang/etherniti/vendor/github.com/prometheus/client_model/go
  170 kB bufio
  170 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-billy.v4/memfs
  168 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/ed25519/internal/edwards25519
  167 kB vendor/golang_org/x/net/http2/hpack
  163 kB crypto/rsa
  157 kB net/http/httptest
  155 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/transport/internal/common
  154 kB encoding/binary
  152 kB net/textproto
  150 kB github.com/zerjioang/go-hpc/thirdparty/echo/middleware
  149 kB net/url
  148 kB github.com/zerjioang/etherniti/vendor/github.com/src-d/gcfg
  148 kB mime/multipart
  147 kB github.com/zerjioang/go-hpc/lib/bip39/wordlists
  147 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/config
  144 kB github.com/zerjioang/etherniti/vendor/github.com/dgraph-io/badger/y
  140 kB crypto/cipher
  137 kB github.com/zerjioang/etherniti/core/controllers/common
  136 kB os/user
  132 kB io
  131 kB github.com/zerjioang/etherniti/vendor/golang.org/x/net/internal/timeseries
  131 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing
  126 kB github.com/zerjioang/etherniti/vendor/github.com/dgraph-io/badger/protos
  126 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/storer
  123 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-billy.v4/osfs
  121 kB github.com/zerjioang/etherniti/core/db
  120 kB github.com/zerjioang/etherniti/core/listener/middleware
  120 kB sync
  119 kB github.com/zerjioang/etherniti/vendor/github.com/dgraph-io/badger/skl
  118 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/ssh/knownhosts
  117 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/format/config
  115 kB github.com/zerjioang/etherniti/shared/protocol
  112 kB sort
  111 kB runtime/cgo
  111 kB github.com/zerjioang/etherniti/core/controllers/project
  111 kB github.com/zerjioang/etherniti/core/model/registry
  110 kB github.com/zerjioang/etherniti/core/listener/https
  110 kB github.com/zerjioang/etherniti/core/controllers/registry
  108 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/internal/revision
  106 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/sha3
  104 kB github.com/zerjioang/etherniti/core/model/project
  103 kB github.com/zerjioang/etherniti/core/api
  102 kB github.com/zerjioang/etherniti/core/controllers/ws
  101 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/utils/merkletrie
  101 kB github.com/zerjioang/etherniti/core/config
  100 kB github.com/zerjioang/etherniti/core/listener/socket
   99 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-billy.v4/helper/chroot
   98 kB github.com/zerjioang/go-hpc/thirdparty/middleware/logger
   98 kB vendor/golang_org/x/text/transform
   98 kB crypto/aes
   98 kB github.com/zerjioang/etherniti/vendor/github.com/dgryski/go-farm
   96 kB github.com/zerjioang/go-hpc/lib/metrics/prometheus_metrics
   96 kB math/rand
   94 kB vendor/golang_org/x/crypto/chacha20poly1305
   94 kB github.com/zerjioang/go-hpc/lib/concurrentmap
   93 kB context
   93 kB expvar
   92 kB github.com/zerjioang/etherniti/core/model/auth
   90 kB path/filepath
   90 kB github.com/zerjioang/go-hpc/lib/eth/fixtures
   89 kB github.com/zerjioang/go-hpc/lib/solc
   88 kB github.com/zerjioang/go-hpc/thirdparty/gommon/log
   87 kB crypto/ecdsa
   85 kB github.com/zerjioang/etherniti/core/listener/http
   85 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/format/diff
   85 kB github.com/zerjioang/etherniti/vendor/github.com/emirpasic/gods/lists/arraylist
   85 kB compress/bzip2
   83 kB net/http/pprof
   82 kB github.com/zerjioang/go-hpc/lib/radix
   81 kB github.com/zerjioang/etherniti/vendor/github.com/src-d/gcfg/token
   81 kB github.com/zerjioang/etherniti/core/server/ratelimit
   81 kB github.com/zerjioang/etherniti/core/listener/common
   79 kB crypto/sha512
   78 kB hash/fnv
   78 kB github.com/zerjioang/go-hpc/lib/cyber
   78 kB image/color
   75 kB github.com/zerjioang/go-hpc/lib/eth
   72 kB crypto/sha256
   72 kB encoding/base64
   71 kB github.com/zerjioang/etherniti/vendor/github.com/beorn7/perks/quantile
   71 kB github.com/zerjioang/go-hpc/lib/bip32
   71 kB github.com/zerjioang/etherniti/vendor/github.com/emirpasic/gods/trees/binaryheap
   71 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/transport/file
   70 kB github.com/zerjioang/etherniti/vendor/github.com/src-d/gcfg/scanner
   70 kB crypto/sha1
   69 kB github.com/zerjioang/etherniti/vendor/github.com/tidwall/pretty
   68 kB github.com/zerjioang/etherniti/shared/mixed
   68 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/transport
   68 kB github.com/zerjioang/go-hpc/thirdparty/gommon/color
   68 kB github.com/zerjioang/etherniti/vendor/github.com/pkg/errors
   67 kB crypto/x509/pkix
   66 kB compress/gzip
   65 kB vendor/golang_org/x/net/http/httpproxy
   65 kB github.com/zerjioang/go-hpc/lib/bip39
   63 kB github.com/zerjioang/go-hpc/lib/hashset
   62 kB github.com/zerjioang/etherniti/core/model/metadata
   62 kB github.com/zerjioang/go-hpc/lib/eth/profile
   60 kB net/http/internal
   59 kB github.com/zerjioang/go-hpc/lib/eth/fixtures/crypto
   58 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/format/objfile
   58 kB text/tabwriter
   58 kB github.com/zerjioang/etherniti/vendor/github.com/modern-go/concurrent
   58 kB github.com/zerjioang/go-hpc/lib/snowflake
   58 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/utils/ioutil
   57 kB github.com/zerjioang/go-hpc/thirdparty/template
   57 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/utils/merkletrie/filesystem
   57 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/bcrypt
   57 kB net/http/httptrace
   55 kB io/ioutil
   55 kB container/list
   55 kB hash/crc32
   54 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/openpgp/armor
   54 kB github.com/zerjioang/etherniti/core/keystore/memory
   54 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/cast5
   54 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-billy.v4
   53 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-billy.v4/helper/polyfill
   52 kB crypto/rand
   52 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/protocol/packp/capability
   52 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/transport/git
   52 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/format/gitignore
   51 kB log
   50 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/internal/chacha20
   50 kB github.com/zerjioang/etherniti/vendor/github.com/src-d/gcfg/types
   48 kB crypto/des
   47 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/storage
   47 kB github.com/zerjioang/etherniti/vendor/github.com/valyala/bytebufferpool
   46 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/revlist
   46 kB github.com/zerjioang/go-hpc/lib/eth/fixtures/common/math
   46 kB github.com/zerjioang/etherniti/vendor/github.com/AndreasBriese/bbloom
   45 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-billy.v4/util
   45 kB compress/zlib
   45 kB crypto/md5
   44 kB encoding/pem
   44 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/openpgp/s2k
   43 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/cache
   42 kB encoding/hex
   41 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/poly1305
   40 kB os/signal
   40 kB github.com/zerjioang/etherniti/vendor/github.com/emirpasic/gods/containers
   40 kB vendor/golang_org/x/crypto/internal/chacha20
   38 kB mime/quotedprintable
   38 kB vendor/golang_org/x/text/secure/bidirule
   38 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/protocol/packp/sideband
   37 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/ripemd160
   36 kB crypto/dsa
   36 kB github.com/zerjioang/etherniti/vendor/github.com/prometheus/client_golang/prometheus/internal
   36 kB github.com/zerjioang/go-hpc/lib/httpclient
   36 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/blowfish
   36 kB github.com/zerjioang/etherniti/vendor/github.com/emirpasic/gods/utils
   36 kB github.com/zerjioang/go-hpc/lib/eth/fixtures/common
   34 kB github.com/zerjioang/go-hpc/util/id
   34 kB github.com/zerjioang/go-hpc/lib/encoding
   33 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/ed25519
   32 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/format/pktline
   32 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/utils/merkletrie/noder
   32 kB internal/cpu
   31 kB path
   30 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/utils/merkletrie/index
   30 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/curve25519
   30 kB github.com/zerjioang/go-hpc/lib/eth/rpc/model
   30 kB vendor/golang_org/x/crypto/curve25519
   30 kB github.com/zerjioang/etherniti/vendor/github.com/zerjioang/go-bus/mutex
   29 kB github.com/zerjioang/etherniti/vendor/github.com/jbenet/go-context/io
   29 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/transport/client
   28 kB github.com/zerjioang/etherniti/vendor/github.com/pelletier/go-buffruneio
   28 kB github.com/zerjioang/etherniti/core/server/mem
   27 kB github.com/zerjioang/go-hpc/lib/worker
   27 kB github.com/zerjioang/etherniti/vendor/github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg
   27 kB runtime/trace
   26 kB vendor/golang_org/x/net/http/httpguts
   26 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/utils/merkletrie/internal/frame
   26 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/utils/binary
   26 kB unicode/utf8
   26 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/plumbing/filemode
   25 kB github.com/zerjioang/go-hpc/lib/tor
   25 kB internal/singleflight
   24 kB crypto
   24 kB math/bits
   23 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/openpgp/errors
   22 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/openpgp/elgamal
   22 kB github.com/zerjioang/etherniti/vendor/github.com/mattn/go-colorable
   22 kB hash/adler32
   21 kB github.com/zerjioang/go-hpc/lib/fastime
   21 kB github.com/zerjioang/etherniti/core/controllers/security
   21 kB github.com/zerjioang/go-hpc/lib/packers
   21 kB github.com/zerjioang/etherniti/core/server/disk
   21 kB runtime/debug
   20 kB github.com/zerjioang/go-hpc/lib/eth/paramencoder/erc20
   20 kB github.com/zerjioang/go-hpc/lib/integrity
   20 kB github.com/zerjioang/etherniti/vendor/gopkg.in/warnings.v0
   19 kB github.com/zerjioang/etherniti/core/data
   19 kB github.com/zerjioang/go-hpc/lib/interval
   19 kB github.com/zerjioang/etherniti/vendor/github.com/mitchellh/go-homedir
   19 kB crypto/rc4
   18 kB sync/atomic
   18 kB github.com/zerjioang/go-hpc/lib/eth/paramencoder
   17 kB vendor/golang_org/x/crypto/poly1305
   16 kB github.com/zerjioang/go-hpc/util/str
   16 kB github.com/zerjioang/etherniti/vendor/github.com/tidwall/match
   15 kB image/internal/imageutil
   15 kB github.com/zerjioang/etherniti/vendor/github.com/emirpasic/gods/lists
   15 kB internal/bytealg
   14 kB hash
   14 kB github.com/zerjioang/go-hpc/lib/cache
   14 kB container/heap
   14 kB github.com/zerjioang/go-hpc/lib/counter32
   14 kB github.com/zerjioang/etherniti/vendor/github.com/xanzy/ssh-agent
   14 kB crypto/hmac
   13 kB internal/testlog
   12 kB github.com/zerjioang/etherniti/core/listener/swagger
   12 kB github.com/zerjioang/etherniti/core/logger
   12 kB github.com/zerjioang/etherniti/core/controllers/tokenlist
   12 kB github.com/zerjioang/etherniti/vendor/github.com/prometheus/procfs/internal/fs
   12 kB github.com/zerjioang/go-hpc/lib/stack
   12 kB github.com/zerjioang/etherniti/vendor/golang.org/x/net/context
   11 kB github.com/zerjioang/go-hpc/util/ip
   11 kB github.com/zerjioang/etherniti/core/cmd
   11 kB runtime/internal/sys
   11 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/utils/diff
   10 kB runtime/internal/atomic
  9.8 kB github.com/zerjioang/etherniti/vendor/github.com/matttproud/golang_protobuf_extensions/pbutil
  9.8 kB github.com/zerjioang/etherniti/core/model/registry/version
  8.6 kB github.com/zerjioang/go-hpc/thirdparty/gommon/random
  8.6 kB github.com/zerjioang/go-hpc/lib/encoding/base58
  8.5 kB github.com/zerjioang/go-hpc/util/banner
  8.4 kB github.com/zerjioang/etherniti/core/listener
  8.0 kB unicode/utf16
  7.9 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/pbkdf2
  7.9 kB github.com/zerjioang/etherniti/vendor/github.com/zerjioang/go-bus
  7.7 kB crypto/internal/randutil
  7.5 kB github.com/zerjioang/go-hpc/lib/badips
  7.4 kB github.com/zerjioang/go-hpc/lib/encoding/hex
  7.1 kB github.com/zerjioang/go-hpc/lib/bots
  6.7 kB encoding
  6.4 kB github.com/zerjioang/etherniti/core/bus
  6.4 kB github.com/zerjioang/etherniti/vendor/github.com/emirpasic/gods/trees
  6.0 kB vendor/golang_org/x/crypto/cryptobyte/asn1
  5.8 kB github.com/zerjioang/etherniti/vendor/gopkg.in/src-d/go-git.v4/internal/url
  5.7 kB crypto/subtle
  5.6 kB internal/syscall/unix
  5.2 kB github.com/zerjioang/etherniti/shared/constants
  4.4 kB github.com/zerjioang/go-hpc/lib/aeshash
  4.3 kB github.com/zerjioang/etherniti/shared/def/listener
  4.2 kB internal/nettrace
  3.6 kB internal/race
  3.4 kB errors
  3.3 kB github.com/zerjioang/etherniti/vendor/golang.org/x/crypto/internal/subtle
  3.1 kB crypto/internal/subtle
  2.5 kB github.com/zerjioang/etherniti/vendor/github.com/mattn/go-isatty
  2.1 kB github.com/zerjioang/go-hpc/util/fs
  1.5 kB github.com/zerjioang/etherniti/vendor/github.com/dgraph-io/badger/options
```

## Performance tips

### Avoid the use of pointers in large heap scenarios

* Strings, slices and time.Time all contain pointers
* Lots of strings
* Timestamps on objects using time.Time
* Maps with slice values
* Maps with string keys

### Compilation time analysis

```bash
go build -gcflags='-m -m' $file.go
```

#### Detecting too complex functions

```bash
go build -gcflags='-m -m' $file.go | grep 'function too complex'
```

#### Detecting heap escapes

```bash
go build -gcflags='-m -m' $file.go | grep 'escapes to heap'
```

#### Bound Check analysis

Note: `bce` stands for `bound check elimination`
Note: `ssa` stands for `static single assignment`
 
```bash
go build -gcflags=-d=ssa/check_bce/debug=1 $file.go
```

#### Other `ssa` flags

```bash
go build -gcflags=-d=ssa/insert_resched_checks/on,ssa/check/on
```

### Going deeper with `GOSSAFUNC`
 
```bash
GOSSAFUNC=pattern go build $file.go
```

Other environment [magical environment](https://github.com/golang/go/blob/master/src/cmd/go/internal/work/exec.go#L233) you might use are `GOCLOBBERDEADHASH`, `GOSSAFUNC`, `GO_SSA_PHI_LOC_CUTOFF`, `GOSSAHASH`

## License

All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

 * Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 * Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 * Uses GPL license described below

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
