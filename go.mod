module github.com/rustatian/GoPlayground

go 1.27.3

require (
	buf.build/gen/go/connectrpc/eliza/connectrpc/go v1.21.0-20230913231627-233fca715f49.1
	buf.build/gen/go/connectrpc/eliza/protocolbuffers/go v1.36.12-20230913231627-233fca715f49.2
	connectrpc.com/connect v1.21.1
	github.com/centrifugal/centrifuge-go v0.13.0
	github.com/go-acme/lego/v4 v4.35.2
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/golang/protobuf v1.5.4
	github.com/google/flatbuffers v25.12.19+incompatible
	github.com/google/go-github/v42 v42.0.0
	github.com/google/gopacket v1.1.19
	github.com/nats-io/nats.go v1.54.0
	github.com/quic-go/quic-go v0.64.1
	github.com/roadrunner-server/endure v1.4.6
	github.com/roadrunner-server/errors v1.6.0
	github.com/roadrunner-server/goridge/v3 v3.8.3
	github.com/spf13/afero v1.15.0
	github.com/tetratelabs/wazero v1.12.0
	github.com/valyala/fasthttp v1.76.0
	go.opentelemetry.io/otel v1.47.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.47.0
	go.opentelemetry.io/otel/sdk v1.47.0
	go.temporal.io/sdk v1.50.0
	go.temporal.io/sdk/contrib/opentelemetry v0.8.1
	gocv.io/x/gocv v0.43.0
	golang.org/x/crypto v0.58.0
	golang.org/x/oauth2 v0.37.0
	golang.org/x/sync v0.24.0
	google.golang.org/api v0.306.0
	google.golang.org/grpc v1.84.0
	google.golang.org/protobuf v1.36.12
	tailscale.com v1.104.2
)

require (
	cloud.google.com/go/auth v0.24.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.3.0 // indirect
	cloud.google.com/go/compute/metadata v0.10.0 // indirect
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/akutz/memconn v0.1.0 // indirect
	github.com/andybalholm/brotli v1.2.6 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/centrifugal/fdelta v0.0.3 // indirect
	github.com/centrifugal/protocol v0.23.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/coder/websocket v1.8.15 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dblohm7/wingoes v0.0.0-20260526185140-fb298caac7ca // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/felixge/httpsnoop v1.1.0 // indirect
	github.com/fxamacker/cbor/v2 v2.9.6 // indirect
	github.com/go-jose/go-jose/v4 v4.1.6 // indirect
	github.com/go-json-experiment/json v0.0.0-20260820222146-c27c302e5fc3 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/goccy/go-json v0.11.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/go-querystring v1.2.0 // indirect
	github.com/google/s2a-go v0.1.11 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.23 // indirect
	github.com/googleapis/gax-go/v2 v2.26.2 // indirect
	github.com/gorilla/websocket v1.5.4-0.20250319132907-e064f32e3674 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.3.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.32.0 // indirect
	github.com/hdevalence/ed25519consensus v0.2.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/jsimonetti/rtnetlink v1.4.2 // indirect
	github.com/klauspost/compress v1.20.3 // indirect
	github.com/mailru/easyjson v0.9.2 // indirect
	github.com/mdlayher/netlink v1.11.2 // indirect
	github.com/mdlayher/socket v0.8.0 // indirect
	github.com/miekg/dns v1.1.73 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/molecule-man/go-brrr v1.3.0 // indirect
	github.com/nats-io/nkeys v0.4.16 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nexus-rpc/nexus-proto-annotations v0.1.0 // indirect
	github.com/nexus-rpc/sdk-go v0.7.0 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/segmentio/asm v1.2.1 // indirect
	github.com/segmentio/encoding v0.5.4 // indirect
	github.com/shadowspore/fossil-delta v0.0.0-20241213113458-1d797d70cbe3 // indirect
	github.com/stretchr/objx v0.8.0 // indirect
	github.com/stretchr/testify v1.12.1 // indirect
	github.com/tailscale/go-winio v0.0.0-20231025203758-c4f33415bf55 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.72.0 // indirect
	go.opentelemetry.io/otel/log v1.47.0 // indirect
	go.opentelemetry.io/otel/metric v1.47.0 // indirect
	go.opentelemetry.io/otel/trace v1.47.0 // indirect
	go.temporal.io/api v1.64.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.28.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	go4.org/mem v0.0.0-20240501181205-ae6ca9944745 // indirect
	go4.org/netipx v0.0.0-20260823151212-3075585bcbeb // indirect
	golang.org/x/exp v0.0.0-20261014143641-f0251eb4aa3e // indirect
	golang.org/x/mod v0.42.0 // indirect
	golang.org/x/net v0.61.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/text v0.43.0 // indirect
	golang.org/x/time v0.16.0 // indirect
	golang.org/x/tools v0.52.0 // indirect
	golang.zx2c4.com/wireguard/windows v1.1.1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20261019180522-ce908d29299b // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20261019180522-ce908d29299b // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
