package codefights

func rotateImage(image [][]int) [][]int {
	l := len(image)

	for layer := 0; layer < l/2; layer++ {
		start := layer
		end := l - 1 - layer

		for i := start; i < end; i++ {
			offset := i - start
			top := image[start][i]
			right := image[i][end]
			bottom := image[end][end-offset]
			left := image[end-offset][start]

			image[start][i], image[i][end], image[end][end-offset], image[end-offset][start] = left, top, right, bottom
		}
	}

	return image
}
