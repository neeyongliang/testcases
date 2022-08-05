package moretypes

import (
	"golang-tour/mypic"
)

func MyPic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := range image {
		image[i] = make([]uint8, dx)
		for j := range image[i] {
			image[i][j] = uint8((i+j)/2)
		}
	}
	return image
}

func ExercieseMyPic() {
	mypic.Show(MyPic)
}

// 验证方法，新建 test.html, 编辑内容：
// <img src="data:image/png;base64, 不带IMAGE:的输出结果">
// 然后在浏览器中打开此文件, 例如：
