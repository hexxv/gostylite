package http

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}


/*
net.Listen("tcp", "127.0.0.1:0")：监听一个未被占用的端口，并返回 Listener。
调用 http.Serve(ln, nil) 启动 http 服务。
使用 http.Get 发起一个 Get 请求，检查返回值是否正确。
尽量不对 http 和 net 库使用 mock，这样可以覆盖较为真实的场景。
*/
func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}

/*
针对 http 开发的场景，使用标准库 net/http/httptest 进行测试更为高效。
使用 httptest 模拟请求对象(req)和响应对象(w)，达到了相同的目的。
*/
func TestConn2(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}