package service

import (
	"backpack-trade-bot/utils"
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Utility function to decode a base64 encoded Ed25519 private key and verify it against the provided public key.
func base64ToEd25519PrivateKey(privateKeyBase64 string) (ed25519.PrivateKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("error decoding private key from base64: %v", err)
	}
	if len(keyBytes) == ed25519.PrivateKeySize {
		return ed25519.PrivateKey(keyBytes), nil
	}
	if len(keyBytes) == ed25519.SeedSize {
		// Convert the seed (32 bytes) to a private key (64 bytes).
		return ed25519.NewKeyFromSeed(keyBytes[:32]), nil
	}
	return nil, errors.New("invalid private key length")
}

// Utility function to decode a base64 encoded Ed25519 public key.
func base64ToEd25519PublicKey(publicKeyBase64 string) (ed25519.PublicKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("error decoding public key from base64: %v", err)
	}
	if len(keyBytes) != ed25519.PublicKeySize {
		return nil, errors.New("invalid public key length")
	}
	return ed25519.PublicKey(keyBytes), nil
}

// BackpackClient structure to hold client configuration.
type BackpackClient struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  string
	BaseURL    string
}

// NewBackpackClient creates a new BackpackClient instance after validating the provided keys.
func NewBackpackClient(config *Config) (*BackpackClient, error) {
	privateKey, err := base64ToEd25519PrivateKey(config.APIConfig.APISecret)
	if err != nil {
		return nil, err
	}

	publicKey, err := base64ToEd25519PublicKey(config.APIConfig.APIKey)
	if err != nil {
		return nil, err
	}

	// Verify the public key matches the private key
	if !ed25519.PublicKey(privateKey.Public().(ed25519.PublicKey)).Equal(publicKey) {
		return nil, errors.New("public key does not match the provided private key")
	}

	return &BackpackClient{
		PrivateKey: privateKey,
		PublicKey:  config.APIConfig.APIKey,
		BaseURL:    "https://api.backpack.exchange/",
	}, nil
}

func (c *BackpackClient) MakePublicAPIRequest(method, endpoint string, params map[string]string) ([]byte, error) {
	client := &http.Client{}

	// This code block is part of the MakeAPIRequest method

	var req *http.Request
	var err error

	urlValues := url.Values{}
	for key, value := range params {
		urlValues.Set(key, value)
	}

	if method == "GET" {
		// For GET requests, append the encoded parameters to the request URL
		reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, endpoint, urlValues.Encode())
		req, err = http.NewRequest(method, reqURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create new request: %w", err)
		}
	} else {
		// For POST, PUT, and other methods, encode the parameters as form data
		req, err = http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, endpoint), strings.NewReader(urlValues.Encode()))
		if err != nil {
			return nil, fmt.Errorf("failed to create new request: %w", err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	if err != nil {
		return nil, err
	}

	// Set headers including the signature
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *BackpackClient) MakeAuthenticatedAPIRequest(method, endpoint, instruction string, params map[string]string) ([]byte, error) {
	// Prepare the request
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	signature, err := c.SignMessage(method, endpoint, params, timestamp, instruction, DefaultTimeoutMs)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	// This code block is part of the MakeAPIRequest method

	var req *http.Request

	urlValues := url.Values{}
	for key, value := range params {
		urlValues.Set(key, value)
	}

	if method == "GET" {
		// For GET requests, append the encoded parameters to the request URL
		reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, endpoint, urlValues.Encode())
		req, err = http.NewRequest(method, reqURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create new request: %w", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		var jsonData []byte
		// For POST, PUT, and other methods, encode the parameters as JSON
		jsonData, err = json.Marshal(params) // Assuming params is a map[string]string or similar
		if err != nil {
			return nil, fmt.Errorf("failed to marshal params to JSON: %w", err)
		}
		req, err = http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, endpoint), bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, fmt.Errorf("failed to create new request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if err != nil {
		return nil, err
	}

	// Set headers including the signature
	req.Header.Set("X-API-KEY", c.PublicKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-TIMESTAMP", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-WINDOW", string(rune(DefaultTimeoutMs)))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *BackpackClient) MakeAuthenticatedAPIInterfaceRequest(method, endpoint, instruction string, params map[string]interface{}) ([]byte, error) {
	// Prepare the request
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	strParams, err := utils.ConvertMapToStringMap(params)
	if err != nil {
		return nil, err
	}
	signature, err := c.SignInterfaceMessage(method, endpoint, strParams, timestamp, instruction, DefaultTimeoutMs)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	// This code block is part of the MakeAPIRequest method

	var req *http.Request

	var jsonData []byte
	// For POST, PUT, and other methods, encode the parameters as JSON
	jsonData, err = json.Marshal(params) // Assuming params is a map[string]string or similar
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params to JSON: %w", err)
	}
	req, err = http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// Set headers including the signature
	req.Header.Set("X-API-KEY", c.PublicKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("X-TIMESTAMP", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-WINDOW", string(rune(DefaultTimeoutMs)))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
