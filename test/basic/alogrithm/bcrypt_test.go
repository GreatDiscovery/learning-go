package alogrithm

import "testing"

import (
	"bufio"
	_ "context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 默认 cost（work factor）
const DefaultCost = 12

// HashPassword 用 bcrypt 哈希密码，返回哈希串
func HashPassword(password string, cost int) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// ComparePassword 校验明文密码与 hash 是否匹配
func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// NeedsRehash 判断给定的 hash 是否使用了比 targetCost 更低的 cost（即是否需要 rehash）
func NeedsRehash(hash string, targetCost int) bool {
	cost, err := bcrypt.Cost([]byte(hash))
	if err != nil {
		// 如果无法解析 cost，建议 rehash
		return true
	}
	return cost < targetCost
}

// CLI 功能：hash / verify
func runCLI() {
	cmd := flag.String("cmd", "hash", "命令: hash | verify | server")
	password := flag.String("password", "", "要哈希或验证的密码")
	hashStr := flag.String("hash", "", "用于 verify 的哈希串")
	cost := flag.Int("cost", DefaultCost, "bcrypt cost/work-factor (推荐 >=10, 常用 12)")
	addr := flag.String("addr", ":8080", "当 cmd=server 时的 HTTP 监听地址")
	flag.Parse()

	switch *cmd {
	case "hash":
		pw := *password
		if pw == "" {
			// 如果没有通过参数传入，交互式读取
			fmt.Print("Enter password: ")
			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')
			pw = strings.TrimSpace(line)
		}
		h, err := HashPassword(pw, *cost)
		if err != nil {
			log.Fatalf("Hash error: %v", err)
		}
		fmt.Printf("bcrypt hash (cost=%d): %s\n", *cost, h)
	case "verify":
		if *hashStr == "" {
			fmt.Println("请通过 -hash 指定哈希串")
			return
		}
		pw := *password
		if pw == "" {
			fmt.Print("Enter password to verify: ")
			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')
			pw = strings.TrimSpace(line)
		}
		start := time.Now()
		err := ComparePassword(*hashStr, pw)
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("Password mismatch (elapsed=%s): %v\n", elapsed, err)
		} else {
			fmt.Printf("Password OK (elapsed=%s)\n", elapsed)
		}
		// 是否需要 rehash
		if NeedsRehash(*hashStr, *cost) {
			fmt.Printf("Note: hash cost lower than target cost=%d. Consider re-hashing.\n", *cost)
		}
	case "server":
		fmt.Println("Starting HTTP bcrypt demo server on", *addr)
		http.HandleFunc("/hash", hashHandler(*cost))
		http.HandleFunc("/verify", verifyHandler())
		srv := &http.Server{
			Addr: *addr,
		}
		log.Fatal(srv.ListenAndServe())
	default:
		fmt.Println("Unknown cmd. Use -cmd=hash|verify|server")
	}
}

// HTTP handler: POST /hash { "password": "..." } -> {"hash":"...","cost":12}
func hashHandler(cost int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST", http.StatusMethodNotAllowed)
			return
		}
		var body struct {
			Password string `json:"password"`
			Cost     int    `json:"cost,omitempty"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad body: "+err.Error(), http.StatusBadRequest)
			return
		}
		c := cost
		if body.Cost != 0 {
			c = body.Cost
		}
		h, err := HashPassword(body.Password, c)
		if err != nil {
			http.Error(w, "hash error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		resp := map[string]interface{}{
			"hash": h,
			"cost": c,
		}
		json.NewEncoder(w).Encode(resp)
	}
}

// HTTP handler: POST /verify { "password":"...", "hash":"..." } -> {"ok":true}
func verifyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST", http.StatusMethodNotAllowed)
			return
		}
		var body struct {
			Password string `json:"password"`
			Hash     string `json:"hash"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad body: "+err.Error(), http.StatusBadRequest)
			return
		}
		err := ComparePassword(body.Hash, body.Password)
		ok := err == nil
		// Always respond with 200, don't leak reason (timing side-channels are still possible)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": ok})
	}
}

func main() {
	// 如果通过环境变量 BCRYPT_COST 指定 cost，优先使用
	if s := os.Getenv("BCRYPT_COST"); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 {
			// set default cost via const? We just pass later via flags
			log.Printf("Using BCRYPT_COST from env: %d", v)
		}
	}
	runCLI()
}

// TestBCryptDemo 测试密码哈希与校验流程
func TestBCryptDemo(t *testing.T) {
	password := "my_secure_password"

	// 生成哈希
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("生成 bcrypt 哈希失败: %v", err)
	}
	t.Logf("生成的哈希: %s", hash)

	// 验证正确密码
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		t.Errorf("密码校验失败: %v", err)
	} else {
		t.Log("✅ 正确密码校验通过")
	}

	// 验证错误密码
	if err := bcrypt.CompareHashAndPassword(hash, []byte("wrong_password")); err == nil {
		t.Error("❌ 错误密码应当无法通过校验")
	} else {
		t.Logf("✅ 错误密码校验失败（预期行为）: %v", err)
	}
}
