package os

import "os"

func CreateFolder(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirname, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
