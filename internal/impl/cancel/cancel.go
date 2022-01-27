package cancel

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	w "github.com/logrusorgru/aurora"

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

func (c *Cancel) Run() (err error) {
	return c.action(cancelWorkflow, "🟨 Delete workflow")
	// err = c.action(cancelWorkflow, "🟨 Cancel workflow")
	// if err != nil {
	// 	 return err
	// }
	// return c.action(deleteWorkflow, "🟨 Delete workflow")
}

func (c *Cancel) action(cmd command, message string) error {
	logger.Info().Msg(message)
	status, err := request(c.f.getImpl(cmd))
	if err != nil {
		logger.Err(err).Msgf("request error (%v)", status)
		return err
	}
	logger.Info().Msgf(" - %v", w.BrightYellow(status))
	return nil
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
