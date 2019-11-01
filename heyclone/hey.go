package main

import (
	"flag"
	"fmt"
	"net/http"
)

type BoomContent struct {
	Host           string
	Request_Number int
}

func (b *BoomContent) ArgsParse() {
	//解析命令行参数，并输出帮助文件
	request_number := flag.Int("n", 0, "Number of requests to run. Default is 200.")
	request_concurrency := flag.Int("c", 50,
		"Number of requests to run concurrently. Total number of requests cannot be smaller than the concurrency level. Default is 50.")

	request_rate := flag.Int("-q", 0, "Rate limit, in queries per second (QPS). Default is no rate limit.")
	request_duration := flag.Int("-z", 10,
		`Duration of application to send requests. 
	When duration is reached,application stops and exits. If duration is specified, n is ignored.Examples: -z 10s -z 3m.`)
	output_type := flag.String("o", "csv", `Output type. If none provided, a summary is printed.
      "csv" is the only supported alternative. Dumps the response
      metrics in comma-separated values format.
		`)
	request_method := flag.String("m", "GET", `HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.`)
	custom_http_header := flag.String("H", "", `Custom HTTP header. You can specify as many as needed by repeating the flag.
      For example, -H "Accept: text/html" -H "Content-Type: application/xml" .`)
	request_time_out := flag.Int("-t", 10, `Timeout for each request in seconds. Default is 20, use 0 for infinite.`)
	request_accept_header := flag.String("A", "", `HTTP Accept header.`)
	request_http_body := flag.String("d", "", `HTTP request body`)
	request_body_from_file := flag.String("D", "", ` HTTP request body from file. For example, /home/user/file.txt or ./file.txt.`)
	request_content_type := flag.String("T", "", `Content-type, defaults to "text/html".`)
	request_authentication := flag.String("a", "", `Basic authentication, username:password.`)
	request_proxy_addr := flag.String("x", "", `HTTP Proxy address as host:port.`)
	request_http_version := flag.Bool("h2", false, `Enable HTTP/2.`)
	request_http_host := flag.String("host", "", `HTTP Host header.`)
	request_disable_comp := flag.Bool("disable-compression", false, "Disable compression.")
	request_keepalive_set := flag.Bool("disable-keepalive", false, `Disable keep-alive, prevents re-use of TCP
                        connections between different HTTP requests.`)
	request_rewrite_set := flag.Bool("disable-redirects", false, `Disable following of HTTP redirects`)
	request_cpu_set := flag.Int("cpus", 1, `Number of used cpu cores.
                        (default for current machine is 8 cores)`)
	flag.Parse()
	fmt.Println(*request_number, *request_concurrency, *request_rate, *request_duration,
		*output_type, *request_method, *custom_http_header, *request_time_out, *request_accept_header, *request_http_body, *request_body_from_file,
		*request_content_type, *request_authentication, *request_proxy_addr, *request_http_version, *request_http_host, *request_disable_comp, *request_keepalive_set,
		*request_rewrite_set, *request_cpu_set)
	b.Host = *request_http_host
	b.Request_Number = *request_number
}

func (b *BoomContent) GenerateRequest(stats chan string) {
	resp, err := http.Get(b.Host)
	if err != nil {
		fmt.Println("连接错误，程序退出!错误原因为:", nil)
		return
	}
	if resp.Status != "200" {
		fmt.Println("返回的状态码非200，程序退出!")
		return
	}
	stats <- resp.Status
}

func (b *BoomContent) MainController(stats chan string) {
	for i := 0; i < b.Request_Number; i++ {
		go b.GenerateRequest(stats)
	}
	for res := range stats {
		fmt.Println(res)
	}

}
func main() {
	b := BoomContent{}
	b.ArgsParse()
	stats := make(chan string, 10)
	b.MainController(stats)

}
