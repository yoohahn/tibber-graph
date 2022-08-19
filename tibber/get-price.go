package tibber

import (
	"context"
	"encoding/json"

	"github.com/machinebox/graphql"
)

// Returns the body as first paramater and if the request was succes or not
func GetPrice() (string, bool) {
	hasError := false
	client := graphql.NewClient("https://api.tibber.com/v1-beta/gql")
	req := graphql.NewRequest(`
        {
            viewer {
                homes {
                    currentSubscription {
                        priceInfo {
                            current {
                                total
                                startsAt
                            }
                            today {
                                total
                                startsAt
                            }
                            tomorrow {
                                total
                                startsAt
                            }
                        }
                    }
                }
            }
        }
    `)
	req.Header.Set("authorization", "Bearer "+GetMyApiKey())
	req.Header.Set("Content-Type", "application/json")

	var graphqlResponse interface{}
	if err := client.Run(context.Background(), req, &graphqlResponse); err != nil {
		panic(err)
	}

	responseAsJson, err := json.Marshal(graphqlResponse)
	result := ""

	if err != nil {
		result = err.Error()
		hasError = true
	} else {
		result = string(responseAsJson)
	}
	return result, hasError
}
