# Service Complier

### Workflow :
![enter image description here](https://lh3.googleusercontent.com/oGd5YcyQ0kqHhRTLaUyGt0z6P0KDik_AeH2M_SamRD6mojHGy4MZQJIVt6na8IObdsJ_cjHYQpdrchLqIcgBFEvZXR5dOMsB0FUL6jbH1AalIgB7D6lKvW7LInWM9xwIf7rqw0-nj3VDtnfFcJh3zOmo_mzeoKg_qLxu4cj7tWfYb47YjmS4Pi0st5WTTG94-WTzfrersm8YNI1F9EhOAaCSxR-mMS7GDwG2q-utIJzJKhdP6LDzXYJ45oW6Q1kK4rrpFCHosA34p9WL6_GuVdWjaFsVVeoqSlObfITm3-_g257wgCE9F-y7eMKrZc4-eBMxj28dyOEVPsTWLb22B1MjeMZp6S2V2LnIpDkMCmfNJjHBicrI3l30IcI76YsKVVrLDVd6JPjtqKdPG1lyDwzd6-LhCsdb9kV4SpGlR_ijJTkUZj5QwF236Z6i818oUEwPeQkZ8TbKvFoeytgl1W2fMMlNLZmvF5A3gNtO2kdo9gI0wIMFcuw_DU5hQJMZvSBb_QGHF8jVaUmgdvASOKqgwsesyro30m2B0JtpUviHKOTxE-xa4hCroMOqFzTupsPe966i09kCv6qE3VgpJym1Ks_ljW47GfkOTA=w1223-h626-no)

### Quy trình hoạt động
**Input**: 
		+ Source : string , nội dung code
		+ Language: enum, ngôn ngữ sử dụng
**Output**:
		+ Status : boolean, trạng thái kết quả trả về( true or false)
		+ Message: string , nội dung lỗi hoặc kết quả của đoạn code
		
**Các bước hoạt động**:
B1: Nhận request và kiểm tra dữ liệu
B2: Tạo 1 container chạy đoạn code với format theo từng loại ngôn ngữ
B3: GetLog của đoạn code và so sánh với kết quả mong muốn(có thể bỏ so sánh)
B4: Trả về kết quả
