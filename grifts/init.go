package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kagundajm/buffalo/jquery_uiautocomplete/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
