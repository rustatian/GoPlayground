package main

import (
	"fmt"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/zipkin"
)

func main() {

	m := make(map[string]string)

	m["key"] = "fsd"
	m["key2"] = "fsd"

	fmt.Println("" +
		bitbucket.org/inturnco/productservice/application/repository/pg.(*Repo).collectExistingData(0xc00012ff80, 0xc000096720, 0x24, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/productservice/application/repository/pg.(*Repo).StoreProductsExtractMovingUnits(0xc00012ff80, 0xc0000102c0, 0xc0000102c8, 0xc000096720, 0x24, 0xc000096750, 0x24, 0xc0000966f0, 0x24, 0x6, ...)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/productservice/application.(*ImportsConfigProcessorService).ImportConfigProcessor(0xc0003b8ab0, 0x13e32c0, 0xc00010fbf0, 0xc000096720, 0x24, 0xc000096750, 0x24, 0xc0000966f0, 0x24, 0xc0000374c0, ...)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/productservice/imports.(*service).ImportConfigProcessor(0xc0003ba3d0, 0x13e32c0, 0xc00010fbf0, 0xc000096720, 0x24, 0xc000096750, 0x24, 0xc0000966f0, 0x24, 0xc0000374c0, ...)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/productservice/imports.makeImportInventoryConfigProcessorEndpoint.func1(0x13e32c0, 0xc00010fbf0, 0xf57300, 0xc0001122a0, 0x1, 0xc0000968d0, 0x2b, 0x5)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/go-sdk/blocking/redis.RedisBlockingScopeMiddleware.func1.1(0x13e32c0, 0xc00010fbf0, 0xf57300, 0xc0001122a0, 0x20, 0x13dcec0, 0xc00000e6a0, 0xc00010fc50)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/go-sdk/security.AuthorizationMiddlewareWithTrusted.func1.1(0x13e32c0, 0xc00010fbf0, 0xf57300, 0xc0001122a0, 0x4ddc0a, 0xf081c0, 0xc00010fbf0, 0xe505a0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 github.com/go-kit/kit/ratelimit.NewErroringLimiter.func1.1(0x13e32c0, 0xc00010fbf0, 0xf57300, 0xc0001122a0, 0xf058a0, 0xc00034a0e0, 0x13e32c0, 0xc00010fbf0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/go-sdk/tracing/zipkin.TraceEndpoint.func1.1(0x13e32c0, 0xc00010faa0, 0xf57300, 0xc0001122a0, 0x0, 0x0, 0x0, 0x0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/go-sdk/instrumenting.(*Middleware).WithMetrics.func1.1(0x13e32c0, 0xc00010faa0, 0xf57300, 0xc0001122a0, 0x0, 0x0, 0x0, 0x0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/go-sdk/logger.(*Middleware).WithLogging.func1.1(0x13e32c0, 0xc00010faa0, 0xf57300, 0xc0001122a0, 0xc0001122a0, 0x0, 0x0, 0x0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/go-sdk/mq/amqp/subscriber.(*BaseSubscriber).serveDelivery(0xc0003b8f30, 0xc0000d9680, 0x0, 0x0, 0x0, 0x0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/go-sdk/mq/amqp/subscriber.(*RetrySubscriber).serveDelivery.func1(0xc00034c000, 0xc00000f7a0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 github.com/cenkalti/backoff.RetryNotify(0xc00010f500, 0x13d4660, 0xc00034c000, 0x0, 0x0, 0x0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 github.com/cenkalti/backoff.Retry(...)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/go-sdk/mq/amqp/subscriber.(*RetrySubscriber).serveDelivery(0xc000375500, 0xc0000d9680, 0x0, 0x0, 0x0, 0x0)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 bitbucket.org/inturnco/go-sdk/mq/amqp/subscriber.(*RetrySubscriber).ServeDelivery(0xc000375500, 0xc0000d9680)
	/var/log/upstart/productservice.log inturn-dev-productservice-i-02a3c33ea1fc26ed7 created by bitbucket.org/inturnco/go-sdk/mq/amqp.(*Server).runNonBlockingListener")

}
