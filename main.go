package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {

	fmt.Println("Starting server on port :8080")

	component := hello("world")

	http.Handle("/", templ.Handler(component))

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			fmt.Println("Request received")
			z, err := mem.VirtualMemory()

			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write([]byte(fmt.Sprintf("Total: %vGB, Free: %vGB, Used: %vGB, UsedPercent: %d\n", z.Total>>30, z.Free>>30, z.Used>>30, int(z.UsedPercent))))
		}
	})

	http.ListenAndServe(":8080", nil)
}
