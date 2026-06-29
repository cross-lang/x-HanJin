// Package tasks 提供定时任务和周期任务的调度功能，
// 基于 robfig/cron 库实现，支持秒级精度的 cron 表达式。
package tasks

import (
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"x-HanJin/pkg/log"
)

// StartPeriodicTask 启动周期任务，支持指定首次执行时间。
//
// 参数:
//   - interval: 执行间隔，支持 Go duration 格式（如 "1h"、"30m"）
//   - startTime: 首次执行时间，格式为 "HH:MM"（空字符串表示立即执行）
//   - cmd: 任务执行函数
//
// 注意：此函数会阻塞调用 goroutine。
func StartPeriodicTask(interval string, startTime string, cmd func()) {
	// 解析间隔时间
	duration, err := time.ParseDuration(interval)
	if err != nil {
		duration, _ = time.ParseDuration(interval + "h")
	}
	log.Info("周期任务间隔已解析", zap.String("duration", duration.String()))

	// 计算距离开始时间的延迟
	var startDelay time.Duration
	if startTime != "" {
		now := time.Now()
		startTimeParts := strings.Split(startTime, ":")
		if len(startTimeParts) == 2 {
			hour, _ := strconv.Atoi(startTimeParts[0])
			minute, _ := strconv.Atoi(startTimeParts[1])

			startToday := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
			if startToday.Before(now) {
				startToday = startToday.Add(24 * time.Hour)
			}
			startDelay = startToday.Sub(now)
		}
	}

	log.Info("周期任务配置",
		zap.String("interval", duration.String()),
		zap.String("startTime", startTime),
		zap.Duration("startDelay", startDelay),
	)

	c := cron.New(cron.WithSeconds())

	if startDelay > 0 {
		log.Info("周期任务将在延迟后首次执行", zap.Duration("delay", startDelay))
		time.AfterFunc(startDelay, func() {
			log.Info("周期任务首次执行", zap.String("time", time.Now().Format("2006-01-02 15:04:05")))
			cmd()

			_, err = c.AddFunc("@every "+duration.String(), func() {
				log.Info("周期任务执行", zap.String("time", time.Now().Format("2006-01-02 15:04:05")))
				cmd()
			})
			if err != nil {
				log.Error("添加周期任务失败", zap.Error(err))
				return
			}
			log.Info("周期任务已调度", zap.String("interval", duration.String()))
		})
	} else {
		log.Info("周期任务立即开始", zap.String("interval", duration.String()))
		_, err = c.AddFunc("@every "+duration.String(), func() {
			log.Info("周期任务执行", zap.String("time", time.Now().Format("2006-01-02 15:04:05")))
			cmd()
		})
		if err != nil {
			log.Error("添加周期任务失败", zap.Error(err))
			return
		}
	}

	c.Start()
	log.Info("周期任务已启动")
	select {} // 阻塞主函数，防止程序退出
}

// StartScheduledTask 启动定时任务（支持秒级精度的 cron 表达式）。
//
// spec 格式（6 位）：
//
//	秒 分 时 日 月 星期
//
// 示例:
//   - "30 0 0 * * *"     每天 0:00:30
//   - "0 0 9 * * *"       每天 9:00:00
//   - "0 0 9 * * 1-5"     工作日 9:00:00
//   - "0 0 0 1 * *"       每月 1 号 0:00:00
//
// 注意：此函数会阻塞调用 goroutine。
func StartScheduledTask(spec string, cmd func()) {
	c := cron.New(cron.WithSeconds())

	entryID, err := c.AddFunc(spec, func() {
		log.Info("定时任务执行", zap.String("time", time.Now().Format("2006-01-02 15:04:05.000")))
		cmd()
	})
	if err != nil {
		log.Error("定时任务创建失败", zap.Error(err), zap.Int("entryID", int(entryID)))
		return
	}
	log.Info("定时任务创建成功", zap.Int("entryID", int(entryID)), zap.String("spec", spec))

	c.Start()
	log.Info("定时任务已启动", zap.Int("entryID", int(entryID)))
	select {} // 阻塞主函数，防止程序退出
}
