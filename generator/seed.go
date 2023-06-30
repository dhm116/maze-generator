package generator

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

func EncodeMapSeed(seed int64, size int) string {
	value := fmt.Sprintf("%d:%v", size, seed)
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func DecodeMapSeed(value string) (seed int64, size int) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		fmt.Println("Error decoding map seed:", err)
	} else {
		parts := strings.Split(string(decoded), ":")
		size, _ = strconv.Atoi(parts[0])
		seed, _ = strconv.ParseInt(parts[1], 10, 64)
	}
	return seed, size
}
