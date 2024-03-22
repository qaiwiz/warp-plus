ticker := time.NewTicker(garbageCollectTime)
for {
	select {
	case <-ticker.C:
		rate.mu.Lock()
		currentTime := rate.timeNow()
		for addr, entry := range rate.table {
			if currentTime.Sub(entry.lastTime) > garbageCollectTime {
	
