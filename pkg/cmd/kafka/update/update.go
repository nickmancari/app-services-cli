package update

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type Options struct {
	name  string
	id    string
	owner string

	outputFormat string

	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

func NewUpdateCommand(f *factory.Factory) *cobra.Command {
	opts := Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		localizer:  f.Localizer,
		Logger:     f.Logger,
		Connection: f.Connection,
	}

	cmd := cobra.Command{
		Use:   "update",
		Short: "Update a Kafka instance",
		Long:  "Update a Kafka instance",
		Args:  cobra.RangeArgs(0, 1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return cmdutil.FilterValidKafkas(f, toComplete)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if len(args) > 0 {
				opts.name = args[0]
			}

			if opts.name != "" && opts.id != "" {
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			var kafkaConfig *config.KafkaConfig
			if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
				return opts.localizer.MustLocalizeError("kafka.common.error.noKafkaSelected")
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return run(&opts)
		},
	}

	// TODO: Add dynamic flag completions
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.update.flag.id"))
	cmd.Flags().StringVar(&opts.owner, "owner", "", opts.localizer.MustLocalize("kafka.update.flag.owner"))

	_ = cmd.MarkFlagRequired("owner")

	return &cmd
}

func run(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	var kafkaInstance *kafkamgmtclient.KafkaRequest
	ctx := context.Background()
	if opts.name != "" {
		kafkaInstance, _, err = kafka.GetKafkaByName(ctx, api.Kafka(), opts.name)
		if err != nil {
			return err
		}
	} else {
		kafkaInstance, _, err = kafka.GetKafkaByID(ctx, api.Kafka(), opts.id)
		if err != nil {
			return err
		}
	}

	updateObj := kafkamgmtclient.NewKafkaUpdateRequest(opts.owner)

	updatedInstance, _, err := api.Kafka().UpdateKafkaById(context.Background(), kafkaInstance.GetId()).KafkaUpdateRequest(*updateObj).Execute()
	if err != nil {
		return err
	}

	accounts, _, err := api.AccountMgmt().ApiAccountsMgmtV1AccountsGet(context.Background()).Execute()
	
	for _, a := range accounts.Items {
		fmt.Println(a.Id)
	}
	return nil

	data, _ := json.Marshal(updatedInstance)
	dump.JSON(opts.IO.Out, data)

	return nil
}
