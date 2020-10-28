// https://geektutu.com/post/quick-gomock.html
package mock

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

/*
这个测试用例有2个目的:
一是使用 ctrl.Finish() 断言 DB.Get() 被是否被调用，如果没有被调用，后续的 mock 就失去了意义；
二是测试方法 GetFromDB() 的逻辑是否正确(如果 DB.Get() 返回 error，那么 GetFromDB() 返回 -1)。

NewMockDB() 的定义在 db_test.go 中，由 mockgen 自动生成。
*/
// go test . -cover -v
func TestGetFromDB(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(controller)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	//m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	//m.EXPECT().Get(gomock.Any()).Return(630, nil)
	//m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	//m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))

	/*
	m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
		if key == "Sam" {
			return 630, nil
		}
		return 0, errors.New("not exist")
	})
	*/

	//m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)
	//m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).MaxTimes(2)
	//m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).MinTimes(2)
	//m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).AnyTimes()
	//
	//Times() 断言 Mock 方法被调用的次数。
	//MaxTimes() 最大次数。
	//MinTimes() 最小次数。
	//AnyTimes() 任意次数（包括 0 次）

	// 调用顺序(InOrder)
	//o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	//o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	//gomock.InOrder(o1, o2)



	if v := GetFromDB(m,"Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
