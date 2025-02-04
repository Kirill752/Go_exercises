package main

func joinStr(sep string, elem ...string) string {
	res := []byte{}
	for _, el := range elem {
		res = append(res, el...)
		res = append(res, sep...)
	}
	res = res[:len(res)-1]
	return string(res)
}
