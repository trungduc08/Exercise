1. Các bước để  push lên github cơ bản"
    - Kiểm tra xem có sự thay đổi nào ko
        + Sử dụng git status: xem file đang ở trên branch nào? trạng thái của file:
            đã được staged hay chưa: nếu chưa được staged đến bước 2.
        + git add .
        + git commit -m "message in commit"
        + git push origin name_branch

        + note: muốn kiếm tra chính xác sự thay đổi sử dụng git diff

2. Sử dụng git --amend để sửa commit
    C1: 
    - Sau khi commit mà thấy chưa đủ cần phải sửa lại file hay thêm file vào, ...
    - sử dụng git commit --amend
    - sau đó push sử dụng: git push -f 
    C2: 
    - Sử dụng git reset --h để trở lại trước khi commit

3. Tạo branch mới
    - git checkout -b name_new_branch
    - để chuyển tới 1 branch: git checkout name_branch


4. Merge giữa các branch
    - Dùng git rebase(thường là từ các nhánh khác đến nhánh master)
        + ví dụ muốn rebase từ dev vào master
            B1: git checkout master
            B2: git pull origin master
            B3: git merge dev

    - Dùng git merge:(thường là từ master đến các nhánh khác)
        + ví dụ muốn merge từ master vào dev
            B1: git checkout dev
            B2: git pull origin dev
            B3: git merge master
        
5. So sánh git rebase và git merge 
    - Giống nhau: đều tích hợp những thay đổi từ 1 nhánh vào 1 nhánh khác
    - Khác nhau: 
        + merge: lấy tất cả các commit từ lúc tách nhánh tạo thành 1 commit merge duy nhất
            - ưu điểm: sẽ biết được dc change của branch nào
            - nhược điểm: nếu có nhiều merge sẽ tạo ra nhiều commit merge
        + rebase: lấy tất cả commit từ lúc tách nhánh rồi đem tất cả các commit đó 
            đặt lần lượt lên nhánh muốn merge đến theo thứ tự

6. Git tag
    - Quản lí phiên bản.
    - Ví dụ như sau khi làm 1 feature mới và được merge vào nhánh master ta sẽ tạo 1 tag mới

7. Rebase 2 commits in 1 commit
    - Dùng git rebase -i HEAD~N với n là số commit muốn gộp
    - sửa trong file theo hướng dẫn. 
        ví dụ: pick afdsbfgnfdbf
                f   dfgdbfrhnfhfhf
    - ESC => :x => commit => push