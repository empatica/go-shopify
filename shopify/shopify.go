// Package shopify provides an easy-to-use API for making CRUD request to shopify.
package shopify

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// Shopify store struct which we use to wrap our request operations.
type Shopify struct {
	// Store domain-name
	store string
	// Store API key
	apiKey string
	// Store password
	pass string
}

const (
	domain = ".myshopify.com/admin"
)

// New Creates a New Shopify Store API object with the store, apiKey and pass of your store.
// Usage: shopify.New("mystore", "XXX","YYY")
func New(store, apiKey, pass string) Shopify {
	return Shopify{store: store, apiKey: apiKey, pass: pass}
}

// Request Creates a new Request to Shopify and returns the response as a map[string]interface{}.
// method: GET/POST/PUT - string
// url: target endpoint like "products" - string
// data: content to be sent with the request
// Usage: shopify.request("GET","products",nil)
func (shopify *Shopify) Request(method, endpoint string, data interface{}) ([]byte, []error) {
	jsonData, _ := getJSONBytesFromMap(data)
	targetURL := shopify.createTargetURL(endpoint)

	request := gorequest.New()
	request.Get(targetURL)

	if jsonData != nil && data != nil {
		request.Send(string(jsonData))
	}

	_, body, errs := request.End()

	return []byte(body), errs
}

// Get Makes a GET request to shopify with the given endpoint.
// Usage:
// shopify.Get("products/5.json")
// shopify.Get("products/5/variants.json")
func (shopify *Shopify) Get(endpoint string) ([]byte, []error) {
	return shopify.GetWithParameters(endpoint, nil)
}

// GetWithParameters Makes a GET request to shopify with the given endpoint and given parameters
func (shopify *Shopify) GetWithParameters(endpoint string, parameters map[string]string) ([]byte, []error) {
	targetURL := shopify.createTargetURLWithParameters(endpoint, parameters)
	request := gorequest.New()
	_, body, errs := request.Get(targetURL).End()

	return []byte(body), errs
}

// Post Makes a POST request to shopify with the given endpoint and data.
// Usage: shopify.Post("products", map[string]interface{} = product data map)
func (shopify *Shopify) Post(endpoint string, data interface{}) ([]byte, []error) {
	targetURL := shopify.createTargetURL(endpoint)
	jsonData, err := getJSONBytesFromMap(data)
	if err != nil {
		return nil, []error{err}
	}

	request := gorequest.New()
	request.Post(targetURL)
	if jsonData != nil && data != nil {
		request.Send(string(jsonData))
	}
	_, body, errs := request.End()

	return []byte(body), errs
}

// Put Makes a PUT request to shopify with the given endpoint and data.
// Usage: shopify.Put("products", map[string]interface{} = product data map)
func (shopify *Shopify) Put(endpoint string, data interface{}) ([]byte, []error) {
	targetURL := shopify.createTargetURL(endpoint)
	jsonData, err := getJSONBytesFromMap(data)
	if err != nil {
		return nil, []error{err}
	}

	request := gorequest.New()
	request.Put(targetURL)
	if jsonData != nil && data != nil {
		request.Send(string(jsonData))
	}
	_, body, errs := request.End()

	return []byte(body), errs
}

// Delete Makes a DELETE request to shopify with the given endpoint.
// Usage: shopify.Delete("products/5.json")
func (shopify *Shopify) Delete(endpoint string) ([]byte, []error) {
	targetURL := shopify.createTargetURL(endpoint)

	request := gorequest.New()
	_, body, errs := request.Delete(targetURL).End()

	return []byte(body), errs
}

// Creates target URL for making a Shopify Request to a given endpoint
func (shopify *Shopify) createTargetURL(endpoint string) string {
	return shopify.createTargetURLWithParameters(endpoint, nil)
}

// Creates target URL for making a Shopify Request to a given endpoint with the given parameters
func (shopify *Shopify) createTargetURLWithParameters(endpoint string, parameters map[string]string) string {
	var parametersString = ""
	if parameters != nil && len(parameters) > 0 {
		parametersString = "?"
		for k := range parameters {
			parametersString = fmt.Sprintf("%v%v=%v&", parametersString, k, parameters[k])
		}
	}
	return fmt.Sprintf("https://%s:%s@%s%s/%s.json%s", shopify.apiKey, shopify.pass, shopify.store, domain, endpoint, parametersString)
}
