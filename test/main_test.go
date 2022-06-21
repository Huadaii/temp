package test

import (
	"fmt"
	"log"
	"strings"
	"temp/queue"
	"testing"
)

func BenchmarkFormat(b *testing.B) {
	var sql strings.Builder
	for i := 0; i < b.N; i++ {
		fmt.Fprint(&sql, `select t1.u_user_id i_d, t1.realname realname,t2.mobile mobile, t2.createdate createdate, t1.grade, t1.name username, t2.is_initpwd is_init_pwd, t1.nickname, t3.head_url avatar_u_r_l, t1.es_school_id school_i_d, t1.es_school_name school_name,
	t1.es_class_id class_i_d, t1.es_class_name class_name, t1.student_no,t1.classnum
	from base.base_user t1 inner join base.base_userauth t2 on t1.u_user_id=t2.u_user_id inner join base.base_userext t3 on t1.u_user_id=t3.u_user_id where 1=1 `)
		fmt.Fprint(&sql, " and t1.u_user_id in ("+"c.String()"+")")
		fmt.Fprint(&sql, " and t1.realname='"+"c.String()"+"'")
		fmt.Fprint(&sql, " and t1.realname is not null ")
		_ = sql.String()
	}
}

func BenchmarkFormats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sql := `select t1.u_user_id i_d, t1.realname realname,t2.mobile mobile, t2.createdate createdate, t1.grade, t1.name username, t2.is_initpwd is_init_pwd, t1.nickname, t3.head_url avatar_u_r_l, t1.es_school_id school_i_d, t1.es_school_name school_name,
	t1.es_class_id class_i_d, t1.es_class_name class_name, t1.student_no,t1.classnum
	from base.base_user t1 inner join base.base_userauth t2 on t1.u_user_id=t2.u_user_id inner join base.base_userext t3 on t1.u_user_id=t3.u_user_id where 1=1 `
		sql += " and t1.u_user_id in (" + "c.String()" + ")"
		sql += " and t1.realname='" + "c.String()" + "'"
		sql += " and t1.realname is not null "
		_ = sql
	}
}

func BenchmarkFormatj(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sql := `select t1.u_user_id i_d, t1.realname realname,t2.mobile mobile, t2.createdate createdate, t1.grade, t1.name username, t2.is_initpwd is_init_pwd, t1.nickname, t3.head_url avatar_u_r_l, t1.es_school_id school_i_d, t1.es_school_name school_name,
	t1.es_class_id class_i_d, t1.es_class_name class_name, t1.student_no,t1.classnum
	from base.base_user t1 inner join base.base_userauth t2 on t1.u_user_id=t2.u_user_id inner join base.base_userext t3 on t1.u_user_id=t3.u_user_id where 1=1 `
		sql += " and t1.u_user_id in (" + "c.String()" + ")"
		sql += " and t1.realname='" + "c.String()" + "'"
		sql += " and t1.realname is not null "
		_ = sql
	}
}

var q queue.ItemQueue

func initQueue() {
	if q.Data == nil {
		q = queue.ItemQueue{}
		q.New()
	}
}

func TestIQueue_queue(t *testing.T) {
	initQueue()
	q.Enqueue("2")
	q.Enqueue("ask")
	log.Println("size", q.Size())
	item := q.Dequeue()
	log.Println("pop item", item)
	item = q.Dequeue()
	log.Println("pop item", item)
	item = q.Dequeue()
	log.Println("pop item", item)
	// for {
	// 	if !q.IsEmpty() {
	// 		item := q.Dequeue()
	// 		log.Println("pop item", item)
	// 	}
	// }

	log.Println("queue front item", q.Front())

	log.Println("size", q.Size())
	log.Println("isEmpty", q.IsEmpty())
}
