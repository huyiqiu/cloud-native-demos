package main

import "coffee-device-plugin/pkg/coffee"
import log "github.com/sirupsen/logrus"


func main() {
    plugin := coffee.NewCoffeePulgin()

    if err := plugin.Serve(); err != nil {
        log.Fatal("cannot serve: ", err)
    }

	if err := coffee.RegisterWithKubelet(); err != nil {
		log.Fatal("cannot register to kubelet: ", err)
	}
    // Keep the plugin running
    select {}
}