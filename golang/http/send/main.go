package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	NewRootCmd(os.Stdout).Execute()
}

func NewRootCmd(w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use: "send",
	}
	cmd.AddCommand(NewSendCommand(w))
	return cmd
}

func NewSendCommand(w io.Writer) *cobra.Command {
	o := NewDefaultSendOptions()

	cmd := &cobra.Command{
		Use:  "send",
		Long: "send http request",
		RunE: func(cmd *cobra.Command, args []string) error {
			ch := NewMaxTimeChan(o.Times)

			var g sync.WaitGroup
			g.Add(o.Threads)

			// 进度条
			bar := progressbar.Default(int64(o.Times))

			// 测试过程中的错误
			var errs []error

			// 开始时间
			s := time.Now()
			for t := 0; t < o.Threads; t++ {
				go func(t int) {
					defer g.Done()
					for i := range ch {
						client := &http.Client{}
						if err := send(client, o); err != nil {
							errs = append(errs, fmt.Errorf("%3d %3d send http request fail: %w", t, i, err))
						}
						bar.Add(1)
					}
				}(t)
			}
			g.Wait()
			bar.Exit()

			// 打印测试过程中发生的错误
			for _, err := range errs {
				fmt.Fprintln(cmd.OutOrStdout(), err)
			}

			// 耗时
			d := time.Since(s)
			fmt.Fprintln(cmd.OutOrStdout(), "duration", d)
			// qps
			qps := float64(o.Times) * float64(time.Second) / float64(d)
			fmt.Fprintf(cmd.OutOrStdout(), "qps %.3f\n", qps)

			return nil
		},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   false,
			DisableNoDescFlag:   false,
			DisableDescriptions: false,
			HiddenDefaultCmd:    false,
		},
		SilenceUsage: true,
	}
	cmd.SetOut(w)

	o.AddFlag(cmd.Flags())

	return cmd
}

type SendOptions struct {
	Method string
	URL    string
	Body   FileValue

	BearerToken string

	Threads int
	Times   int
}

func NewDefaultSendOptions() *SendOptions {
	return &SendOptions{
		Method:  http.MethodGet,
		URL:     "http://localhost:8080",
		Threads: 1,
		Times:   10,
	}
}

type FileValue struct {
	Name  string
	Bytes []byte
}

// Set implements pflag.Value.
func (f *FileValue) Set(s string) (err error) {
	f.Name = s

	if f.Name == "" {
		return nil
	}

	f.Bytes, err = os.ReadFile(f.Name)
	return err
}

// String implements pflag.Value.
func (f *FileValue) String() string {
	return f.Name
}

// Type implements pflag.Value.
func (*FileValue) Type() string {
	return "file"
}

var _ pflag.Value = &FileValue{}

func (o *SendOptions) AddFlag(fs *pflag.FlagSet) {
	fs.SortFlags = false

	httpFlagSet := pflag.NewFlagSet("http", pflag.ContinueOnError)
	httpFlagSet.SortFlags = false
	httpFlagSet.StringVar(&o.Method, "method", o.Method, "http method")
	httpFlagSet.StringVar(&o.URL, "url", o.URL, "http url")
	httpFlagSet.Var(&o.Body, "body", "load http body from this file, if specified")
	httpFlagSet.StringVar(&o.BearerToken, "bearer-token", o.BearerToken, "bearer authentication token")
	fs.AddFlagSet(httpFlagSet)

	fs.IntVar(&o.Times, "times", o.Times, "number of times to send http request")
	fs.IntVar(&o.Threads, "threads", o.Threads, "threads number")
}

func NewMaxTimeChan(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func send(client *http.Client, o *SendOptions) error {
	req, err := http.NewRequest(o.Method, o.URL, bytes.NewReader(o.Body.Bytes))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if o.BearerToken != "" {
		resp.Header.Add("Authorization", "Bearer "+o.BearerToken)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			body = fmt.Appendf(body, "read http response body fail: %v", err)
		}
		if len(body) == 0 {
			return fmt.Errorf("%s", resp.Status)
		}
		return fmt.Errorf("%d, body: %s", resp.StatusCode, body)
	}

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		return err
	}

	return nil
}
