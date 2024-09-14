// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func mainComp() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><link href=\"https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css\" rel=\"stylesheet\"><link href=\"./assets/styles.css\" rel=\"stylesheet\"><title>Hello Template</title><script>\r\n            var memoryData = {};\r\n            var chartInstance = null;\r\n            var data = null;\r\n            var socket;\r\n            var setButton = \"memory\";\r\n            var intervalId;\r\n            var firstHash = false;\r\n            var secondClickOnMemory = false;\r\n            var isInputFocused = false;\r\n\r\n        function connectWebSocket() {\r\n            socket = new WebSocket(\"ws://localhost:8080/ws\");\r\n\r\n            socket.onopen = function(event) {\r\n                setInterval(() => {\r\n                    firstHash = false;\r\n                }, 30000);\r\n\r\n                setIntervalHandler();\r\n            };\r\n\r\n            socket.onmessage = function(event) {\r\n                data = JSON.parse(event.data);\r\n                const type = data?.type;\r\n                console.log(type);\r\n                if (type === \"memory\") {\r\n                    const memoryDiv = document.createElement('div');\r\n                    memoryDiv.classList.add('bg-gray-800', 'text-white', 'p-4', 'mb-4');\r\n                    document.querySelector('.info').innerText = \"\";\r\n                    for (let [key, value] of Object.entries(data)) {\r\n                        if (key === \"type\" || key === \"memory\") {\r\n                            continue;\r\n                        }\r\n                        const p = document.createElement('p');\r\n                        p.innerText = `${key}: ${value}`;\r\n                        if (!firstHash || !secondClickOnMemory) {\r\n                            memoryData[key] = parseFloat(value.split(\" \")[0]);\r\n                        }\r\n                        memoryDiv.appendChild(p);\r\n                    }\r\n                    document.querySelector('.info').appendChild(memoryDiv);\r\n                    if (!firstHash || !secondClickOnMemory) {\r\n                        initializeChart();\r\n                        // Reset after rendering the chart\r\n                    }\r\n                    secondClickOnMemory = false; \r\n                    firstHash = true;\r\n                }\r\n                if (type === \"os\") {\r\n                    document.querySelector('.info').innerText = \"\";\r\n                    const osDiv = document.createElement('div');\r\n                    osDiv.classList.add('bg-gray-800', 'text-white', 'p-4', 'mb-4');\r\n                    document.querySelector('.info').innerText = \"\";\r\n                    for (let [key, value] of Object.entries(data)) {\r\n                        if (key === \"type\") {\r\n                            continue;\r\n                        }\r\n                        const p = document.createElement('p');\r\n                        p.innerText = `${key}: ${value}`;\r\n                        osDiv.appendChild(p);\r\n                    }\r\n                    document.querySelector('.info').appendChild(osDiv);\r\n                    closeChart();\r\n                }\r\n                if (type === \"process\") {\r\n                    // create info for processes\r\n                    document.querySelector('.info').innerText = \"\";\r\n                    const processesInfo = document.createElement('div');\r\n                    processesInfo.classList.add('bg-gray-800', 'text-white', 'p-4', 'mb-4');\r\n                    processesInfo.innerText = data.processes.length + \" processes running\";\r\n\r\n                    const searchInput = document.createElement('input');\r\n                    searchInput.setAttribute('type', 'text');\r\n                    searchInput.setAttribute('placeholder', 'Search process');\r\n                    searchInput.classList.add('bg-gray-700', \"focus:bg-black\", \"focus:outline-none\", 'w-1/2', 'flex', 'flex-col', 'my-4', 'text-white', 'p-2', 'mb-4');\r\n                    searchInput.addEventListener('input', (event) => {\r\n                        const searchValue = event.target.value;\r\n                        const processDivs = document.querySelectorAll('.process-div');\r\n                        \r\n                        processDivs.forEach((div) => {\r\n                            const processName = div.querySelector('.process-name').innerText;\r\n\r\n                            if (processName.toLowerCase().includes(searchValue.toLowerCase())) {\r\n                                div.style.display = 'block';\r\n                            } else {\r\n                                div.style.display = 'none';\r\n                            }\r\n                        });\r\n                    });\r\n\r\n                    searchInput.addEventListener('focus', (event) => {\r\n                        isInputFocused = true;\r\n                        clearInterval(intervalId);\r\n                        intervalId = null;\r\n                    });\r\n\r\n                    searchInput.addEventListener('blur', (event) => {\r\n                        isInputFocused = false;\r\n                        setIntervalHandler();\r\n                    });\r\n\r\n\r\n\r\n                    processesInfo.appendChild(searchInput);\r\n                    document.querySelector('.info').appendChild(processesInfo);\r\n                    \r\n                    // processes mapping\r\n                    const processDiv = document.createElement('div');\r\n                    processDiv.classList.add('bg-gray-800', 'text-white', 'p-4', 'mb-4');\r\n                    for (process of data.processes) {\r\n                        const div = document.createElement('div');\r\n                        div.classList.add('bg-gray-700', 'px-4', 'mb-4', 'break-words', 'process-div')\r\n                        for (let [key, value] of Object.entries(process)) {\r\n                            const p = document.createElement('p');\r\n                            if (key === \"Name\") {\r\n                                p.classList.add('process-name');\r\n                            }\r\n                            p.innerText = `${key} : ${value}`;\r\n                            div.appendChild(p);\r\n                        }\r\n                        processDiv.appendChild(div);\r\n                    }\r\n                    document.querySelector('.info').appendChild(processDiv);\r\n                    closeChart();\r\n                }\r\n                if (type === \"cpu\") {\r\n                    document.querySelector('.info').innerText = \"\";\r\n                    const cpuDiv = document.createElement('div');\r\n                    cpuDiv.classList.add('bg-gray-800', 'text-white', 'p-4', 'mb-4');\r\n                    \r\n                    if (data.cpuInfo) {\r\n                        const cpuInfoDiv = document.createElement('div');\r\n                        data.cpuInfo.forEach((info, index) => {\r\n                            const infoDiv = document.createElement('div');\r\n                            infoDiv.classList.add('bg-gray-700', 'p-4', 'mb-4');\r\n                            for (let [key, value] of Object.entries(info)) {\r\n                                const p = document.createElement('p');\r\n                                p.innerText = `${key}: ${value}`;\r\n                                infoDiv.appendChild(p);\r\n                            }\r\n                            cpuInfoDiv.appendChild(infoDiv);\r\n                        })\r\n                        cpuDiv.appendChild(cpuInfoDiv);\r\n                    }\r\n\r\n                    if (data.cpuTimes) {\r\n                        const cpuTimesDiv = document.createElement('div');\r\n                        cpuTimesDiv.classList.add('bg-gray-700', 'p-4', 'mb-4');\r\n                        data.cpuTimes.forEach((cpu, index) => {\r\n                            const cpuTimeDiv = document.createElement('div');\r\n                            cpuTimeDiv.classList.add('bg-gray-800', 'p-4', 'mb-4');\r\n                            for (let [key, value] of Object.entries(cpu)) {\r\n                                const p = document.createElement('p');\r\n                                p.innerText = `${key}: ${value}`;\r\n                                cpuTimeDiv.appendChild(p);\r\n                            }\r\n                            cpuTimesDiv.appendChild(cpuTimeDiv);\r\n                        })\r\n                        cpuDiv.appendChild(cpuTimesDiv);\r\n                    }\r\n\r\n                    document.querySelector('.info').appendChild(cpuDiv);\r\n                    closeChart();\r\n                }\r\n                if (type === \"network\") {\r\n                    const netDiv = document.createElement('div');\r\n                    netDiv.classList.add('bg-gray-800', 'text-white', 'p-4')\r\n                    document.querySelector('.info').innerText = \"\";\r\n\r\n                    data.net.forEach((network, index) => {\r\n                        const singleNetworkDiv = document.createElement('div');\r\n                        singleNetworkDiv.classList.add('bg-gray-700', 'p-4', 'mb-4')\r\n                        for (let [key, value] of Object.entries(network)) {\r\n                            const p = document.createElement('p');\r\n                            p.innerText = `${key}: ${value}`;\r\n                            singleNetworkDiv.appendChild(p);\r\n                        }\r\n                        netDiv.appendChild(singleNetworkDiv);\r\n                    })\r\n                    closeChart();\r\n                    document.querySelector('.info').appendChild(netDiv);\r\n                }\r\n            };\r\n\r\n            socket.onclose = function(event) {\r\n                console.log(\"WebSocket is closed now. Reconnecting...\");\r\n                setTimeout(connectWebSocket, 1000); // Reconnect after 1 second\r\n            };\r\n\r\n            socket.onerror = function(error) {\r\n                console.log(\"WebSocket error:\", error);\r\n            };\r\n        }\r\n\r\n        function setIntervalHandler() {\r\n            if (intervalId) {\r\n                clearInterval(intervalId);\r\n            }\r\n\r\n            if (!isInputFocused) {\r\n                if (setButton === \"memory\") {\r\n                    intervalId = setInterval(fetchMemoryInfo, 2000);\r\n                } else if (setButton === \"os\") {\r\n                    intervalId = setInterval(fetchOSInfo, 2000);\r\n                } else if (setButton === \"process\") {\r\n                    intervalId = setInterval(fetchProcessInfo, 2000);\r\n                } else if (setButton === \"cpu\") {\r\n                    intervalId = setInterval(fetchCPUInfo, 2000);\r\n                } else if (setButton === \"network\") {\r\n                    intervalId = setInterval(fetchNetworkInfo, 2000);\r\n                }\r\n            }\r\n\r\n        }\r\n\r\n        async function fetchMemoryInfo() {\r\n            try {\r\n                if (setButton === \"memory\") {\r\n                    secondClickOnMemory = true;\r\n                } else {\r\n                    setButton = \"memory\";\r\n                    setIntervalHandler();\r\n                }\r\n                socket.send(\"fetchMemoryInfo\");\r\n            } catch (error) {\r\n                console.error('Error fetching memory info:', error);\r\n            }\r\n        }\r\n\r\n        async function fetchOSInfo() {\r\n            try {\r\n                setButton = \"os\";\r\n                setIntervalHandler();\r\n                socket.send(\"fetchOSInfo\");\r\n            } catch (error) {\r\n                console.error('Error fetching OS info:', error);\r\n            }\r\n        }\r\n\r\n        async function fetchNetworkInfo() {\r\n            try {\r\n                setButton = \"network\";\r\n                setIntervalHandler();\r\n                socket.send(\"fetchNetworkInfo\");\r\n            } catch (error) {\r\n                console.error('Error fetching network info:', error);\r\n            }\r\n        }\r\n\r\n        async function fetchProcessInfo() {\r\n            try {\r\n                setButton = \"process\";\r\n                setIntervalHandler();\r\n                socket.send(\"fetchProcessInfo\");\r\n            } catch (error) {\r\n                console.error('Error fetching process info:', error);\r\n            }\r\n        }\r\n\r\n        async function fetchCPUInfo() {\r\n            try {\r\n                setButton = \"cpu\";\r\n                setIntervalHandler();\r\n                socket.send(\"fetchCPUInfo\");\r\n            } catch (error) {\r\n                console.error('Error fetching CPU info:', error);\r\n            }\r\n        }\r\n\r\n        function initializeChart() {\r\n            const ctx = document.getElementById('myChart').getContext('2d');\r\n            document.getElementById('myChart').style.display = 'block';\r\n\r\n            closeChart();\r\n\r\n            chartInstance = new Chart(ctx, {\r\n                type: 'doughnut',\r\n                data: {\r\n                    labels: ['Total', 'Used', 'Free', 'Available'],\r\n                    datasets: [{\r\n                        label: 'Memory usage',\r\n                        backgroundColor: [\r\n                            'rgba(255, 99, 132, 0.2)', // Red\r\n                            'rgba(54, 162, 235, 0.2)', // Blue\r\n                            'rgba(255, 206, 86, 0.2)', // Yellow\r\n                            'rgba(75, 192, 192, 0.2)'  // Green\r\n                        ],\r\n                        borderColor: [\r\n                            'rgba(255, 99, 132, 1)', // Red\r\n                            'rgba(54, 162, 235, 1)', // Blue\r\n                            'rgba(255, 206, 86, 1)', // Yellow\r\n                            'rgba(75, 192, 192, 1)'  // Green\r\n                        ],\r\n                        data: [memoryData.Total, memoryData.Used, memoryData.Free, memoryData.Available],\r\n                        borderWidth: 2\r\n                    }]\r\n                },\r\n            });\r\n        }\r\n\r\n        function closeChart() {\r\n            if (chartInstance) {\r\n                chartInstance.destroy();\r\n                chartInstance = null;\r\n            }\r\n        }\r\n\r\n        window.onload = () => {\r\n            connectWebSocket();\r\n            document.getElementById('memory-btn').addEventListener('click', fetchMemoryInfo);\r\n            document.getElementById('process-btn').addEventListener('click', fetchProcessInfo);\r\n            document.getElementById('os-btn').addEventListener('click', fetchOSInfo);\r\n            document.getElementById('cpu-btn').addEventListener('click', fetchCPUInfo);\r\n            document.getElementById('network-btn').addEventListener('click', fetchNetworkInfo);\r\n            fetchMemoryInfo();\r\n        };\r\n        </script></head><body class=\"bg-black\"><div class=\"p-4\"><button id=\"memory-btn\" class=\"bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded\">Memory</button> <button id=\"process-btn\" class=\"bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded\">Processes</button> <button id=\"os-btn\" class=\"bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded\">OS</button> <button id=\"cpu-btn\" class=\"bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded\">CPU</button> <button id=\"network-btn\" class=\"bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded\">Network</button></div><div class=\"info p-4\"></div><div class=\"w-1/3 flex content-center mx-auto\"><div class=\"flex w-full justify-center items-center\"><canvas id=\"myChart\"></canvas></div></div><script src=\"https://cdn.jsdelivr.net/npm/chart.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
