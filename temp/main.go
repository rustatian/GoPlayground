package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

var esClient *elastic.Client

func main() {
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()
	// ELASTICSEARCH index storage
	args := os.Args[1:]
	const BULK_SIZE = 200

	if len(args) < 2 {
		panic("argument filepath and index name missing")
	}

	esOptions := make([]elastic.ClientOptionFunc, 0, 4)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	esOptions = append(esOptions,
		elastic.SetHttpClient(client),
		elastic.SetURL("https://localhost:8443"),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
	)
	var err error
	esClient, err = elastic.NewClient(esOptions...)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	indexName := fmt.Sprintf("%s-test", args[1])

	ic := fmt.Sprintf(InventoryProductMapping, 1, 1)
	createResp, err := esClient.CreateIndex(indexName).BodyJson(ic).Do(ctx)
	if err != nil {
		panic(err)
	}
	if createResp.Acknowledged == false {
		panic("index not awk")
	}

	resetInterval(ctx, esClient, indexName, "-1")

	var products []ProductIdx

	var wg sync.WaitGroup
	for {
		obj, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(products) > 0 {
				wg.Add(1)
				go indexer(ctx, indexName, products, &wg)
			}
			break
		}
		if err != nil {
			panic(err)
		}

		var p ProductIdx
		err = json.Unmarshal([]byte(obj), &p)
		if err != nil {
			// Deserialization failed
			panic(err)
		}
		products = append(products, p)
		if len(products) >= BULK_SIZE {
			fmt.Println("Send Bulk...")
			wg.Add(1)
			go indexer(ctx, indexName, products, &wg)
			products = nil
		}
	}
	wg.Wait()
	resetInterval(ctx, esClient, indexName, "1s")
	//... not sure what this time is returning
	elapsed := time.Since(start)
	fmt.Printf("Indexed products in %f seconds", elapsed.Seconds())

}

func indexer(ctx context.Context, indexName string, products []ProductIdx, wg *sync.WaitGroup) {
	bulkClient := esClient.Bulk().Index(indexName)

	for _, product := range products {
		bdr := elastic.NewBulkIndexRequest().UseEasyJSON(true).
			Type("products").
			Id(*product.UUID).
			Doc(product)
		bulkClient.Add(bdr)
	}
	bulkRequest, err := bulkClient.Do(ctx)
	if err != nil {
		panic(err)
	}
	ff := bulkRequest.Failed()
	if len(ff) != 0 {
		panic("some bulk items failed")
	}

	wg.Done()
}

func resetInterval(ctx context.Context, client *elastic.Client, indexName string, interval string) {
	obj := map[string]interface{}{}
	obj["index"] = struct {
		Refresh string `json:"refresh_interval"`
	}{
		interval,
	}
	resp, err := client.IndexPutSettings(indexName).BodyJson(obj).Do(ctx)
	if err != nil {
		panic(err)
	}
	if resp.Acknowledged == false {
		panic("index not awk")
	}
}

