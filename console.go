package main

import "flag"

const (
	InfoColor    = "\033[1;34m%s\033[0m\n"
	SuccessColor = "\033[1;32m%s\033[0m\n"
	NoticeColor  = "\033[1;36m%s\033[0m\n"
	WarningColor = "\033[1;33m%s\033[0m\n"
	ErrorColor   = "\033[1;31m%s\033[0m\n"
	DebugColor   = "\033[0;36m%s\033[0m\n"
)

func parseConsoleArgs() {
	flag.BoolVar(&migrateFlag, "migrate", false, "add migrate option for migrate tables")
	flag.BoolVar(&dbSeedFlag, "seed", false, "add this option for fill table some fake data")
	flag.Parse()
}
