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
    => len(ExpSlicde) = 3, cap(ExpSlicde) = 5
    - duyệt slice: for i,v:= range slice_name;++{

    }
    - So sánh 2 slices: 2 slice có thể được so sánh khi chúng có len bằng nhau và kiểu giống nhau

3. Runes
    - Là 1 slice có kiểu là byte trả về mã ASC2 
    - Khai báo và duyệt làm tương tự như Slices
4. Append
    - z = []int{}
    - x:= []int{1,2,3,4,5} 
    - cơ chế hoạt động: 
        + nếu sức chứa của slices(cap) đủ lớn thì tạo 1 slice có low = 0, high = len(s) +1. 
        + nếu sức chứa của slices(cap) ko đủ lớn thì 
            ++ nới rộng lên và copy toàn bộ giá trị x vào z
            
    - gán lại giá trị tại index = len(x) bằng giá trị muốn append vào mảng.

5. Remove 
    - Cách 1: Copy các phần tử  từ (vị trí muốn xóa) + 1 đến cuối mảng vaò
                vị tri từ vị trí muốn xóa đến cuối mảng sau đó lấy slice có low = 0, high = len-1
    - Cách 2: Gán phân tử tại vị trí muốn xóa bằng phần tử cuối mảng, lấy 1 slide có low = 0, high = len -1

6. Maps 
    - Giống Dictionary trong C#
    - Mỗi 1 maps đều có key và value, mỗi 1 key là duy nhất 
    - Khai báo map_name := make(map[type]type)
    - duyệt các phần tử trong map:
        for key, value := range map_name
    - so sánh 2 map: 2 map được gọi là bằng nhau nếu có kích thước bằng nhau, value tại mỗi key tương ứng bằng nhau


7. Struct
    - Là 1 nhóm các biến có thể khác kiểu dữ liệu
    - khai báo: type name_struct struct{name_variable, type ...}
    - Cách truy xuất các variable trong struct: name_struct.name_variable
    - Trong 1 struct có thể gồm các struct khác

8. Json 
    - Là 1 cấu trúc để lưu trữ thông tin
    - Muốn lấy dữ liệu từ json trước hết tạo 1 struct có format giống với trong file .json
    - Đặt tag cho các varialbe: `json:"tag"`
    - Thường thì các Variable nên được Viết hoa chữ cái đấu còn tag thì viết thường 
    - Lấy dữ liệu từ json ta sẽ sử dụng json.unmarshal

9. Text và html template 
    - Được sử dụng để hiển thị dữ liệu lên template
    - 1 template cũng bao gồm các thẻ html cơ bản, giá trị truyền vào các thẻ được đặt trong 
        cặp dấu {{}} được gọi là các actions. Mỗi actions chứa 1 biểu thức để lấy giá trị tương ứng 
        với các varialbe trong structute
    - Ví dụ: <td <a href ='{{.Name_struct.varibale}}'>{{.Name_struct.varibale}}</a></td>