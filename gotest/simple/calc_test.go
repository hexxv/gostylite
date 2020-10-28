package simple

import (
	"fmt"
	"os"
	"testing"
)

// 最基础写法(不推荐)
func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

// 子测试
func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2,3) != 7 {
			t.Fail()
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fail()
		}
	})
}

/*
表格驱动测试:

所有用例的数据组织在切片 cases 中，看起来就像一张表，借助循环创建子测试。这样写的好处有：

新增用例非常简单，只需给 cases 新增一条测试数据即可。
测试代码可读性好，直观地能够看到每个子测试的参数和期待的返回值。
用例失败时，报错信息的格式比较统一，测试报告易于阅读。
*/
func TestMul2(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}

/*
帮助函数(helpers)

对一些重复的逻辑，抽取出来作为公共的帮助函数(helpers)，可以增加测试代码的可读性和可维护性。 借助帮助函数，可以让测试用例的主逻辑看起来更清晰。

例如，我们可以将创建子测试的逻辑抽取出来：
*/
type calcCase struct{ A, B, Expected int }

func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got",
			c.A, c.B, c.Expected, ans)
	}

}

func TestMul3(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, -3, -6})
	createMulTestCase(t, &calcCase{2, 0, 1}) // wrong case
}



/*
如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。
调用 m.Run() 触发所有测试用例的执行，并使用 os.Exit() 处理返回的状态码，如果不为0，说明有用例失败。
因此可以在调用 m.Run() 前后做一些额外的准备(setup)和回收(teardown)工作。
*/

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}


func TestAdd2(t *testing.T) {
	var a = make([]int, 0)
	a = append(a, 0,1,2,3,4,5,6,7,8,9)
	var b = a[1:5:5]
	t.Log(a)
	t.Log(b)
}