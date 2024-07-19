package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/f1bonacc1/process-compose/src/types"
	"github.com/rs/zerolog/log"
)

func (p *PcClient) addProcess(processConfig types.ProcessConfig) error {
	url := fmt.Sprintf("http://%s/process/add", p.address)

	b, err := json.Marshal(processConfig)
	if err != nil {
		log.Error().Msgf("failed to marshal add process body: %v", err)
		return err
	}

	r := bytes.NewReader(b)

	resp, err := p.client.Post(url, "application/json", r)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	var respErr pcError
	if err = json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
		log.Error().Msgf("failed to decode add process response: %v", err)
		return err
	}
	return fmt.Errorf(respErr.Error)
}
