断点续传网络与任务管理
client, err := GetClient("http://127.0.0.1:8080")
if err != nil {
	panic(err)
}

client.AddUpload("./file.tar")