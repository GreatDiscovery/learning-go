package k8s

import (
	bolt "go.etcd.io/bbolt"
	"path/filepath"
)

// /var/lib/docker/buildkit下面有containerdmeta.db、metadata_v2.db等db文件

func main() {
	dirname := ""
	db, _ := bolt.Open(filepath.Join(dirname, "meta.db"), 0644, nil)
	db.View(func(tx *bolt.Tx) error {
		return nil
	})
}
