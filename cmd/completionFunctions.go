package cmd

import (
	"fmt"
	"time"

	"github.com/buonotti/apisense/api/db"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/conversion"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/definitions"
	"github.com/buonotti/apisense/validation/pipeline"
)

func validReportsFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		reports, _ := pipeline.Reports()
		return util.Map(reports, func(r pipeline.Report) string {
			return fmt.Sprintf("%s\t%s with %d result(s)", r.Id, time.Time(r.Time).Format("2006-01-02 at 15-04-05.000Z"), len(r.Endpoints))
		}), cobra.ShellCompDirectiveNoFileComp
	}
}

func validConfigKeysFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		keys := viper.AllKeys()
		return keys, cobra.ShellCompDirectiveNoFileComp
	}
}

func validFormatsFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return conversion.Converters(), cobra.ShellCompDirectiveNoFileComp
	}
}

func validDefinitionsFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		endpointDefinitions, _ := definitions.Endpoints()
		return util.Map(endpointDefinitions, func(d definitions.Endpoint) string {
			return fmt.Sprintf("%s\t%s", d.FileName, d.Name)
		}), cobra.ShellCompDirectiveNoFileComp
	}
}

func validEnabledDefinitionFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		endpointDefinitions, _ := definitions.Endpoints()
		mappedDefinitions := util.Map(endpointDefinitions, func(d definitions.Endpoint) string {
			if d.IsEnabled {
				return fmt.Sprintf("%s\t%s", d.FileName, d.Name)
			}
			return ""
		})
		return util.Where(mappedDefinitions, func(s string) bool {
			return s != ""
		}), cobra.ShellCompDirectiveNoFileComp
	}
}

func validDisabledDefinitionFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		endpointDefinitions, _ := definitions.Endpoints()
		mappedDefinitions := util.Map(endpointDefinitions, func(d definitions.Endpoint) string {
			if !d.IsEnabled {
				return fmt.Sprintf("%s%s\t%s", viper.GetString("daemon.ignore-prefix"), d.FileName, d.Name)
			}
			return ""
		})
		return util.Where(mappedDefinitions, func(s string) bool {
			return s != ""
		}), cobra.ShellCompDirectiveNoFileComp
	}
}

func validUsersFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		users, _ := db.ListUsers()
		return util.Map(users, func(u db.User) string {
			return fmt.Sprintf("%s\t%v", u.Username, u.Enabled)
		}), cobra.ShellCompDirectiveNoFileComp
	}
}

func validEnabledUserFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		users, _ := db.ListUsers()
		mappedUsers := util.Map(users, func(u db.User) string {
			if u.Enabled {
				return fmt.Sprintf("%s", u.Username)
			}
			return ""
		})
		return util.Where(mappedUsers, func(s string) bool {
			return s != ""
		}), cobra.ShellCompDirectiveNoFileComp
	}
}

func validDisabledUserFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		users, _ := db.ListUsers()
		mappedUsers := util.Map(users, func(u db.User) string {
			if !u.Enabled {
				return fmt.Sprintf("%s", u.Username)
			}
			return ""
		})
		return util.Where(mappedUsers, func(s string) bool {
			return s != ""
		}), cobra.ShellCompDirectiveNoFileComp
	}
}

func validLogLevelsFunc() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"debug", "info", "warning", "error", "fatal"}, cobra.ShellCompDirectiveNoFileComp
	}
}
