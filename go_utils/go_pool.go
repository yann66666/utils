// Author: yann
// Date: 2021/9/27 3:21 下午
// Desc:

package utils

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 如果传入任务超时返回的错误信息
var ErrScheduleTimeout = fmt.Errorf("schedule error: timed out")

// 线程池模型
type GoPool struct {
	sem  chan struct{} //线程池队列
	work chan func()   //任务队列
	stop chan int
	min  uint32 //最小线程数数量
	lock sync.Mutex
}

//GoroutinePoolOptions 协程池初始化参数
type GoroutinePoolOptions struct {
	MaxSize   uint32 //MaxSize 协程最大数量
	QueueSize uint32 //QueueSize 队列
	CSize     uint32 //CSize 当前/最小协程数量
}

func (g *GoPool) Close() {
	close(g.stop)
	close(g.work)
	close(g.sem)
}

//***************************************************
//Description : 初始化线程池
//param :       GoroutinePoolOptions选项参数
//return :      线程池
//***************************************************
func NewPool(opt *GoroutinePoolOptions) *GoPool {
	if opt.CSize <= 0 {
		panic("The current number of threads must be greater than or equal to zero")
	}
	if opt.QueueSize <= 0 {
		panic("The task queue size must be greater than zero")
	}
	if opt.CSize > opt.MaxSize {
		panic("The current number of threads must be less than the maximum number of threads")
	}
	p := &GoPool{
		sem:  make(chan struct{}, opt.MaxSize),
		work: make(chan func(), opt.QueueSize),
		min:  opt.CSize,
		stop: make(chan int),
	}
	for i := uint32(0); i < opt.CSize; i++ {
		p.sem <- struct{}{}
		go p.worker(func() {})
	}

	return p
}

//***************************************************
//Description : 传入任务
//param :       任务句柄
//***************************************************
func (p *GoPool) Schedule(task func()) {
	_ = p.schedule(task, nil)
}

//***************************************************
//Description : 传入任务设置超时
//param :       超时时间
//param :       任务句柄
//return :      如果任务队列和协程队列均满一直无法写入,超过时间返回超时错误
//***************************************************
func (p *GoPool) ScheduleTimeout(timeout time.Duration, task func()) error {
	return p.schedule(task, time.After(timeout))
}

func (p *GoPool) schedule(task func(), timeout <-chan time.Time) error {
	select {
	case <-p.stop:
		return nil
	case <-timeout:
		return ErrScheduleTimeout
	case p.work <- task: //如果任务队列未满传入任务
		return nil
	case p.sem <- struct{}{}: //如果任务队列已满,说明做的速度慢则创建协程
		go p.worker(task)
		return nil

	}
}

//线程工作
func (p *GoPool) worker(task func()) {
	defer func() {
		<-p.sem
	}()

	task()
	//循环获取任务, 当超时后判断当前线程数量是否大于最小线程数量, 大于则退出本线程
	for {
		select {
		case <-p.stop:
			return
		case f := <-p.work:
			f()
		case <-time.After(time.Minute):
			p.lock.Lock()
			if uint32(len(p.sem)) > atomic.LoadUint32(&p.min) {
				p.lock.Unlock()
				return
			}
			p.lock.Unlock()
		}
	}
}

