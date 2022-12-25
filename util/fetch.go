package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func getAOCToken() string {
	homeDirName, _ := os.UserHomeDir()
	tokenFilePath := filepath.Join(homeDirName, ".config", "aocd", "token")
	f, _ := os.ReadFile(tokenFilePath)
	return strings.TrimSpace(string(f))
}

func FetchInputData(day int) string {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day), nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: getAOCToken()})

	client := &http.Client{}
	resp, respError := client.Do(req)
	if respError != nil {
		return ""
	}
	defer resp.Body.Close()

	text, _ := io.ReadAll(resp.Body)

	return strings.TrimRightFunc(string(text), unicode.IsSpace)
}
