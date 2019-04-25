package boss

func Run() {
	defer sqliteClose()
	defer redisClose()

	getArea()
	getBusiness()
	getJobList()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
