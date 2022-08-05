package main

import "io"

// io 中有:
// 写入器 Writer
// 关闭器 Closer
// 写入关闭器 WriterCloser，嵌入了 Writer，Closer 的接口

type device struct {
}

func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func main() {
	var wc io.WriteCloser = new(device)
	wc.Write(nil)
	wc.Close()
	var writeOnly io.Writer = new(device)
	writeOnly.Write(nil)
}
