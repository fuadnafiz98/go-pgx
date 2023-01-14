package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	// "strconv"
	"time"
)

var letters = []rune("abcdefghijklkmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genRandomUsername(length int) string {
	str := make([]rune, length)

	for index := range str {
		str[index] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}

func main() {

	// don't forget this
	rand.Seed(time.Now().UnixNano())

	// Create 5000 tags
	/**
	INSERT INTO tag(tag_name) values('');
	file, err := os.Create("./gen/sql/tag.sql")
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i < 5000; i++ {
		sql := "INSERT INTO tag(tag_name) values(" + "'" + genRandomUsername(rand.Intn(100-10)+10) + "');"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	/**/

	// Create 1000 writer
	/**
	file, err := os.Create("./gen/sql/writer.sql")
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 1000; i++ {
		username := genRandomUsername(rand.Intn(90-5) + 5)
		sql := "INSERT INTO writer(writer_username, writer_email) values(" +
			"'" + username + "'," +
			"'" + username + "@gmail.com'" +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	defer file.Close()
	/**/

	// Create 1000 editor
	/**
	file, err := os.Create("./gen/sql/editor.sql")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 1000; i++ {
		username := genRandomUsername(rand.Intn(90-5) + 5)
		sql := "INSERT INTO editor(editor_username, editor_email) values(" +
			"'" + username + "'," +
			"'" + username + "@gmail.com'" +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	defer file.Close()
	/**/

	// Create 1000 webmaster
	/**
	file, err := os.Create("./gen/sql/webmaster.sql")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 1000; i++ {
		username := genRandomUsername(rand.Intn(90-5) + 5)
		sql := "INSERT INTO webmaster(webmaster_username, webmaster_email) values(" +
			"'" + username + "'," +
			"'" + username + "@gmail.com'" +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	defer file.Close()
	/**/

	// Create 10000 websites
	// 20% webmaster will have 80% websites
	/**
	file, err := os.Create("./gen/sql/website.sql")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 8000; i++ {
		url := genRandomUsername(rand.Intn(100-5) + 5)
		sql := "INSERT INTO website(website_url, webmaster_id) values(" +
			"'https://" + url + ".com'," +
			strconv.Itoa(rand.Intn(200-1)+1) +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	for i := 0; i < 2000; i++ {
		url := genRandomUsername(rand.Intn(100-5) + 5)
		sql := "INSERT INTO website(website_url, webmaster_id) values(" +
			"'https://" + url + ".com'," +
			strconv.Itoa(rand.Intn(999-201)+201) +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	defer file.Close()
	/**/

	// Create 100 host_admin
	/**
	file, err := os.Create("./gen/sql/hosting_admin.sql")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 1000; i++ {
		username := genRandomUsername(rand.Intn(90-5) + 5)
		sql := "INSERT INTO hosting_admin(hosting_admin_username, hosting_admin_email) values(" +
			"'" + username + "'," +
			"'" + username + "@gmail.com'" +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	defer file.Close()
	/**/

	// Create 1000 hosting
	// 20% hosting_admin will have 80% hosting
	/**
	file, err := os.Create("./gen/sql/hosting.sql")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 800; i++ {
		name := genRandomUsername(rand.Intn(100-5) + 5)
		sql := "INSERT INTO hosting(hosting_name, hosting_admin_id) values(" +
			"'" + name + "'," +
			strconv.Itoa(rand.Intn(200-1)+1) +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	for i := 0; i < 200; i++ {
		name := genRandomUsername(rand.Intn(100-5) + 5)
		sql := "INSERT INTO hosting(hosting_name, hosting_admin_id) values(" +
			"'" + name + "'," +
			strconv.Itoa(rand.Intn(200-1)+1) +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	defer file.Close()
	/**/

	// Create 100,000 Blogs
	/**/
	file, err := os.Create("./gen/sql/blog.sql")
	if err != nil {
		fmt.Println(err)
	}

	status := [5]string{"requested", "review", "approved", "published", "deleted"}

	for i := 1; i <= 80000; i++ {
		name := genRandomUsername(rand.Intn(100-10) + 10)
		sql := "INSERT INTO blog(blog_title, blog_url, blog_status, blog_writer_id, blog_editor_id, blog_website_id) values(" +
			"'" + name + "'," +
			"'" + name + "'," +
			"'" + status[i%5] + "'," +
			strconv.Itoa(rand.Intn(200-1)+1) + "," +
			strconv.Itoa(rand.Intn(200-1)+1) + "," +
			strconv.Itoa(rand.Intn(200-1)+1) +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	for i := 1; i <= 20000; i++ {
		name := genRandomUsername(rand.Intn(100-10) + 10)
		sql := "INSERT INTO blog(blog_title, blog_url, blog_status, blog_writer_id, blog_editor_id, blog_website_id) values(" +
			"'" + name + "'," +
			"'" + name + "'," +
			"'" + status[i%5] + "'," +
			strconv.Itoa(rand.Intn(200-1)+1) + "," +
			strconv.Itoa(rand.Intn(200-1)+1) + "," +
			strconv.Itoa(rand.Intn(200-1)+1) +
			");"
		// fmt.Println(sql)
		fmt.Fprintln(file, sql)
	}
	defer file.Close()
	/**/

}
