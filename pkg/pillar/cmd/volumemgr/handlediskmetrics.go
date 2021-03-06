/*
 * Copyright (c) 2020. Zededa, Inc.
 * SPDX-License-Identifier: Apache-2.0
 */

package volumemgr

import (
	"fmt"
	"time"

	"github.com/lf-edge/eve/pkg/pillar/diskmetrics"
	"github.com/lf-edge/eve/pkg/pillar/flextimer"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/utils"
	"github.com/shirou/gopsutil/disk"
)

func publishDiskMetrics(ctx *volumemgrContext, statuses ...*types.DiskMetric) {
	for _, status := range statuses {
		key := status.Key()
		log.Debugf("publishDiskMetrics(%s)", key)
		pub := ctx.pubDiskMetric
		pub.Publish(key, *status)
		log.Debugf("publishDiskMetrics(%s) Done", key)
	}
}

func unpublishDiskMetrics(ctx *volumemgrContext, statuses ...*types.DiskMetric) {
	for _, status := range statuses {
		key := status.Key()
		log.Debugf("unpublishDiskMetrics(%s)", key)
		pub := ctx.pubDiskMetric
		c, _ := pub.Get(key)
		if c == nil {
			log.Errorf("unpublishDiskMetrics(%s) not found", key)
			continue
		}
		pub.Unpublish(key)
		log.Debugf("unpublishDiskMetrics(%s) Done", key)
	}
}

func lookupDiskMetric(ctx *volumemgrContext, key string) *types.DiskMetric {
	key = types.PathToKey(key)
	log.Debugf("lookupDiskMetric(%s)", key)
	pub := ctx.pubDiskMetric
	c, _ := pub.Get(key)
	if c == nil {
		log.Debugf("lookupDiskMetric(%s) not found", key)
		return nil
	}
	status := c.(types.DiskMetric)
	log.Debugf("lookupDiskMetric(%s) Done", key)
	return &status
}

func publishAppDiskMetrics(ctx *volumemgrContext, statuses ...*types.AppDiskMetric) {
	for _, status := range statuses {
		key := status.Key()
		log.Debugf("publishAppDiskMetrics(%s)", key)
		pub := ctx.pubAppDiskMetric
		pub.Publish(key, *status)
		log.Debugf("publishAppDiskMetrics(%s) Done", key)
	}
}

func unpublishAppDiskMetrics(ctx *volumemgrContext, statuses ...*types.AppDiskMetric) {
	for _, status := range statuses {
		key := status.Key()
		log.Debugf("unpublishAppDiskMetrics(%s)", key)
		pub := ctx.pubAppDiskMetric
		c, _ := pub.Get(key)
		if c == nil {
			log.Errorf("unpublishAppDiskMetrics(%s) not found", key)
			continue
		}
		pub.Unpublish(key)
		log.Debugf("unpublishAppDiskMetrics(%s) Done", key)
	}
}

func lookupAppDiskMetric(ctx *volumemgrContext, key string) *types.AppDiskMetric {
	key = types.PathToKey(key)
	log.Debugf("lookupAppDiskMetric(%s)", key)
	pub := ctx.pubAppDiskMetric
	c, _ := pub.Get(key)
	if c == nil {
		log.Debugf("lookupAppDiskMetric(%s) not found", key)
		return nil
	}
	status := c.(types.AppDiskMetric)
	log.Debugf("lookupAppDiskMetric(%s) Done", key)
	return &status
}

//diskMetricsTimerTask calculates and publishes disk metrics periodically
func diskMetricsTimerTask(ctx *volumemgrContext, handleChannel chan interface{}) {
	log.Infoln("starting report diskMetricsTimerTask timer task")
	createOrUpdateDiskMetrics(ctx)

	diskMetricInterval := time.Duration(ctx.globalConfig.GlobalValueInt(types.DiskScanMetricInterval)) * time.Second
	max := float64(diskMetricInterval)
	min := max * 0.3
	diskMetricTicker := flextimer.NewRangeTicker(time.Duration(min), time.Duration(max))
	// Return handle to caller
	handleChannel <- diskMetricTicker

	// Run a periodic timer so we always update StillRunning
	stillRunning := time.NewTicker(25 * time.Second)
	ctx.ps.StillRunning(diskMetricsAgentName, warningTime, errorTime)

	for {
		select {
		case <-diskMetricTicker.C:
			start := time.Now()
			createOrUpdateDiskMetrics(ctx)
			ctx.ps.CheckMaxTimeTopic(diskMetricsAgentName, "createOrUpdateDiskMetrics", start,
				warningTime, errorTime)

		case <-stillRunning.C:
		}
		ctx.ps.StillRunning(diskMetricsAgentName, warningTime, errorTime)
	}
}

