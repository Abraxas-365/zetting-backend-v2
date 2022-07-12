package application

import "bytes"

func (app *application) ViewImage(key string) (*bytes.Buffer, error) {
	buf, err := app.bucket.Get(key)
	if err != nil {
		return nil, err
	}
	return buf, nil

}
