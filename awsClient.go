package main

import (
  "github.com/sirupsen/logrus"
  "os"
)

var logger = logrus.New()

func init() {
  logger.Formatter = new(logrus.JSONFormatter)

  // Only log the warning severity or above.
  //log.SetLevel(log.WarnLevel)
  logger.Level = logrus.DebugLevel
}

func main() {
  app := initApp()
  app.Run(os.Args)
}
