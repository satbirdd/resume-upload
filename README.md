断点续传客户端
```
client, err := resume_upload.NewClient("http://127.0.0.1:8084/files/", nil, nil)
if err != nil {
	panic(err)
}

done := make(chan struct{})
go func() {
	fp := "./file.rar"
	err := client.Upload(fp, done)
	if err != nil {
		log.Printf("[Resumable Upload]upload file %v failed，%v", fp, err)
	}
	os.Exit(1)
}()

<-done
fmt.Println("upload finished")
```