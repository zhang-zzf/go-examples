package google

import (
	"context"
	"encoding/json"
	"example/0340contexts/userip"
	"net/http"
)

type Result struct {
	Title, URL string
}
type Results []Result

func Search(ctx context.Context, query string) (Results, error) {
	url := "https://ajax.googleapis.com/ajax/services/search/web?v=1.0"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	// encode
	req.URL.RawQuery = q.Encode()
	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var data struct {
			ResponseData struct {
				Results []struct {
					TitleNoFormatting string
					URL               string
				}
			}
		}
		if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		for _, r := range data.ResponseData.Results {
			results = append(results, Result{
				Title: r.TitleNoFormatting,
				URL:   r.URL,
			})
		}
		return nil
	})
	return results, err
}

func httpDo(ctx context.Context, req *http.Request,
	respHandle func(*http.Response, error) error) error {
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() {
		resp, err := http.DefaultClient.Do(req)
		c <- respHandle(resp, err)
	}()
	select {
	case <-ctx.Done():
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}
