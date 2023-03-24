package parser

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/chain4energy/juno/v4/logging"
	"github.com/chain4energy/juno/v4/node"

	"github.com/chain4energy/juno/v4/database"
	"github.com/chain4energy/juno/v4/modules"
)

// Context represents the context that is shared among different workers
type Context struct {
	EncodingConfig *params.EncodingConfig
	Node           node.Node
	Database       database.Database
	Logger         logging.Logger
	Modules        []modules.Module
}

// NewContext builds a new Context instance
func NewContext(
	encodingConfig *params.EncodingConfig, proxy node.Node, db database.Database,
	logger logging.Logger, modules []modules.Module,
) *Context {
	return &Context{
		EncodingConfig: encodingConfig,
		Node:           proxy,
		Database:       db,
		Modules:        modules,
		Logger:         logger,
	}
}
