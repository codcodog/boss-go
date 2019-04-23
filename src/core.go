package boss

func Run() {
	defer sqliteClose()
	defer redisClose()

	getArea()
	getBusiness()
	getJobList()
	getJD()
}
