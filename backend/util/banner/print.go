package banner

import "fmt"

// Print the application banner to the console
func Print() {
	fmt.Printf(`
 ____   ___   _        ___      ____   ____  __ __
|    \ /   \ | |      /  _]    |    \ /    ||  |  |
|  D  )     || |     /  [_     |  o  )  o  ||  |  |
|    /|  O  || |___ |    _]    |   _/|     ||  ~  |
|    \|     ||     ||   [_     |  |  |  _  ||___, |
|  .  \     ||     ||     |    |  |  |  |  ||     |
|__|\_|\___/ |_____||_____|    |__|  |__|__||____/
choose a role, get paid
alpha version - v0.0.1

`)
}