//InventoryProductMapping is an ES mapping used for products index
const InventoryProductMapping = `
{
  "settings": {
    "index": {
      "number_of_shards": %d,
      "number_of_replicas": %d
    },
    "analysis": {
    "filter": {
            "length_filter": {
               "type": "length",
               "min": 3               
            }
    },
	  "normalizer": {
        "lowercase-normalizer": {
          "type": "custom",
          "char_filter": [],
          "filter": ["lowercase"]
        }
      },
      "analyzer": {
        "path-analyzer": {
          "tokenizer": "path-tokenizer"
        },
        "char-group-analyzer": {
               "type": "custom",
               "tokenizer": "char_group_tokenizer",
               "filter": [
                  "length_filter",
                  "lowercase"
               ]
        },
        "lowercase-analyzer": {
               "type": "custom",
               "tokenizer": "whitespace",
               "filter": [
                  "lowercase"
               ]
		    }
      },
      "tokenizer": {
        "path-tokenizer": {
          "type": "path_hierarchy",
          "delimiter": "/"
        },
         "char_group_tokenizer": {
            "type": "char_group",
            "tokenize_on_chars": [
              " ",
              "-",
              "'",
              ".",
              "_",
              "\n"
            ]
        }
      }
    }
  },
  "mappings": {
    "products": {
	  "dynamic": false,
      "properties": {
        "uuid": {
          "type": "keyword"
        },
        "companyUUID": {
          "type": "keyword"
        },
        "name": {
          "type": "text",
          "analyzer": "char-group-analyzer",
          "fields": {
			"raw": {
			  "type": "keyword"
			}
		  }
        },
		"description": {
          "type": "text",
          "analyzer": "char-group-analyzer"
        },
        "sku": {
          "type": "keyword",
          "normalizer": "lowercase-normalizer"
        },
        "createdAt": {
          "type": "date"
        },
        "updatedAt": {
          "type": "date"
        },
        "createdByUUID": {
          "type": "keyword"
        },
        "options": {
          "type": "nested",
          "properties": {
            "valueUUID": {
              "type": "keyword"
            },
            "name": {
              "type": "keyword"
            },
            "value": {
              "type": "keyword"
            },
            "inturnType": {
               "type": "keyword"
            },
            "position": {
              "type": "integer"
            }
          }
        },
        "attributes": {
          "type": "nested",
          "properties": {
            "id": {
              "type": "keyword",
              "eager_global_ordinals": true
            },
            "label": {
              "type": "keyword",
              "eager_global_ordinals": true
            },
            "value": {
              "type": "keyword",
              "eager_global_ordinals": true
            },
            "inturnType": {
              "type": "keyword",
              "eager_global_ordinals": true
            },
            "position": {
              "type": "integer"
            }
          }
        },
        "variants": {
          "type": "nested",
          "properties": {
            "uuid": {
              "type": "keyword"
            },
            "name": {
              "type": "text",
              "fielddata": true,
			  "analyzer": "lowercase-analyzer"
            },
            "type": {
              "type": "keyword"
            },
            "upc": {
              "type": "keyword"
            },
            "position": {
              "type": "integer"
            },
			"productID": {
              "type": "keyword"
            },
            "markets": {
					"type": "nested",
					"properties": {
					"uuid": {
						"type": "keyword",
								"eager_global_ordinals": true
							},
							"name": {
								"type": "keyword",
								"eager_global_ordinals": true
							},
							"currency": {
								"type": "keyword",
								"eager_global_ordinals": true
							},
							"currencyCode": {
								"type": "keyword",
									"eager_global_ordinals": true
							},
							"cost": {
								"type": "float"
							},
							"wholesale": {
								"type": "float"
							},
							"retail": {
								"type": "float"
							},
							"suggested": {
								"type": "float"
							}
							}
						},
            "images": {
              "type": "nested",
              "properties": {
                "url": {
                  "type": "keyword"
                },
                "position": {
                  "type": "integer"
                },
                "colorMatch": {
                  "type": "boolean"
                }
              }
            },
            "options": {
              "type": "nested",
              "properties": {
                "valueUUID": {
                  "type": "keyword"
                },
                "name": {
                  "type": "keyword"
                },
                "value": {
                  "type": "keyword"
                },
				"position": {
                  "type": "integer"
            	}
              }
            }
          }
        },
        "packages": {
          "type": "nested",
          "properties": {
            "uuid": {
              "type": "keyword"
            },
            "name": {
              "type": "text",
              "fielddata": true,
         	  "analyzer": "lowercase-analyzer"
            },
            "type": {
              "type": "keyword"
            },
            "variants": {
              "type": "nested",
              "properties": {
                "uuid": {
                  "type": "keyword"
                },
                "ratio": {
                  "type": "integer"
                },
                "position": {
                  "type": "integer"
                }
              }
            },
            "groups": {
              "type": "nested",
              "properties": {
                "uuid": {
                  "type": "keyword"
                },
                "name": {
                  "type": "text",
                  "fielddata": true,
          		  "analyzer": "lowercase-analyzer"
                },
                "type": {
                  "type": "keyword"
                },
                "status": {
                  "type": "keyword"
                },
                "path": {
                  "type": "text",
                  "analyzer": "path-analyzer",
                  "search_analyzer": "keyword",
                  "fielddata": true
                },
                "total": {
                  "type": "integer"
                },
                "available": {
                  "type": "integer"
                },
                "unavailable": {
                  "type": "integer"
                },
                "marketId": {
                  "type": "keyword"
                },
                "marketName": {
                  "type": "keyword"
                },
                "currencyCode": {
                  "type": "keyword"
                },
                "isLocked": {
                  "type": "boolean"
                }
              }
            }
          }
        },
        "priceSummaries": {
          "type": "nested",
          "properties": {
            "groupId": {
              "type": "keyword"
            },
            "packageType": {
              "type": "keyword"
            },
            "marketId": {
              "type": "keyword"
            },
            "marketName": {
              "type": "keyword"
            },
            "currencyCode": {
              "type": "keyword"
            },
            "cost": {
              "type": "float"
            },
            "wholesale": {
              "type": "float"
            },
			"retail": {
                  "type": "float"
			},
            "suggested": {
              "type": "float"
            }
          }
        },
        "averagePriceSummaries": {
          "type": "nested",
          "properties": {
            "marketId": {
              "type": "keyword"
            },
            "marketName": {
              "type": "keyword"
            },
            "currencyCode": {
              "type": "keyword"
            },
            "cost": {
              "type": "float"
            },
            "wholesale": {
              "type": "float"
            },
            "retail": {
                  "type": "float"
            },
            "suggested": {
              "type": "float"
            }
          }
        },
        "quantitySummaries": {
          "type": "nested",
          "properties": {
            "groupId": {
              "type": "keyword"
            },
            "groupType": {
              "type": "keyword"
            },
            "packageType": {
              "type": "keyword"
            },
            "marketId": {
              "type": "keyword"
            },
            "marketName": {
              "type": "keyword"
            },
            "currencyCode": {
              "type": "keyword"
            },
            "totalQuantity": {
              "type": "integer"
            },
            "availableQuantity": {
              "type": "integer"
            }
          }
        },
        "sizeSummaries": {
          "type": "nested",
          "properties": {
            "groupId": {
              "type": "keyword"
            },
            "packageType": {
              "type": "keyword"
            },
            "minAvailablePerSize": {
              "type": "integer"
            },
            "maxAvailablePerSize": {
              "type": "integer"
            },
            "sizePerGroupSummaries": {
              "type": "nested",
              "properties": {
                "value": {
                  "type": "keyword"
                },
				"valueUUID": {
                  "type": "keyword"
                },
                "totalQuantity": {
                  "type": "integer"
                },
				"position": {
                  "type": "integer"
                },
                "availableQuantity": {
                  "type": "integer"
                }
              }
            }
          }
        }
      }
    }
  }
}`

