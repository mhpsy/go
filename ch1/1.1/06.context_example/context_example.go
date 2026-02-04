package main

import (
	"context"
	"fmt"
	"time"
)

// 模拟耗时操作：睡眠指定秒数
func slowTask(name string, duration time.Duration) string {
	time.Sleep(duration)
	return fmt.Sprintf("%s 完成！耗时 %v", name, duration)
}

func example1() {
	fmt.Println("========== 例子1：无超时 ==========")
	// 创建一个永不超时的 context
	ctx := context.Background()

	result := slowTask("任务1", 2*time.Second)
	fmt.Println(result)

	// 检查是否被取消（不会被取消）
	fmt.Printf("被取消了吗？%v\n\n", ctx.Err())
}

func example2() {
	fmt.Println("========== 例子2：有超时，任务完成 ==========")
	// 创建 3 秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 任务只需 2 秒，在超时前完成
	result := slowTask("任务2", 2*time.Second)
	fmt.Println(result)

	// 检查状态
	fmt.Printf("被取消了吗？%v\n", ctx.Err())
	fmt.Printf("还活着吗？%v\n\n", ctx.Done())
}

func example3() {
	fmt.Println("========== 例子3：有超时，任务超时 ==========")
	// 创建 1 秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 模拟：启动一个 goroutine 执行任务
	done := make(chan string)
	go func() {
		result := slowTask("任务3", 3*time.Second)
		done <- result
	}()

	// 等待任务完成或超时
	select {
	case result := <-done:
		// 任务在超时前完成
		fmt.Println(result)
	case <-ctx.Done():
		// context 超时了
		fmt.Println("❌ 任务超时了！被中止了")
		fmt.Printf("错误原因：%v\n\n", ctx.Err())
	}
}

func example4() {
	fmt.Println("========== 例子4：手动取消（不等超时）==========")
	// 创建 10 秒超时的 context（但我们会提前取消）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 启动一个慢任务
	done := make(chan string)
	go func() {
		result := slowTask("任务4", 5*time.Second)
		done <- result
	}()

	// 立即取消（不等任务完成）
	time.Sleep(1 * time.Second)
	cancel() // 🔴 主动取消，不用等 10 秒超时

	select {
	case result := <-done:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("❌ 任务被主动取消了！")
		fmt.Printf("错误原因：%v\n\n", ctx.Err())
	}
}

func example5() {
	fmt.Println("========== 例子5：Context 链式传递 ==========")
	// 父 context
	parentCtx, parentCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer parentCancel()

	// 子 context（继承父的超时）
	childCtx, childCancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer childCancel()

	// 如果父被取消，子也自动被取消
	fmt.Println("父超时：5秒，子超时：3秒")
	fmt.Println("子会在 3 秒时超时（更早的那个）")

	done := make(chan string)
	go func() {
		result := slowTask("子任务", 4*time.Second)
		done <- result
	}()

	select {
	case result := <-done:
		fmt.Println(result)
	case <-childCtx.Done():
		fmt.Println("❌ 子 context 超时了！")
		fmt.Printf("错误原因：%v\n\n", childCtx.Err())
	}
}

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()

	fmt.Println("========== Context 核心要点 ==========")
	fmt.Println("1. context.Background() = 无限期，永不超时")
	fmt.Println("2. WithTimeout() = 添加超时时间")
	fmt.Println("3. cancel() = 手动取消（不必等超时）")
	fmt.Println("4. <-ctx.Done() = 监听是否被取消/超时")
	fmt.Println("5. ctx.Err() = 获取取消原因")
	fmt.Println("6. defer cancel() = 必须调用，释放资源")
}
