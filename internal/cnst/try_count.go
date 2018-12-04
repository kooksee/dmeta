package cnst

type tryCount struct {
	HttpProxy int
	NsqPost   int
	OssOpt    int
	CallBack  int
}

var TryCount = tryCount{
	HttpProxy: 10,
	NsqPost:   100,
	OssOpt:    100,
	CallBack:  50,
}
