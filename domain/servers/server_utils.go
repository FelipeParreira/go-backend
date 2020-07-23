package servers

func Filter(vs []Info, f func(Info) bool) []Info {
	vsf := make([]Info, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
