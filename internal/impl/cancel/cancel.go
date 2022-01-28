package cancel

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kuritka/cancel-action/internal/common/log"
	"github.com/kuritka/cancel-action/internal/impl"
)

var logger = log.Log

type Cancel struct {
	f *requestFactory
}

func NewCommand(o impl.ActionOpts) *Cancel {
	return &Cancel{
		f: newRequestFactory(context.TODO(), o),
	}
}

func (c *Cancel) Run() error {
	logger.Info().Msg("Cancelling....")

	status, err := request(c.f.getImpl(cancelWorkflow))
	if err != nil {
		logger.Err(err).Msgf("error during github request; code: %v", status)
		return err
	}
	logger.Info().Msgf("returned status code: %v", status)

	logger.Info().Msg("Deleting....")
	status, err = request(c.f.getImpl(deleteWorkflow))
	if err != nil {
		logger.Err(err).Msgf("error during github DELETE request; code: %v", status)
		return err
	}
	logger.Info().Msgf("returned status code: %v", status)

	return err
}

func (c *Cancel) String() string {
	return "CANCEL ACTION"
}

func request(f func() (*http.Request, error)) (status int, err error) {
	req, err := f()
	if err != nil {
		return 0, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, err
	}
	// nolint:errcheck
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 208 {
		return resp.StatusCode, fmt.Errorf(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	logger.Debug().Msgf("response body: %s", string(body))
	return resp.StatusCode, nil
}
