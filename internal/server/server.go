package server
 
func Start() {
  router := setRouter()
 
  router.Run(":9000")
}