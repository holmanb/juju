// Copyright 2014, 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package model

import (
	"fmt"

	"github.com/juju/cmd/v3"
	"github.com/juju/errors"
	"github.com/juju/names/v4"

	"github.com/juju/juju/api/client/machinemanager"
	jujucmd "github.com/juju/juju/cmd"
	"github.com/juju/juju/cmd/juju/block"
	"github.com/juju/juju/cmd/modelcmd"
	"github.com/juju/juju/rpc/params"
)

func NewRetryProvisioningCommand() cmd.Command {
	return modelcmd.Wrap(&retryProvisioningCommand{})
}

// retryProvisioningCommand updates machines' error status to tell
// the provisoner that it should try to re-provision the machine.
type retryProvisioningCommand struct {
	modelcmd.ModelCommandBase
	modelcmd.IAASOnlyCommand
	Machines []names.MachineTag
	api      RetryProvisioningAPI
}

// RetryProvisioningAPI defines methods on the client API
// that the retry-provisioning command calls.
type RetryProvisioningAPI interface {
	Close() error
	RetryProvisioning(machines ...names.MachineTag) ([]params.ErrorResult, error)
}

func (c *retryProvisioningCommand) Info() *cmd.Info {
	return jujucmd.Info(&cmd.Info{
		Name:    "retry-provisioning",
		Args:    "<machine> [...]",
		Purpose: "Retries provisioning for failed machines.",
	})
}

func (c *retryProvisioningCommand) Init(args []string) error {
	if len(args) == 0 {
		return errors.Errorf("no machine specified")
	}
	c.Machines = make([]names.MachineTag, len(args))
	for i, arg := range args {
		if !names.IsValidMachine(arg) {
			return errors.Errorf("invalid machine %q", arg)
		}
		if names.IsContainerMachine(arg) {
			return errors.Errorf("invalid machine %q retry-provisioning does not support containers", arg)
		}
		c.Machines[i] = names.NewMachineTag(arg)
	}
	return nil
}

func (c *retryProvisioningCommand) getAPI() (RetryProvisioningAPI, error) {
	if c.api != nil {
		return c.api, nil
	}
	root, err := c.NewAPIRoot()
	if err != nil {
		return nil, errors.Trace(err)
	}
	client := machinemanager.NewClient(root)
	return client, nil
}

func (c *retryProvisioningCommand) Run(context *cmd.Context) error {
	client, err := c.getAPI()
	if err != nil {
		return err
	}
	defer client.Close()

	results, err := client.RetryProvisioning(c.Machines...)
	if err != nil {
		return block.ProcessBlockedError(err, block.BlockChange)
	}
	for _, result := range results {
		if result.Error != nil {
			fmt.Fprintf(context.Stderr, "%v\n", result.Error)
		}
	}
	return nil
}
