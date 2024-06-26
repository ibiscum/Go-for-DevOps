// Package handlers provides an Ops type that has methods that implement bot.HandleFunc for various commands that could be sent to a bot.
package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ibiscum/Go-for-DevOps/chapter11/chatbot/bot"
	"github.com/ibiscum/Go-for-DevOps/chapter11/ops/client"

	"github.com/olekukonko/tablewriter"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"

	pb "github.com/ibiscum/Go-for-DevOps/chapter11/ops/proto"
)

// Ops provides bot.HandleFunc methods that can reuse the connections to the Ops service.
type Ops struct {
	OpsClient *client.Ops
	API       *slack.Client
	SMClient  *socketmode.Client
}

// write writes a formatted string to the event output in the bot.Message.
func (o Ops) write(m bot.Message, s string, i ...interface{}) error {
	_, _, err := o.API.PostMessage(
		m.AppMention.Channel,
		slack.MsgOptionText(fmt.Sprintf(s, i...), false),
	)
	return err
}

// Register registers all the commands held in Ops with the bot.
func (o Ops) Register(b *bot.Bot) {
	b.Register(regexp.MustCompile(`^\s*help`), o.Help)
	b.Register(regexp.MustCompile(`^\s*list traces`), o.ListTraces)
	b.Register(regexp.MustCompile(`^\s*show trace`), o.ShowTrace)
	b.Register(regexp.MustCompile(`^\s*change sampling`), o.ChangeSampling)
	b.Register(regexp.MustCompile(`^\s*show logs`), o.ShowLogs)
	b.Register(nil, o.lastResort)
}

// opt stores the key/value pair for an option to a command.
type opt struct {
	key string
	val string
}

// listTracesRE teases the options from a `list traces` command.
var listTracesRE = regexp.MustCompile(`(\S+)=(?:(\S+))`)

