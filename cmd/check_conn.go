package cmd

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
	"github.com/spf13/cobra"
)

var checkConnCmd = &cobra.Command{
	Use:   "check",
	Short: "Runs a single connection check",
	Long:  "Checks the connection status by starting a set of requests using Speedtest by Ookla, then reports the results back to the user.",
	RunE:  checkConn,
}

func checkConn(cmd *cobra.Command, args []string) error {
	fmt.Println("Starting connection test...")
	var speedtestClient = speedtest.New()

	serverList, err := speedtestClient.FetchServers()
	if err != nil {
		return err
	}

	targets, err := serverList.FindServer([]int{})
	if err != nil {
		return err
	}

	for _, s := range targets {
		err = s.PingTest(nil)
		if err != nil {
			return err
		}
		s.DownloadTest()
		s.UploadTest()
		fmt.Printf("Latency: %s, Download %.2f Mbps, Upload: %.2f Mbps\n", s.Latency, s.DLSpeed, s.ULSpeed)
		s.Context.Reset()
	}

	return nil
}

func init() {
	rootCmd.AddCommand(checkConnCmd)
}