// ProductIdx represents single product used by import flow mechanism
type ProductIdx struct {
	UUID                  *string                 `json:"uuid,omitempty"`
	CompanyUUID           *string                 `json:"companyUUID,omitempty"`
	Name                  *string                 `json:"name,omitempty"`
	Description           *string                 `json:"description,omitempty"`
	SKU                   *string                 `json:"sku,omitempty"`
	CreatedAt             *int64                  `json:"createdAt,omitempty"`
	CreatedByUUID         *string                 `json:"createdByUUID,omitempty"`
	Options               []OptionIdx             `json:"options,omitempty"`
	Attributes            []AttributeIdx          `json:"attributes,omitempty"`
	Variants              []VariantIdx            `json:"variants,omitempty"`
	Packages              []PackageIdx            `json:"packages,omitempty"`
	ImageUUIDs            []string                `json:"-"`                               // technical field used for collecting the images ids for requesting urld from documentservice
	PriceSummaries        []PriceSummaries        `json:"priceSummaries,omitempty"`        // field enriched on side of searchservice with summaries calculation before loading to ES
	AveragePriceSummaries []AveragePriceSummaries `json:"averagePriceSummaries,omitempty"` // field enriched on side of searchservice with summaries calculation before loading to ES
	QuantitySummaries     []QuantitySummaries     `json:"quantitySummaries,omitempty"`     // field enriched on side of searchservice with summaries calculation before loading to ES
	SizeSummaries         []SizeSummaries         `json:"sizeSummaries,omitempty"`         // field enriched on side of searchservice with summaries calculation before loading to ES
}

