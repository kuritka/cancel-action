package cancel

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/kuritka/cancel-action/internal/impl"
)

type command int

const (
	deleteWorkflow command = iota
	cancelWorkflow
	readCommit
)

type requestFactory struct {
	opts impl.ActionOpts
	ctx  context.Context
}

func newRequestFactory(ctx context.Context, opts impl.ActionOpts) *requestFactory {
	return &requestFactory{ctx: ctx, opts: opts}
}

func (f *requestFactory) getImpl(cmd command) func() (*http.Request, error) {
	switch cmd {
	case cancelWorkflow:
		return f.cancelWorkflow
	case deleteWorkflow:
		return f.deleteWorkflow
	case readCommit:
		return f.readCommit
	}
	return func() (*http.Request, error) { return nil, fmt.Errorf("not implemented") }
}

func (f *requestFactory) deleteWorkflow() (req *http.Request, err error) {
	json := new(bytes.Buffer)
	var url = fmt.Sprintf("https://api.github.com/repos/%s/actions/runs/%s", f.opts.GitHub.Repository, f.opts.GitHub.RunID)
	req, err = http.NewRequestWithContext(f.ctx, http.MethodDelete, url, json)
	logger.Debug().Msgf("hitting URL: %s", url)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "token "+f.opts.GitHub.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	return
}

func (f *requestFactory) cancelWorkflow() (req *http.Request, err error) {
	json := new(bytes.Buffer)
	var url = fmt.Sprintf("https://api.github.com/repos/%s/actions/runs/%s/cancel", f.opts.GitHub.Repository, f.opts.GitHub.RunID)
	logger.Debug().Msgf("hitting URL: %s", url)
	req, err = http.NewRequestWithContext(f.ctx, http.MethodPost, url, json)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+f.opts.GitHub.Token)
	req.Header.Set("User-Agent", "actions/cancel-action")
	return
}

func (f *requestFactory) readCommit() (req *http.Request, err error) {
	json := new(bytes.Buffer)
	var url = fmt.Sprintf("https://api.github.com/repos/%s/commits/%s", f.opts.GitHub.Repository, f.opts.GitHub.CommitSHA)
	logger.Debug().Msgf("hitting URL: %s", url)
	req, err = http.NewRequestWithContext(f.ctx, http.MethodGet, url, json)
	if err != nil {
		return
	}
	req.Header.Set("Accept","application/vnd.github.v3+json")
	return
}
