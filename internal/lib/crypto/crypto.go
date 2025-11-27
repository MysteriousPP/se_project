package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

// 生成随机盐值（16 字节 = 128 位，安全足够，可调整长度）
func generateSalt() (string, error) {
	salt := make([]byte, 16)
	// 从加密安全的随机源读取字节（避免伪随机）
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}
	// 转为 base64 字符串（方便存储，16 字节转 base64 后为 24 字符）
	return base64.URLEncoding.EncodeToString(salt), nil
}

// SHA256WithSalt 加盐哈希：明文 + 盐值 拼接后计算 SHA256
// 返回格式：盐值:哈希值（用冒号分隔，方便后续验证）
func SHA256WithSalt(plainText string) (string, error) {
	// 1. 生成随机盐
	salt, err := generateSalt()
	if err != nil {
		return "", fmt.Errorf("生成盐值失败：%v", err)
	}

	// 2. 明文 + 盐值 拼接（也可使用 HMAC 算法，更规范，下文补充）
	combined := plainText + salt

	// 3. 计算 SHA256 哈希
	hash := sha256.New()
	hash.Write([]byte(combined))
	hashBytes := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashBytes)

	// 4. 返回“盐值:哈希值”（存储到数据库，后续验证需提取盐值）
	return fmt.Sprintf("%s:%s", salt, hashStr), nil
}

// VerifySHA256WithSalt 验证加盐哈希：用存储的盐值重新计算哈希，对比是否一致
func VerifySHA256WithSalt(plainText, storedHash string) (bool, error) {
	// 1. 从存储的字符串中拆分盐值和原始哈希（按冒号分割）
	sepIndex := strings.Index(storedHash, ":")
	if sepIndex == -1 {
		return false, fmt.Errorf("存储的哈希格式错误")
	}
	salt := storedHash[:sepIndex]
	originalHash := storedHash[sepIndex+1:]

	// 2. 用相同的盐值重新计算哈希
	combined := plainText + salt
	hash := sha256.New()
	hash.Write([]byte(combined))
	currentHash := hex.EncodeToString(hash.Sum(nil))

	// 3. 对比哈希值（一致则验证通过）
	return currentHash == originalHash, nil
}

// func main() {
// 	// 示例：密码存储与验证
// 	password := "12345678"

// 	// 1. 加密（存储到数据库的是 salt:hash 字符串）
// 	storedHash, err := SHA256WithSalt(password)
// 	if err != nil {
// 		fmt.Printf("加密失败：%v\n", err)
// 		return
// 	}
// 	fmt.Printf("存储到数据库的密文：%s\n", storedHash)

// 	// 2. 验证（用户登录时，用输入的密码验证）
// 	inputPassword := "12345678" // 正确密码
// 	ok, err := VerifySHA256WithSalt(inputPassword, storedHash)
// 	if err != nil {
// 		fmt.Printf("验证失败：%v\n", err)
// 		return
// 	}
// 	fmt.Printf("密码验证结果（正确密码）：%v\n", ok) // 输出 true

// 	// 验证错误密码
// 	inputPasswordWrong := "87654321"
// 	ok, _ = VerifySHA256WithSalt(inputPasswordWrong, storedHash)
// 	fmt.Printf("密码验证结果（错误密码）：%v\n", ok) // 输出 false
// }
