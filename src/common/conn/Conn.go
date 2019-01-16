package conn

import "sync"

var once sync.Once
var maxPoolSize = 10
