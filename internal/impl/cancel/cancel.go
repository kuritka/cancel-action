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
	var b []byte
	b, err = c.action(readCommit, "🟩 READ commit")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	_, err = c.action(cancelWorkflow, "🟨 CANCEL workflow")
	return err
}

func (c *Cancel) action(cmd command, message string) ([]byte, error) {
	logger.Info().Msg(message)
	body, status, err := request(c.f.getImpl(cmd))
	if err != nil {
		logger.Err(err).Msgf("request error (%v)", status)
		return body, err
	}
	logger.Info().Msgf(" - %v", w.BrightYellow(status))
	return body, nil
}

func (c *Cancel) String() string {
	return "CANCEL ACTION"
}

func request(f func() (*http.Request, error)) (b []byte, status int, err error) {
	req, err := f()
	if err != nil {
		return nil, 0, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	// nolint:errcheck
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 208 {
		return nil, resp.StatusCode, fmt.Errorf(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}
