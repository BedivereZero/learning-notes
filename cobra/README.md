# Cobra

## 创建 root 命令

main.go:

```go
package main

import (
	"github.com/BedivereZero/learning-notes/cobra/cmd"
)

func main() {
	cmd.Execute()
}
```

cmd/root.go:

```go
package

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "This is a cobra example",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

添加子命令

cmd/version.go

```go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sample v0.1.0")
	},
}
```

添加 flag

cmd/version.go

```go
var (
	format string
)

func init() {
    	versionCmd.Flags().StringVarP(&format, "format", "f", "", "format of output")
}
```

添加在子命令中可用的 flag

cmd/root.go

```go
var (
    verbose string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
```

添加必要的 flag

cmd/root.go

```go
var (
    region string
)

func init() {
    rootCmd.Flags().StringVarP(&region, "region", "r", "", "AWS region (required)")
    rootCmd.MarkFlagRequired("region")
}
```

定义命令的参数

cmd/version.go:

```go
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sample v0.1.0")
	},
	Args: cobra.NoArgs,
}
```

- `NoArgs`: 如果有任何参数，命令将报错
- `ArbitraryArgs`: 命令将接受任何参数
- `OnlyValidArgs`: 命令仅接受 `ValidArgs` 属性中定义的参数
- `MinimumNArgs(int)`: 如果参数数量少于指定数量，命令将报错
- `MaximumNArgs(int)`: 如果参数数量多于指定数量，命令将报错
- `ExactArgs(int)`: 如果命令参数数量不是指定数量，命令将报错
- `ExactValidArgs(int)`: 如果命令参数数量部署指定数量，或不是 `ValidArgs` 属性中定义的参数，命令将报错
- `RangeArgs(min, max)`: 如果命令参数数量不在指定范围内，命令将报错
