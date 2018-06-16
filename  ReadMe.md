# Service Complier  
  

### Workflow :  
![enter image description here](https://i.imgur.com/NJX8cWV.png)  
  
  
  
### Quy trình hoạt động  
**Input**:   
      + Source : string , nội dung code  
      + Language: enum, ngôn ngữ sử dụng  
  
  
  
**Output**:  
      + Status : boolean, trạng thái kết quả trả về( true or false)  
      + Message: string , nội dung lỗi hoặc kết quả của đoạn code  
        
  
  
**Các bước hoạt động**:  
  
  
B1: Nhận request và kiểm tra dữ liệu  
  
  
B2: Chạy đoạn code với format theo từng loại ngôn ngữ  trong cotainer tương ứng( tối đa timeout là 3s)
  
  
B3: GetLog của đoạn code và so sánh với kết quả mong muốn(có thể bỏ so sánh)  
  
  
B4: Trả về kết quả


Cron job sẽ tìm kiếm các file cũ(file tồn tại quá 1 tiếng để delete)


  