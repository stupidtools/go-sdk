// Package citesting provides a record/replay harness for CI tests.
package citesting

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type CacheMode string

const (
	Locked   CacheMode = "Locked"
	Unlocked CacheMode = "Unlocked"
)

const timestampKey = "_fixedTimestamp"

type CiTestingHarness struct {
	proxyBase            string
	targetServiceURL     string
	cacheMode            CacheMode
	testName             string
	generatedIds         map[string]string
	fixedTimestampMillis int64

	language        string
	category        string
	subcategory     string
	currentFunction string
	varPayload      []string
	varQuery        []string
}

func New(targetServiceURL string, mode CacheMode, testName string) *CiTestingHarness {
	proxyHost := os.Getenv("PROXY_ENDPOINT")
	if proxyHost == "" {
		panic("PROXY_ENDPOINT environment variable is not set. " +
			"Set it to the proxy hostname (e.g. PROXY_ENDPOINT=\"your-proxy-host.example.com\").")
	}
	h := &CiTestingHarness{
		proxyBase:        "https://" + proxyHost + "/proxy",
		targetServiceURL: targetServiceURL,
		cacheMode:        mode,
		testName:         testName,
		language:         "go",
	}
	h.setDefaultVariableFields(targetServiceURL)

	if mode == Locked {
		h.generatedIds = h.loadIds(true)
	} else {
		h.generatedIds = h.loadIds(false)
		h.generatedIds[timestampKey] = fmt.Sprintf("%d", time.Now().UnixMilli())
		h.writeIds()
	}
	fmt.Sscanf(h.generatedIds[timestampKey], "%d", &h.fixedTimestampMillis)

	os.Setenv("SDK_TESTING_TIMESTAMP", fmt.Sprintf("%d", h.fixedTimestampMillis))

	time.Local = time.UTC

	return h
}

func (h *CiTestingHarness) SetFunction(functionPath string) {
	if functionPath == "" {
		h.category, h.subcategory, h.currentFunction = "", "", ""
		return
	}
	parts := strings.SplitN(functionPath, "|", 3)
	if len(parts) == 3 {
		h.category, h.subcategory, h.currentFunction = parts[0], parts[1], parts[2]
	} else {
		h.category, h.subcategory, h.currentFunction = "", "", functionPath
	}
}

func (h *CiTestingHarness) SetLanguage(language string) {
	h.language = language
}

func (h *CiTestingHarness) SetVariableFields(payload, query []string) {
	h.varPayload = payload
	h.varQuery = query
}

func (h *CiTestingHarness) GetTestingURL() string {
	cacheReturns := "false"
	if h.cacheMode == Locked {
		cacheReturns = "true"
	}

	targetHost := h.targetServiceURL
	for _, prefix := range []string{"https://", "http://"} {
		if strings.HasPrefix(targetHost, prefix) {
			targetHost = "https:/" + targetHost[len(prefix):]
			break
		}
	}

	args := "cacheReturns:" + cacheReturns

	var varsParts []string
	for _, f := range h.varPayload {
		varsParts = append(varsParts, "payload."+f)
	}
	for _, f := range h.varQuery {
		varsParts = append(varsParts, "query."+f)
	}
	if len(varsParts) > 0 {
		args += ",vars:" + strings.Join(varsParts, "|")
	}

	if h.language != "" && h.category != "" && h.subcategory != "" && h.currentFunction != "" {
		args += ",language:" + encode(h.language) +
			",category:" + encode(h.category) +
			",subcategory:" + encode(h.subcategory) +
			",function:" + encode(h.currentFunction)
	}

	return h.proxyBase + "/(" + args + ")/" + targetHost
}

func (h *CiTestingHarness) GetCurrentTime() time.Time {
	return time.UnixMilli(h.fixedTimestampMillis).UTC()
}

func (h *CiTestingHarness) GenerateRandomID(key string) string {
	if existing, ok := h.generatedIds[key]; ok {
		return existing
	}
	if h.cacheMode == Locked {
		panic("Harness is locked but no cached ID found for key: " + key +
			". Run tests in Unlocked mode first to generate IDs.")
	}
	id := fmt.Sprintf("%d", rand.Intn(99999998-1000000)+1000000)
	h.generatedIds[key] = id
	h.writeIds()
	return id
}

func (h *CiTestingHarness) setDefaultVariableFields(_ string) {
	h.varPayload = nil
	h.varQuery = nil
}

func encode(value string) string {
	return url.PathEscape(value) // space -> %20
}

func (h *CiTestingHarness) resourceDir() string {
	return filepath.Join("resources", "citesting")
}

func (h *CiTestingHarness) writeIds() {
	dir := h.resourceDir()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(h.generatedIds, "", "    ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(filepath.Join(dir, h.testName+".json"), data, 0o644); err != nil {
		panic(err)
	}
}

func (h *CiTestingHarness) loadIds(required bool) map[string]string {
	path := filepath.Join(h.resourceDir(), h.testName+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		if required {
			panic("Harness is locked but ID file not found: " + path +
				". Run tests in Unlocked mode first to generate the file.")
		}
		return map[string]string{}
	}
	result := map[string]string{}
	if err := json.Unmarshal(data, &result); err != nil {
		panic(err)
	}
	return result
}
