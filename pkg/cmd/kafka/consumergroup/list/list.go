package list

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/consumergroup"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams

	output  string
	kafkaID string
	limit   int32
	topic   string
}

type consumerGroupRow struct {
	ConsumerGroupID   string `json:"groupId,omitempty" header:"Consumer group ID"`
	ActiveMembers     int    `json:"active_members,omitempty" header:"Active members"`
	PartitionsWithLag int    `json:"lag,omitempty" header:"Partitions with lag"`
}

// NewListConsumerGroupCommand creates a new command to list consumer groups
func NewListConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.consumerGroup.list.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.consumerGroup.list.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.consumerGroup.list.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.consumerGroup.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" && !flagutil.IsValidInput(opts.output, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.output, flagutil.ValidOutputFormats...)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return fmt.Errorf(localizer.MustLocalizeFromID("kafka.consumerGroup.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runList(opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.limit, "limit", "", 1000, localizer.MustLocalizeFromID("kafka.consumerGroup.list.flag.limit"))
	cmd.Flags().StringVarP(&opts.output, "output", "o", "", localizer.MustLocalize(&localizer.Config{
		MessageID:   "kafka.consumerGroup.common.flag.output.description",
		PluralCount: 2,
	}))
	cmd.Flags().StringVar(&opts.topic, "topic", "", localizer.MustLocalizeFromID("kafka.consumerGroup.list.flag.topic.description"))

	return cmd
}

func runList(opts *Options) (err error) {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()

	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	req := api.GetConsumerGroupList(ctx)
	req = req.Limit(opts.limit)
	if opts.topic != "" {
		req = req.Topic(opts.topic)
	}
	consumerGroupData, httpRes, err := req.Execute()
	if err != nil {
		if httpRes == nil {
			return err
		}

		switch httpRes.StatusCode {
		case 401:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID:   "kafka.consumerGroup.common.error.unauthorized",
				PluralCount: 2,
				TemplateData: map[string]interface{}{
					"Operation": "list",
				},
			}))
		case 403:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID:   "kafka.consumerGroup.common.error.forbidden",
				PluralCount: 2,
				TemplateData: map[string]interface{}{
					"Operation": "list",
				},
			}))
		case 500:
			return errors.New(localizer.MustLocalizeFromID("kafka.consumerGroup.common.error.internalServerError"))
		case 503:
			return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.common.error.unableToConnectToKafka",
				TemplateData: map[string]interface{}{
					"Name": kafkaInstance.GetName(),
				},
			}), err)
		default:
			return err
		}
	}

	ok, err := checkForConsumerGroups(int(consumerGroupData.GetCount()), opts, kafkaInstance.GetName())
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	switch opts.output {
	case "json":
		data, _ := json.Marshal(consumerGroupData)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(consumerGroupData)
		_ = dump.YAML(opts.IO.Out, data)
	default:
		logger.Info("")
		topics := consumerGroupData.GetItems()
		rows := mapConsumerGroupResultsToTableFormat(topics)
		dump.Table(opts.IO.Out, rows)

		return nil
	}

	return nil

}

func mapConsumerGroupResultsToTableFormat(consumerGroups []strimziadminclient.ConsumerGroup) []consumerGroupRow {
	var rows []consumerGroupRow = []consumerGroupRow{}

	for _, t := range consumerGroups {
		consumers := t.GetConsumers()
		row := consumerGroupRow{
			ConsumerGroupID:   t.GetGroupId(),
			ActiveMembers:     consumergroup.GetActiveConsumersCount(consumers),
			PartitionsWithLag: consumergroup.GetPartitionsWithLag(consumers),
		}
		rows = append(rows, row)
	}

	return rows
}

// checks if there are any consumer groups available
// prints to stderr if not
func checkForConsumerGroups(count int, opts *Options, kafkaName string) (hasCount bool, err error) {
	logger, err := opts.Logger()
	if err != nil {
		return false, err
	}

	if count == 0 && opts.output == "" {
		if opts.topic == "" {
			logger.Info(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.list.log.info.noConsumerGroups",
				TemplateData: map[string]interface{}{
					"InstanceName": kafkaName,
				},
			}))
		} else {
			logger.Info(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.list.log.info.noConsumerGroupsForTopic",
				TemplateData: map[string]interface{}{
					"InstanceName": kafkaName,
					"TopicName":    opts.topic,
				},
			}))
		}

		return false, nil
	}

	return true, nil
}