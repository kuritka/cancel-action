package cancel

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kuritka/cancel-action/internal/common/log"
	"github.com/kuritka/cancel-action/internal/impl"
)

var logger = log.Log

type Cancel struct {
	o impl.ActionOpts
}

func NewCommand(o impl.ActionOpts) *Cancel {
	return &Cancel{
		o: o,
	}
}

func (c *Cancel) Run() error {
	result, err := request(context.TODO(), c.o.GitHub)
	if err != nil {
		logger.Err(err).Msg("error during github request")
	}
	logger.Debug().Msg(result)
	return err
}

func (c *Cancel) String() string {
	return "CANCEL ACTION"
}

func request(ctx context.Context, gh impl.GitHub) (result string, err error) {
	json := new(bytes.Buffer)
	var url = fmt.Sprintf("https://api.github.com/repos/%s/actions/runs/%s/cancel", gh.Repository, gh.RunID)
	logger.Debug().Msgf("hitting URL: %s", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, json)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+gh.Token)
	req.Header.Set("User-Agent", "actions/cancel-action")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	// nolint:errcheck
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 208 {
		return "", fmt.Errorf(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
