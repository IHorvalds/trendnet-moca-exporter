package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/PaesslerAG/jsonpath"
)

type linkStatus int

const (
	UP   linkStatus = 1
	DOWN linkStatus = 0
)

func getMOCALinkStatus(adapter *adapter) (linkStatus, error) {
	url := fmt.Sprintf("http://%s/ms/0/0x15", adapter.MocaAdapterAddress)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer([]byte(`{"data":[]}`)))
	if err != nil {
		return DOWN, err
	}

	req.SetBasicAuth(adapter.MocaUser, adapter.MocaPass)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return DOWN, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return DOWN, fmt.Errorf("%s returned non-200 status code %d", url, resp.StatusCode)
	}

	var json_data map[string]interface{}
	if json.NewDecoder(resp.Body).Decode(&json_data) != nil {
		return DOWN, fmt.Errorf("invalid JSON response: '%s'", resp.Body)
	}

	mocaStatus, err := jsonpath.Get("$.data[5]", json_data)
	if err != nil {
		return DOWN, fmt.Errorf("$.data[5] not fount in JSON response: '%s'", json_data)
	}

	str, ok := mocaStatus.(string)
	if !ok {
		return DOWN, fmt.Errorf("expected string value, got %s", reflect.TypeOf(mocaStatus).String())
	}

	val, err := strconv.ParseInt(str, 0, 16)
	if err != nil {
		return DOWN, err
	}

	if val == 1 {
		return UP, nil
	}

	return DOWN, nil
}
