/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var rev uint32
	for i := 0; i < 32 && num > 0; i++ {
		rev |= num & 1 << (31 - i)
		num >>= 1
	}
	return rev
}

// @lc code=end

