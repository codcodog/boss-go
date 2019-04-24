package boss

func Run() {
	defer sqliteClose()
	defer redisClose()

	if !isCachedArea() {
		getArea()
	}
	if !isCachedBusiness() {
		getBusiness()
	}
	getJobList()
	getJD()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