// ListTraces lists all the traces requested in a table that is output to the user.
func (o Ops) ListTraces(ctx context.Context, m bot.Message) {
	sp := strings.Split(m.Text, "list traces")
	if len(sp) != 2 {
		err := o.write(m, "The 'list traces' command is malformed")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	t := strings.TrimSpace(sp[1])

	kvOpts := []opt{}
	for _, match := range listTracesRE.FindAllStringSubmatch(t, -1) {
		kvOpts = append(kvOpts, opt{strings.TrimSpace(match[1]), strings.TrimSpace(match[2])})
	}
	options := []client.CallOption{}

	for _, opt := range kvOpts {
		switch opt.key {
		case "operation":
			options = append(options, client.WithOperation(opt.val))
		case "start":
			t, err := time.Parse(`01/02/2006-15:04:05`, opt.val)
			if err != nil {
				err := o.write(m, "The start option must be in the form `01/02/2006-15:04:05` for UTC")
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			options = append(options, client.WithStart(t))
		case "end":
			if opt.val == "now" {
				continue
			}
			t, err := time.Parse(`01/02/2006-15:04:05`, opt.val)
			if err != nil {
				err := o.write(m, "The end option must be in the form `01/02/2006-15:04:05` for UTC")
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			options = append(options, client.WithEnd(t))
		case "limit":
			i, err := strconv.Atoi(opt.val)
			if err != nil {
				err = o.write(m, "The limit option must be an integer")
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			if i > 100 {
				err = o.write(m, "Cannot request more than 100 traces")
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			options = append(options, client.WithLimit(int32(i)))
		case "tags":
			tags, err := convertList(opt.val)
			if err != nil {
				err := o.write(m, "tags: must enclosed in [], like tags=[tag,tag2]")
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			options = append(options, client.WithLabels(tags))
		default:
			err := o.write(m, "don't understand an option type(%s)", opt.key)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	traces, err := o.OpsClient.ListTraces(ctx, options...)
	if err != nil {
		err = o.write(m, "Ops server had an error: %s", err)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	b := strings.Builder{}
	b.WriteString("Here are the traces you requested:\n")

	table := tablewriter.NewWriter(&b)
	table.SetHeader([]string{"Start Time(UTC)", "Trace ID"})
	for _, item := range traces {
		table.Append(
			[]string{
				item.Start.Format("01/02/2006 04:05"),
				"http://127.0.0.1:16686/trace/" + item.ID,
			},
		)
	}
	table.Render()
	err = o.write(m, b.String())
	if err != nil {
		log.Fatal(err)
	}
}

// ShowTrace gives the URL to a trace ID.
func (o Ops) ShowTrace(ctx context.Context, m bot.Message) {
	sp := strings.Split(m.Text, "show trace")
	if len(sp) != 2 {
		err := o.write(m, `show trace command should be in form: show trace <id>`)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	id := strings.TrimSpace(sp[1])

	trace, err := o.OpsClient.ShowTrace(ctx, id)
	if err != nil {
		err := o.write(m, "Ops server had an error: %s", err)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	b := strings.Builder{}
	table := tablewriter.NewWriter(&b)
	b.WriteString("Here is some basic trace data:\n")
	table.Append([]string{"ID", trace.Id})
	table.Append([]string{"Duration", trace.Duration.AsDuration().String()})
	table.Append([]string{"Jaeger URL", "http://127.0.0.1:16686/trace/" + trace.Id})
	if len(trace.Errors) > 0 {
		table.Append([]string{"Had Errors", "true"})
	} else {
		table.Append([]string{"Had Errors", "false"})
	}
	table.Render()
	b.WriteString("\n")

	if len(trace.Errors) > 0 {
		table = tablewriter.NewWriter(&b)
		b.WriteString("Here are the errors from the trace:\n")
		for _, err := range trace.Errors {
			table.Append([]string{err})
		}
		table.Render()
		b.WriteString("\n")
	}

	b.WriteString("Here are the operations in the trace:\n")
	table = tablewriter.NewWriter(&b)
	for _, op := range trace.Operations {
		table.Append([]string{op})
	}
	table.Render()
	b.WriteString("\n")

	err = o.write(m, "%s,\nHere is the trace info you requested:\n\n%s", m.User.Name, b.String())
	if err != nil {
		log.Fatal(err)
	}
}

// ShowLogs outputs the logs given a trace ID.
func (o Ops) ShowLogs(ctx context.Context, m bot.Message) {
	sp := strings.Split(m.Text, "show logs")
	if len(sp) != 2 {
		err := o.write(m, `show logs command should be in form: show logs <id>`)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	id := strings.TrimSpace(sp[1])
	log.Println("show logs id==", id)
	logs, err := o.OpsClient.ShowLogs(ctx, id)
	if err != nil {
		err := o.write(m, "Ops server had an error: %s", err)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	b := strings.Builder{}
	n := time.Now().UTC()
	for _, l := range logs {
		var t string
		if l.Time.Year() == n.Year() && l.Time.Month() == n.Month() && l.Time.Day() == n.Day() {
			t = l.Time.Format(`15:04:05`)

		} else {
			t = l.Time.Format(`a01/02/2006 15:04:05`)
		}
		b.WriteString(fmt.Sprintf("%s: %s: %s\n", t, l.Key, l.Value))
	}

	err = o.write(m, "%s,\nHere are the logs you requested for trace %s:\n\n%s", m.User.Name, id, b.String())
	if err != nil {
		log.Fatal(err)
	}
}

var sampleTypeRE = regexp.MustCompile(`^\s*(never|always|float)`)

// ChangeSampling changes the sampling type/rate on the server.
func (o Ops) ChangeSampling(ctx context.Context, m bot.Message) {
	sp := strings.Split(m.Text, "change sampling")
	if len(sp) != 2 {
		err := o.write(m, `change sampling command should be in form: change sampling <type> <options>`)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	t := strings.TrimSpace((sp[1]))

	sub := sampleTypeRE.FindStringSubmatch(t)
	if len(sub) == 0 {
		err := o.write(m, `I don't have support for the samplling type you requested, sorry...`)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	req := &pb.ChangeSamplingReq{}

	switch sub[1] {
	case "never":
		req.Type = pb.SamplerType_STNever
	case "always":
		req.Type = pb.SamplerType_STAlways
	case "float":
		req.Type = pb.SamplerType_STFloat

		sp := strings.Split(t, "float")
		if len(sp) != 2 {
			err := o.write(m, `'change sampling float' must be followed by a float that is > 0 and <= 1`)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		f, err := strconv.ParseFloat(strings.TrimSpace(sp[1]), 64)
		if err != nil {
			err := o.write(m, `'change sampling float' had an invalid float option: %q`, strings.TrimSpace(sp[1]))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		if f <= 0 || f > 1 {
			err := o.write(m, `'change sampling float' must be followed by a float that is > 0 and <= 1`)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		req.FloatValue = f
	default:
		err := o.write(m, `sorry, I hit a bug, I kinda understand %q, so you need to talk to my creator`, m.Text)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err := o.OpsClient.ChangeSampling(ctx, req)
	if err != nil {
		err := o.write(m, "Ops server gave an error on changing the sampling: %s", err)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
}

var cmdList string

func init() {
	cmds := []string{}
	for k := range help {
		cmds = append(cmds, k)
	}
	sort.Strings(cmds)

	b := strings.Builder{}
	for _, cmd := range cmds {
		b.WriteString(cmd + "\n")
	}
	b.WriteString("You can get more help by saying `help <cmd>` with a command from above.\n")
	cmdList = b.String()
}

// Help returns help about various commands.
func (o Ops) Help(ctx context.Context, m bot.Message) {
	sp := strings.Split(m.Text, "help")
	if len(sp) < 2 {
		err := o.write(m, "%s,\nYou have to give me a command you want help with", m.User.Name)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	cmd := strings.TrimSpace(strings.Join(sp[1:], ""))
	if cmd == "" {
		err := o.write(m, "Here are all the commands that I can help you with:\n%s", cmdList)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if v, ok := help[cmd]; ok {
		err := o.write(m, "I can help you waith that:\n%s", v)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err := o.write(m, "%s,\nI don't know what %q is to give you help", m.User.Name, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func (o Ops) lastResort(ctx context.Context, m bot.Message) {
	err := o.write(m, "%s,\nI don't have anything that handles what you sent", m.User.Name)
	if err != nil {
		log.Fatal(err)
	}
}

func convertList(s string) ([]string, error) {
	if string(s[0]) != `[` || string(s[len(s)-1]) != `]` {
		return nil, errors.New("must enclosed in [], like [tag,tag2] comma deliminated with no spaces")
	}

	s = strings.TrimPrefix(s, `[`)
	s = strings.TrimSuffix(s, `]`)
	sp := strings.Split(s, ",")
	tags := []string{}
	for _, t := range sp {
		tags = append(tags, strings.TrimSpace(t))
	}
	return tags, nil
}
