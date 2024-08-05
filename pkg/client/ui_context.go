package client

import "github.com/bndrmrtn/fate_fighter/pkg/i18n"

type Context struct {
	DarkMode bool
	i18n.ITranslator
}