// MarketIdx entity compatible with proto structure in go-agreements
type MarketIdx struct {
	UUID           *string  `json:"uuid,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Currency       *string  `json:"currency,omitempty"`
	CountryCode    *string  `json:"countryCode,omitempty"`
	CurrencyCode   *string  `json:"currencyCode,omitempty"`
	WholesalePrice *float64 `json:"wholesale,omitempty"`
	RetailPrice    *float64 `json:"retail,omitempty"`
	CostPrice      *float64 `json:"cost,omitempty"`
	SuggestedPrice *float64 `json:"suggested,omitempty"`
}

// PriceIdx entity compatible with proto structure in go-agreements
type PriceIdx struct {
	Type  *string  `json:"type,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

// OptionIdx entity compatible with proto structure in go-agreements
type OptionIdx struct {
	UUID       *string `json:"uuid,omitempty"`
	Name       *string `json:"name,omitempty"`
	Value      *string `json:"value,omitempty"`
	ValueUUID  *string `json:"valueUUID,omitempty"`
	InturnType *string `json:"inturnType,omitempty"`
	Position   *int64  `json:"position,omitempty"`
}

// AttributeIdx entity compatible with proto structure in go-agreements
type AttributeIdx struct {
	UUID       *string `json:"id,omitempty"`
	Label      *string `json:"label,omitempty"`
	Type       *string `json:"type,omitempty"`
	Value      *string `json:"value,omitempty"`
	InturnType *string `json:"inturnType,omitempty"`
	Position   *int64  `json:"position,omitempty"`
}

// ImageIdx images for the variants
type ImageIdx struct {
	Position   *int64  `json:"position,omitempty"`
	UUID       *string `json:"-"`
	URL        *string `json:"url,omitempty"`
	ColorMatch *bool   `json:"colorMatch,omitempty"`
}

// VariantIdx variants for the product
type VariantIdx struct {
	UUID      *string     `json:"uuid,omitempty"`
	Name      *string     `json:"name,omitempty"`
	Type      *string     `json:"type,omitempty"`
	Position  *int64      `json:"position,omitempty"`
	ProductID *string     `json:"productID,omitempty"`
	Markets   []MarketIdx `json:"markets,omitempty"`
	Images    []ImageIdx  `json:"images,omitempty"`
	Options   []OptionIdx `json:"options,omitempty"`
}

// PackageIdx packages for the product
type PackageIdx struct {
	UUID     *string             `json:"uuid,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Variants []PackageVariantIdx `json:"variants,omitempty"`
	Groups   []GroupIdx          `json:"groups,omitempty"`
}

// PackageVariantIdx variant for the packages inside the product
type PackageVariantIdx struct {
	UUID     *string `json:"uuid" db:"uuid,omitempty"`
	Ratio    *int64  `json:"ratio" db:"ratio,omitempty"`
	Position *int64  `json:"position" db:"position,omitempty"`
}

// GroupIdx group for the product
type GroupIdx struct {
	UUID         *string `json:"uuid,omitempty"`
	Name         *string `json:"name,omitempty"`
	Status       *string `json:"status,omitempty"`
	Type         *string `json:"type,omitempty"`
	Path         *string `json:"path,omitempty"`
	Total        *int64  `json:"total,omitempty"`
	Available    *int64  `json:"available,omitempty"`
	Unavailable  *int64  `json:"unavailable,omitempty"`
	MarketID     *string `json:"marketId,omitempty"`
	MarketName   *string `json:"marketName,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
	IsLocked     *bool   `json:"isLocked,omitempty"`
}

// PriceSummaries prices distribution for the product
type PriceSummaries struct {
	GroupID      *string  `json:"groupId,omitempty"`
	PackageType  *string  `json:"packageType,omitempty"`
	MarketID     *string  `json:"marketId,omitempty"`
	MarketName   *string  `json:"marketName,omitempty"`
	CurrencyCode *string  `json:"currencyCode,omitempty"`
	Wholesale    *float64 `json:"wholesale,omitempty"`
	Retail       *float64 `json:"retail,omitempty"`
	Cost         *float64 `json:"cost,omitempty"`
	Suggested    *float64 `json:"suggested,omitempty"`
}

// AveragePriceSummaries sum total of variant prices divided by the number of variants for the product for a given market
type AveragePriceSummaries struct {
	MarketID     *string  `json:"marketId,omitempty"`
	MarketName   *string  `json:"marketName,omitempty"`
	CurrencyCode *string  `json:"currencyCode,omitempty"`
	Wholesale    *float64 `json:"wholesale,omitempty"`
	Retail       *float64 `json:"retail,omitempty"`
	Cost         *float64 `json:"cost,omitempty"`
	Suggested    *float64 `json:"suggested,omitempty"`
}

// QuantitySummaries quantity distribution for the product
type QuantitySummaries struct {
	GroupID           *string `json:"groupId,omitempty"`
	GroupType         *string `json:"groupType,omitempty"`
	PackageType       *string `json:"packageType,omitempty"`
	TotalQuantity     *int64  `json:"totalQuantity,omitempty"`
	AvailableQuantity *int64  `json:"availableQuantity,omitempty"`
}

// SizePerGroupSummaries size distribution for the single group/package type
type SizePerGroupSummaries struct {
	ValueUUID         *string `json:"valueUUID,omitempty"`
	Value             *string `json:"value,omitempty"`
	TotalQuantity     *int64  `json:"totalQuantity,omitempty"`
	AvailableQuantity *int64  `json:"availableQuantity,omitempty"`
	Position          *int64  `json:"position,omitempty"`
}

// SizeSummaries size distribution for the product
type SizeSummaries struct {
	GroupID               *string                 `json:"groupId,omitempty"`
	PackageType           *string                 `json:"packageType,omitempty"`
	MinAvailablePerSize   *int64                  `json:"minAvailablePerSize,omitempty"`
	MaxAvailablePerSize   *int64                  `json:"maxAvailablePerSize,omitempty"`
	SizePerGroupSummaries []SizePerGroupSummaries `json:"sizePerGroupSummaries,omitempty"`
}
