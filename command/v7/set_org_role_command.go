package v7

import (
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/cf/errors"
	"code.cloudfoundry.org/cli/command/translatableerror"
	"code.cloudfoundry.org/clock"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/v7/shared"
)

//go:generate counterfeiter . SetOrgRoleActor

type SetOrgRoleActor interface {
	CreateOrgRole(roleType constant.RoleType, orgGUID string, userNameOrGUID string, userOrigin string, isClient bool) (v7action.Warnings, error)
	GetOrganizationByName(name string) (v7action.Organization, v7action.Warnings, error)
	GetUser(username, origin string) (v7action.User, error)
}

type SetOrgRoleCommand struct {
	Args            flag.OrgRoleArgs `positional-args:"yes"`
	IsClient        bool             `long:"client" description:"Assign an org role to a client-id of a (non-user) service account"`
	Origin          string           `long:"origin" description:"Indicates the identity provider to be used for authentication"`
	usage           interface{}      `usage:"CF_NAME set-org-role USERNAME ORG ROLE\n   CF_NAME set-org-role USERNAME ORG ROLE [--client]\n   CF_NAME set-org-role USERNAME ORG ROLE [--origin ORIGIN]\n\nROLES:\n   OrgManager - Invite and manage users, select and change plans, and set spending limits\n   BillingManager - Create and manage the billing account and payment info\n   OrgAuditor - Read-only access to org info and reports"`
	relatedCommands interface{}      `related_commands:"org-users, set-space-role"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       SetOrgRoleActor
}

func (cmd *SetOrgRoleCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	sharedActor := sharedaction.NewActor(config)
	cmd.SharedActor = sharedActor

	ccClient, uaaClient, err := shared.GetNewClientsAndConnectToCF(config, ui, "")
	if err != nil {
		return err
	}
	cmd.Actor = v7action.NewActor(ccClient, config, sharedActor, uaaClient, clock.NewClock())
	return nil
}

func (cmd *SetOrgRoleCommand) Execute(args []string) error {
	err := cmd.validateFlags()
	if err != nil {
		return err
	}

	err = cmd.SharedActor.CheckTarget(false, false)
	if err != nil {
		return err
	}

	currentUser, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor("Assigning role {{.RoleType}} to user {{.TargetUserName}} in org {{.OrgName}} as {{.CurrentUserName}}...", map[string]interface{}{
		"RoleType":        cmd.Args.Role.Role,
		"TargetUserName":  cmd.Args.Username,
		"OrgName":         cmd.Args.Organization,
		"CurrentUserName": currentUser.Name,
	})

	roleType, err := convertRoleType(cmd.Args.Role)
	if err != nil {
		return err
	}

	org, warnings, err := cmd.Actor.GetOrganizationByName(cmd.Args.Organization)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	origin := cmd.Origin
	if cmd.Origin == "" {
		origin = constant.DefaultOriginUaa
	}

	warnings, err = cmd.Actor.CreateOrgRole(roleType, org.GUID, cmd.Args.Username, origin, cmd.IsClient)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		if _, ok := err.(ccerror.RoleAlreadyExistsError); ok {
			cmd.UI.DisplayWarning("User '{{.TargetUserName}}' already has role '{{.RoleType}}' in org '{{.OrgName}}'.", map[string]interface{}{
				"RoleType":       cmd.Args.Role.Role,
				"TargetUserName": cmd.Args.Username,
				"OrgName":        cmd.Args.Organization,
			})
		} else {
			return err
		}
	}

	cmd.UI.DisplayOK()

	return nil
}

func (cmd SetOrgRoleCommand) validateFlags() error {
	if cmd.IsClient && cmd.Origin != "" {
		return translatableerror.ArgumentCombinationError{
			Args: []string{"--client", "--origin"},
		}
	}

	return nil
}

func convertRoleType(givenRole flag.OrgRole) (constant.RoleType, error) {
	switch givenRole.Role {
	case "OrgAuditor":
		return constant.OrgAuditorRole, nil
	case "OrgManager":
		return constant.OrgManagerRole, nil
	case "BillingManager":
		return constant.OrgBillingManagerRole, nil
	default:
		return "", errors.New("Invalid role type.")
	}
}
