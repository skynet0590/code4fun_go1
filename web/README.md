1. **Tài liệu sử dụng:**

* [html template](https://golang.org/pkg/html/template/)
2. **Bài tập về nhà**
> Như mọi người thấy trong file ./web/handlers/init.go dòng 24 mình chỉ parse duy nhất 1 file. 
> Trong thực tế project sẽ bao gồm rất nhiều file nằm trong 1 thư mục và code của chúng ta phải đảm bảo
> tìm kiếm tất cả các file trong thư mục đó. Vậy nên bài tập về nhà là viết một func. Biến đầu vào là đường 
> dẫn đến thư mục và tên đuôi file. Biến đầu ra là danh sách đường dẫn đến các files nằm trong thư mục đó có 
> đuôi file là là đuôi file đã truyền vào 