package service

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

const DefaultTimeoutMs = 5000

func (c *BackpackClient) SignMessage(method, endpoint string, params map[string]string, timestamp int64, instruction string, window int64) (string, error) {
	// Ensure window defaults to DEFAULT_TIMEOUT_MS if not specified
	if window == 0 {
		window = DefaultTimeoutMs
	}

	// Sort the parameter keys alphabetically
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Build the parameter string
	var paramString strings.Builder
	for _, key := range keys {
		paramString.WriteString(fmt.Sprintf("%s=%s&", key, url.QueryEscape(params[key])))
	}

	// Construct the header string
	headerInfo := fmt.Sprintf("timestamp=%d&window=%d", timestamp, window)

	// Construct the message to sign
	messageToSign := fmt.Sprintf("instruction=%s&%s%s", instruction, paramString.String(), headerInfo)

	// Sign the message
	signature, err := c.PrivateKey.Sign(rand.Reader, []byte(messageToSign), crypto.Hash(0))
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %w", err)
	}

	// Encode the signature in base64 and return
	return base64.StdEncoding.EncodeToString(signature), nil
}

func (c *BackpackClient) SignInterfaceMessage(method, endpoint string, params map[string]string, timestamp int64, instruction string, window int64) (string, error) {
	// Ensure window defaults to DEFAULT_TIMEOUT_MS if not specified
	if window == 0 {
		window = DefaultTimeoutMs
	}

	// Sort the parameter keys alphabetically
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Build the parameter string
	var paramString strings.Builder
	for _, key := range keys {
		paramString.WriteString(fmt.Sprintf("%s=%s&", key, url.QueryEscape(params[key])))
	}

	// Construct the header string
	headerInfo := fmt.Sprintf("timestamp=%d&window=%d", timestamp, window)

	// Construct the message to sign
	messageToSign := fmt.Sprintf("instruction=%s&%s%s", instruction, paramString.String(), headerInfo)

	// Sign the message
	signature, err := c.PrivateKey.Sign(rand.Reader, []byte(messageToSign), crypto.Hash(0))
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %w", err)
	}

	// Encode the signature in base64 and return
	return base64.StdEncoding.EncodeToString(signature), nil
}
