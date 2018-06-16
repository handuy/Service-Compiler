
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
  
  
B4: Trả về kết quả# Service Complier  
  
### Workflow :  
![enter image description here](https://lh3.googleusercontent.com/h9CW14p3YPgSa2TWVgrwmsEW6GMkakngn4IB936OFBJTm6nqlg_LMJKauWAGXrOD-HO1m8zSY7VHdldy3PHpZohF2PpIQDbVCZn2W9N51JNzDz8rM5GQs2X9QExsGw95SqDv8zRmS4PjlMRHzPV7QDFG_RWf9TL9OLTdbqa2p9CT1NACOX2CbpsH5VBiEjo2RhxUachJpA4_1yZWkkW3UH2ARS5C4wT0w62CXJWwulerfVT8h-dKKpxlXflwGuO0swgR-axEY7VcsreaRWzqLfrfJwy1w9Zd3_2uz-F24xswK9MA31SoNycNUN8Rw7J4WV1d2thY1Wik61Uq1U2h7sG9P_Vu9aTQ9S4O8MCm6vwdougRTD9SlKMWIPXNpCKLwcL62jlzK4taM1mWZtAxp--I874km9HOMP3jwDev_9rxtkzkh1uZkpvOEIwSa9jt1zN9oHF8sN91Y4mCszMFiTz38mZqGFYGPPezlggNglelxBZS3t8ux6ELOBxhHUfWe-cnYdpRRVwirR5FakoVu3OZSgTxWzdnT-Rlu-8JzNRvnMfGBI9VkEd3KAsFcEJ4zfcWqv-_L3cvaxZNb5c9jj4TOlWKE-0B0KydkQ=w1223-h626-no)  
  
  
  
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

**Các bước hoạt động**:  
  