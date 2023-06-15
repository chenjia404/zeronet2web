package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chenjia404/zeronet2web/models"
	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db = make(map[string]string)

type Post struct {
	Post_id       int     `json:"post_id"`
	Title         string  `json:"title"`
	DatePublished float64 `json:"date_published"`
	Body          string  `json:"body"`
}

type UserData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Links       string `json:"links"`
	Post        []Post `json:"post"`
}

func setupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		blogs := make([]models.Blog, 0, 20)
		db.Limit(20).Find(&blogs)
		for _, blog := range blogs {
			fmt.Printf("Address:%s %s\n", blog.Address, blog.Title)
		}
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title":       "zeronet 2 web",
			"description": "显示全部zeronet博客",
			"blogs":       blogs,
		})
	})
	r.GET("/:address/", func(c *gin.Context) {
		address := c.Param("address")
		post_id := c.Query("post_id")
		postId, err := strconv.Atoi(post_id)
		if err != nil {
			postId = 0
		}
		jsonFile, err := os.Open(ZeroNetDataPath + address + "/data/data.json")
		if err != nil {
			fmt.Println("文件不存在，请查看该文件")
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var result UserData
		err = json.Unmarshal([]byte(byteValue), &result)
		if err == nil {
			result := db.Where(models.Blog{Address: address}).FirstOrCreate(&models.Blog{
				Title:       result.Title,
				Address:     address,
				Modified:    0,
				Description: "Zeronet",
			})
			if result.RowsAffected >= 1 {
				fmt.Printf("插入成功\n")
			} else {
				fmt.Printf("插入失败%s\n", address)
			}
		} else {
			fmt.Println(err)
		}

		fmt.Println(result.Title)
		fmt.Println(result.Description)
		fmt.Printf("文章数：%d\n", len(result.Post))
		fmt.Printf("postId:%d\n", postId)

		if postId == 0 {

			var buf bytes.Buffer
			if err := md.Convert([]byte(result.Description), &buf); err != nil {
				panic(err)
			}

			var linksBuf bytes.Buffer
			if err := md.Convert([]byte(result.Links), &buf); err != nil {
				panic(err)
			}
			description := strings.Replace(buf.String(), "http://127.0.0.1:43110/", ProxyHost, -1)
			links := strings.Replace(linksBuf.String(), "http://127.0.0.1:43110/", ProxyHost, -1)

			c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
				"address":     address,
				"proxy_host":  ProxyHost,
				"title":       result.Title,
				"description": template.HTML(description),
				"links":       template.HTML(links),
				"Post":        result.Post,
			})
		} else {
			for _, post := range result.Post {
				if post.Post_id == postId {
					var buf bytes.Buffer
					if err := md.Convert([]byte(post.Body), &buf); err != nil {
						panic(err)
					}

					body := strings.Replace(buf.String(), "http://127.0.0.1:43110/", ProxyHost, -1)
					c.HTML(http.StatusOK, "posts/post.tmpl", gin.H{
						"address":        address,
						"proxy_host":     ProxyHost,
						"title":          post.Title,
						"date_published": time.Unix(int64(post.DatePublished), 0).String(),
						"body":           template.HTML(body),
					})
					break
				}

			}

		}

		// c.String(http.StatusOK, "Hello %s %s", name, post_id)
	})

	return r
}

var ZeroNetDataPath = ""
var ProxyHost = ""

func main() {
	db, _ := gorm.Open(sqlite.Open("zeronet.blogs.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	db.AutoMigrate(&models.Blog{})
	db.Migrator().CreateIndex(&models.Blog{}, "id")
	db.Migrator().CreateIndex(&models.Blog{}, "id")

	var _ZeroNetDataPath = flag.String("dir", "", "ZeroNet Data Path")
	var _ProxyHost = flag.String("host", "http://127.0.0.1:43110/", "Proxy Host")
	var _port = flag.String("port", "20236", "web port")
	flag.Parse()
	ZeroNetDataPath = *_ZeroNetDataPath
	ProxyHost = *_ProxyHost
	fmt.Printf("ZeroNet Data Path:%s\n", ZeroNetDataPath)
	fmt.Printf("Proxy Host:%s\n", ProxyHost)

	r := setupRouter(db)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + *_port)
}
