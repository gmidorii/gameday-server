package main

import (
	"html/template"
	"log"
	"net/http"
)

// Result is html response struct
type Result struct {
	Message string
}

func outernalHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s /outernal\n", r.Method)
	tmpl, err := template.ParseFiles("./template/outernal.html")
	if err != nil {
		failed(err, w)
		return
	}
	resp, err := http.Get(cfg.OuternalURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, Result{Message: "【失敗】外部接続できません。設定を確認してください"})
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, Result{Message: "【失敗】外部接続に失敗しました。設定を確認してください"})
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Result{Message: "【成功】外部接続できました！！"})
}
