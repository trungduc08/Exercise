Báo cáo chương 4:
1. Array
    - Mảng là 1 tập hợp các phần tử có cùng kiểu
    - MẢng có kích thước cố định
    - Cách khai báo var array_name []element_type
    - Khi khai báo có thể chỉ ra kích thước, các phần tử của mảng hoặc là không
    - Truy cập 1 phần tử trong mảng: array_name[index]
    - So sánh 2 mảng: 2 mảng có thể so sánh được khi có cùng kích thước và kiểu

    - SHA256: là 1 hàm băm mã hóa 1 thông điệp thành 1 slice có kiểu là byte
    - Kết quả trả về là giống nhau nếu đầu vào là giống nhau

    - ArrayPointer: khi khai báo bắt buộc phải có kích thước của array

2. Slices
    - Slices cũng là 1 tập hợp các phần tử có cùng kiểu
    - Khai báo slice_name := make([]type, len, cap)
    - ExpArray :=[]int{1,2,3,4,5,6,7}
    - ExpSlicde := ExpArray[2:5]
    => len(ExpSlicde) = 3 cap(ExpSlicde) = 5