package service

import (
	"fmt"
	"github.com/khorevaa/ras-client/serialize"
)

func (s service) getFromCache(key string) (interface{}, bool) {
	//fmt.Sprintf("%s.clusters", ctt.App.Name)
	return s.cache.Get(key)
}

func (s service) getCacheClusters(appName string) ([]*serialize.ClusterInfo, bool) {

	if clusters, ok := s.cache.Get(fmt.Sprintf(clustersTpl, appName)); ok {
		return clusters.([]*serialize.ClusterInfo), ok
	}
	return nil, false
}

func (s service) setCacheClusters(appName string, clusters []*serialize.ClusterInfo) {
	s.cache.Set(fmt.Sprintf(clustersTpl, appName), clusters)
}

func (s service) getCacheInfobases(cluster string) (serialize.InfobaseSummaryList, bool) {

	if list, ok := s.cache.Get(fmt.Sprintf(infobasesTpl, cluster)); ok {
		return list.(serialize.InfobaseSummaryList), ok
	}
	return nil, false
}

func (s service) setCacheInfobases(cluster string, list serialize.InfobaseSummaryList) {
	s.cache.Set(fmt.Sprintf(infobasesTpl, cluster), list)
}

func (s service) clearCacheInfobases(cluster string) {
	s.cache.Clear(fmt.Sprintf(infobasesTpl, cluster))
}
