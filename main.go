package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	fmt.Println("Starting server on port :8080")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		for {
			_, message, err := conn.ReadMessage()
			fmt.Println(string(message))
			if err != nil {
				fmt.Println(err)
				break
			}

			if string(message) == "fetchMemoryInfo" {
				memInfo, err := mem.VirtualMemory()
				if err != nil {
					fmt.Println(err)
					return
				}

				memoryJSON := map[string]string{
					"Total":       fmt.Sprintf("%.3f GB", float64(memInfo.Total)/float64(1<<30)),
					"Available":   fmt.Sprintf("%.3f GB", float64(memInfo.Available)/float64(1<<30)),
					"Used":        fmt.Sprintf("%.3f GB", float64(memInfo.Used)/float64(1<<30)),
					"Free":        fmt.Sprintf("%.3f GB", float64(memInfo.Free)/float64(1<<30)),
					"UsedPercent": fmt.Sprintf("%d%%", int(memInfo.UsedPercent)),
				}

				infoJSON := map[string]interface{}{
					"memory": memoryJSON,
				}

				err = conn.WriteJSON(infoJSON)
				if err != nil {
					fmt.Println(err)
					break
				}
			} else if string(message) == "fetchOSInfo" {

				hostInfo, err := host.Info()

				if err != nil {
					fmt.Println(err)
					return
				}

				bootTime := int64(hostInfo.BootTime)
				uptime := int64(hostInfo.Uptime)

				days := uptime / (24 * 3600)
				hours := (uptime % (24 * 3600)) / 3600
				minutes := (uptime % 3600) / 60
				seconds := uptime % 60

				uptimeString := fmt.Sprintf("%d days, %d hours, %d minutes %d seconds", days, hours, minutes, seconds)
				formattedBootTime := time.Unix(bootTime, 0).Format("2006-01-02 15:04:05")

				hostJSON := map[string]string{
					"Host":            hostInfo.Hostname,
					"OS":              strings.ToUpper(hostInfo.OS[:1]) + hostInfo.OS[1:],
					"Platform":        hostInfo.Platform,
					"PlatformFamily":  hostInfo.PlatformFamily,
					"PlatformVersion": hostInfo.PlatformVersion,
					"KernelVersion":   hostInfo.KernelVersion,
					"KernelArch":      hostInfo.KernelArch,
					"BootTime":        formattedBootTime,
					"Procs":           fmt.Sprintf("%d", hostInfo.Procs),
					"Uptime":          uptimeString,
				}

				err = conn.WriteJSON(hostJSON)

				if err != nil {
					fmt.Println(err)
					break
				}
			}
		}
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			fmt.Println("Request received")
			memInfo, err := mem.VirtualMemory()

			if err != nil {
				fmt.Println(err)
				return
			}

			hostInfo, err := host.Info()

			if err != nil {
				fmt.Println(err)
				return
			}

			memoryJSON := map[string]string{
				"Total":       fmt.Sprintf("%.3f GB", float64(memInfo.Total)/float64(1<<30)),
				"Available":   fmt.Sprintf("%.3f GB", float64(memInfo.Available)/float64(1<<30)),
				"Used":        fmt.Sprintf("%.3f GB", float64(memInfo.Used)/float64(1<<30)),
				"Free":        fmt.Sprintf("%.3f GB", float64(memInfo.Free)/float64(1<<30)),
				"UsedPercent": fmt.Sprintf("%d%%", int(memInfo.UsedPercent)),
			}

			bootTime := int64(hostInfo.BootTime)
			uptime := int64(hostInfo.Uptime)

			days := uptime / (24 * 3600)
			hours := (uptime % (24 * 3600)) / 3600
			minutes := (uptime % 3600) / 60
			seconds := uptime % 60

			uptimeString := fmt.Sprintf("%d days, %d hours, %d minutes %d seconds", days, hours, minutes, seconds)
			formattedBootTime := time.Unix(bootTime, 0).Format("2006-01-02 15:04:05")

			hostJSON := map[string]string{
				"Host":            hostInfo.Hostname,
				"OS":              strings.ToUpper(hostInfo.OS[:1]) + hostInfo.OS[1:],
				"Platform":        hostInfo.Platform,
				"PlatformFamily":  hostInfo.PlatformFamily,
				"PlatformVersion": hostInfo.PlatformVersion,
				"KernelVersion":   hostInfo.KernelVersion,
				"KernelArch":      hostInfo.KernelArch,
				"BootTime":        formattedBootTime,
				"Procs":           fmt.Sprintf("%d", hostInfo.Procs),
				"Uptime":          uptimeString,
			}

			processes, err := process.Processes()

			if err != nil {
				fmt.Println(err)
				return
			}

			processJSON := make([]map[string]string, 0)

			for _, p := range processes {
				name, _ := p.Name()
				pid := p.Pid
				ppid, _ := p.Ppid()
				username, _ := p.Username()
				cmdline, _ := p.Cmdline()

				// if len(cmdline) > 100 {
				// 	cmdline = cmdline[:100]
				// }

				processJSON = append(processJSON, map[string]string{
					"Name":     name,
					"PID":      fmt.Sprintf("%d", pid),
					"PPID":     fmt.Sprintf("%d", ppid),
					"Username": username,
					"Cmdline":  cmdline,
				})
			}

			infoJSON := map[string]interface{}{
				"memory":    memoryJSON,
				"host":      hostJSON,
				"processes": processJSON,
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(infoJSON)
		}
	})

	component := hello("world")
	http.Handle("/", templ.Handler(component))

	http.ListenAndServe(":8080", nil)
}
