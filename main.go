package main

import (
	"fmt"
)

func main() {
	parseConsoleArgs()

	connectDB()

	if migrateFlag {
		migrate()
		fmt.Printf(SuccessColor, "All tables migrated successfully")
	}
	fmt.Printf(NoticeColor, "App loaded successfully")

	endPoint()

}
