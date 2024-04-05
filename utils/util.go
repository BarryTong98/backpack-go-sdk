package utils

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func ConvertMapToStringMap(originalMap map[string]interface{}) (map[string]string, error) {
	convertedMap := make(map[string]string)
	for key, value := range originalMap {
		switch v := value.(type) {
		case string:
			convertedMap[key] = v
		case int, int32, int64, float32, float64:
			convertedMap[key] = fmt.Sprintf("%v", v)
		case bool:
			convertedMap[key] = strconv.FormatBool(v)
		default:
			jsonValue, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("error marshaling value for key '%s': %w", key, err)
			}
			convertedMap[key] = string(jsonValue)
		}
	}
	return convertedMap, nil
}

// Base64ToEd25519PrivateKey Utility function to decode a base64 encoded Ed25519 private key and verify it against the provided public key.
func Base64ToEd25519PrivateKey(privateKeyBase64 string) (ed25519.PrivateKey, error) {
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

// Base64ToEd25519PublicKey Utility function to decode a base64 encoded Ed25519 public key.
func Base64ToEd25519PublicKey(publicKeyBase64 string) (ed25519.PublicKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("error decoding public key from base64: %v", err)
	}
	if len(keyBytes) != ed25519.PublicKeySize {
		return nil, errors.New("invalid public key length")
	}
	return ed25519.PublicKey(keyBytes), nil
}
