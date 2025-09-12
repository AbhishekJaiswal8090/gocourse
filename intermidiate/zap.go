package main 

import ("log" 
"go.uber.org/zap")


func zapLogger(){
   logger,err:=zap.NewProduction()
   if err!=nil{
	log.Println("Error in initialisisng zap logger")
   }

   defer logger.Sync()
   logger.Info(("this is an info message"))

   logger.Info("User logged in",zap.String("username","Abhi"),zap.String("method","GET"))
   
}