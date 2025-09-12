package main

import ("github.com/sirupsen/logrus")

func logrus1(){
	//create instances of logrus
    log :=logrus.New()

	//set level for what kind of this logger will be use ex- info ,error ,debug
	log.SetLevel(logrus.DebugLevel)

	//now set the format whether is is kind of any simple text or json or anyhting else
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Debug("This is a degub message")
	log.Info("This is info message")
	log.Error("This is a error message")


	log.WithFields(logrus.Fields{
		"username ":"John Doe",
		"method ":"GET",
	}).Info("User logged in.")
}