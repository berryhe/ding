package internal

// StrArrGroupAlg 字符串分组算法 字符串切片按n分组
func StrArrGroupAlg(strArr []string, n int) [][]string {
	saLen := len(strArr)

	if saLen == 0 {
		return nil
	}

	// flag 记录是否有余数
	count, flag := saLen/n, 0

	if saLen%n != 0 {
		count++
		flag = 1
	}

	res := make([][]string, count)

	for i := 0; i < count-flag; i++ {
		res[i] = strArr[i*n : i*n+n]
	}

	if flag != 0 {
		res[count-1] = strArr[n*(count-1):]
	}

	return res
}

// FlipByteSlice  翻转byte切片，字符串通用
func FlipByteSlice(bArr []byte) []byte {
	for i, j := 0, len(bArr)-1; i < j; i++ {
		bArr[i], bArr[j] = bArr[j], bArr[i]
		j--
	}
	return bArr
}
