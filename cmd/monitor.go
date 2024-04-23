package cmd

import (
	"fmt"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
	"github.com/spf13/cobra"
)

var monitorConnCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitors the connection status",
	Long:  "Monitors the connection status by running in a loop and periodically checking the connection.",
	RunE:  monitorConn,
}

func monitorConn(cmd *cobra.Command, args []string) error {
	var speedtestClient = speedtest.New()

	fmt.Printf("Starting up connection monitoring at %s...\n", time.Now().Format("2006-01-02 15:04:05"))

	for {
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
			connStatus := "UP"
			currTime := time.Now().Format("2006-01-02 15:04:05")
			if err != nil {
				fmt.Printf("Connection is DOWN at %s\n", currTime)
				connStatus = "DOWN"
			}
			fmt.Printf("Connection is UP with %s of latency (%s jitter) at %s\n", s.Latency, s.Jitter, currTime)
			cfg.Logger.Info("Connection status", "time", currTime, "latency", s.Latency.Milliseconds(), "jitter", s.Jitter.Milliseconds(), "status", connStatus)
			s.Context.Reset()
		}

		time.Sleep(time.Duration(cfg.CheckInterval) * time.Second) // Adjust the interval as needed
	}
}

func init() {
	rootCmd.AddCommand(monitorConnCmd)
}
