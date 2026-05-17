package service

import "bishe/dao/es"

func ServiceInitEs() (err error) {
	return es.InitEs()
}