//createOrUpdateDiskMetrics creates or updates metrics for all disks, mountpaths and volumeStatuses
func createOrUpdateDiskMetrics(ctx *volumemgrContext) {
	log.Infof("createOrUpdateDiskMetrics")
	var diskMetricList []*types.DiskMetric
	startPubTime := time.Now()

	disks := diskmetrics.FindDisksPartitions(log)
	for _, d := range disks {
		size, _ := diskmetrics.PartitionSize(log, d)
		log.Debugf("createOrUpdateDiskMetrics: Disk/partition %s size %d", d, size)
		var metric *types.DiskMetric
		metric = lookupDiskMetric(ctx, d)
		if metric == nil {
			log.Infof("createOrUpdateDiskMetrics: creating new DiskMetric for %s", d)
			metric = &(types.DiskMetric{DiskPath: d, IsDir: false})
		} else {
			log.Infof("createOrUpdateDiskMetrics: updating DiskMetric for %s", d)
		}
		metric.TotalBytes = size
		stat, err := disk.IOCounters(d)
		if err == nil {
			metric.ReadBytes = stat[d].ReadBytes
			metric.WriteBytes = stat[d].WriteBytes
			metric.ReadCount = stat[d].ReadCount
			metric.WriteCount = stat[d].WriteCount
		}
		// XXX do we have a mountpath? Combine with paths below if same?
		diskMetricList = append(diskMetricList, metric)
	}

	var persistUsage uint64
	for _, path := range types.ReportDiskPaths {
		u, err := disk.Usage(path)
		if err != nil {
			// Happens e.g., if we don't have a /persist
			log.Errorf("createOrUpdateDiskMetrics: disk.Usage: %s", err)
			continue
		}
		// We can not run diskmetrics.SizeFromDir("/persist") below in reportDirPaths, get the usage
		// data here for persistUsage
		if path == types.PersistDir {
			persistUsage = u.Used
		}
		log.Debugf("createOrUpdateDiskMetrics: Path %s total %d used %d free %d",
			path, u.Total, u.Used, u.Free)
		var metric *types.DiskMetric
		metric = lookupDiskMetric(ctx, path)
		if metric == nil {
			log.Infof("createOrUpdateDiskMetrics: creating new DiskMetric for %s", path)
			metric = &(types.DiskMetric{DiskPath: path, IsDir: true})
		} else {
			log.Infof("createOrUpdateDiskMetrics: updating DiskMetric for %s", path)
		}
		metric.TotalBytes = u.Total
		metric.UsedBytes = u.Used
		metric.FreeBytes = u.Free
		diskMetricList = append(diskMetricList, metric)
	}
	log.Debugf("createOrUpdateDiskMetrics: persistUsage %d, elapse sec %v", persistUsage, time.Since(startPubTime).Seconds())

	for _, path := range types.ReportDirPaths {
		usage := diskmetrics.SizeFromDir(log, path)
		log.Debugf("createOrUpdateDiskMetrics: ReportDirPath %s usage %d", path, usage)
		var metric *types.DiskMetric
		metric = lookupDiskMetric(ctx, path)
		if metric == nil {
			log.Infof("createOrUpdateDiskMetrics: creating new DiskMetric for %s", path)
			metric = &(types.DiskMetric{DiskPath: path, IsDir: true})
		} else {
			log.Infof("createOrUpdateDiskMetrics: updating DiskMetric for %s", path)
		}

		metric.UsedBytes = usage

		diskMetricList = append(diskMetricList, metric)
	}
	log.Debugf("createOrUpdateDiskMetrics: DirPaths in persist, elapse sec %v", time.Since(startPubTime).Seconds())

	for _, path := range types.AppPersistPaths {
		usage := diskmetrics.SizeFromDir(log, path)
		log.Debugf("createOrUpdateDiskMetrics: AppPersistPath %s usage %d", path, usage)
		var metric *types.DiskMetric
		metric = lookupDiskMetric(ctx, path)
		if metric == nil {
			log.Infof("createOrUpdateDiskMetrics: creating new DiskMetric for %s", path)
			metric = &(types.DiskMetric{DiskPath: path, IsDir: true})
		} else {
			log.Infof("createOrUpdateDiskMetrics: updating DiskMetric for %s", path)
		}

		metric.UsedBytes = usage

		diskMetricList = append(diskMetricList, metric)
	}
	publishDiskMetrics(ctx, diskMetricList...)
	for _, volumeStatus := range getAllVolumeStatus(ctx) {
		if err := createOrUpdateAppDiskMetrics(ctx, volumeStatus); err != nil {
			log.Errorf("CreateOrUpdateCommonDiskMetrics: exception while publishing diskmetric. %s", err.Error())
		}
	}
}

func createOrUpdateAppDiskMetrics(ctx *volumemgrContext, volumeStatus *types.VolumeStatus) error {
	log.Infof("createOrUpdateAppDiskMetrics(%s, %s)", volumeStatus.VolumeID, volumeStatus.FileLocation)
	actualSize, maxSize, diskType, dirtyFlag, err := utils.GetVolumeSize(log, volumeStatus.FileLocation)
	if err != nil {
		err = fmt.Errorf("createOrUpdateAppDiskMetrics(%s, %s): exception while getting volume size. %s",
			volumeStatus.VolumeID, volumeStatus.FileLocation, err)
		log.Error(err.Error())
		return err
	}
	appDiskMetric := types.AppDiskMetric{DiskPath: volumeStatus.FileLocation,
		ProvisionedBytes: maxSize,
		UsedBytes:        actualSize,
		DiskType:         diskType,
		Dirty:            dirtyFlag,
	}
	publishAppDiskMetrics(ctx, &appDiskMetric)
	return nil
}
