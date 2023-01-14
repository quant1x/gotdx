# cstruct-go

a fast c-style struct packer & unpacker for golang


## 特性

- Go struct 与 C struct 一一对应，内存分布一致。
- 快速。
- 方便，只需要定义 struct 即可。


## 例子1

  1. 定义举例

      ```go
      type mystruct1 struct {
          F1  bool
          F2  float32
          F3  float64
          F4  string
          F5  []byte
          F6  int8
          F7  int16
          F9  int32
          F11 int64
          F12 uint8
          F13 uint16
          F15 uint32
          F17 uint64
          F18 [20]byte
          F19 [16]uint32
      }
      ```

  2. 使用举例

      ```go
      a := &mystruct1{}
      // ...(初始化a 代码略)...

      // 序列化代码如下，返回 []byte类型数据
      buf_l, _ := cstruct.Marshal(a)

      // 反序列化代码如下
      b := &mystruct1{}
      if err := cstruct.Unmarshal(buf_l, b); err != nil {
          fmt.Println(err)
          return
      }
      ```

  详细例子可以参考：

  - [x_test.go](tests/x_test.go)
  - [array_test.go](tests/array_test.go)


## 例子2

演示 C struct 与 Go struct 内存分布一致

  - C struct：

    ```c++
    #include <string>

    #pragma pack(1)

    struct StructA {
      uint8_t A1;
      uint32_t A2;
      uint8_t A3[5];
    };


    struct StructB {
      uint8_t B1;
      StructA B2;
      uint16_t B3;
      float B4;
      StructA B5[3];
    };

    int main()
    {
      StructB b;
      b.B1 = 127;
      b.B2.A1 = 56;
      b.B2.A2 = 999;
      b.B2.A3[0] = 0;
      b.B2.A3[1] = 1;
      b.B2.A3[2] = 2;
      b.B2.A3[3] = 3;
      b.B2.A3[4] = 4;
      b.B3 = 8888;
      b.B4 = 88.8f;
      b.B5[0] = b.B2;
      b.B5[1] = b.B2;
      b.B5[2] = b.B2;

      printf("len(b) = %llu\n", sizeof(b));
      printf("struct data len = %llu\n", sizeof(b));
      printf("struct data is:\n");

      unsigned char buff[1024];
      memcpy(buff, &b, sizeof(b));
      for (int i = 0; i < sizeof(b); i++) {
        printf("%d ", buff[i]);
      }
      return 0;
    }
    ```

  - Go struct

    ```golang
    type StructA struct {
      A1 uint8
      A2 uint32
      A3 [5]uint8
    }

    type StructB struct {
      B1 uint8
      B2 StructA
      B3 uint16
      B4 float32
      B5 [3]StructA
    }

    func main() {
      b := StructB{}
      b.B1 = 127
      b.B2.A1 = 56
      b.B2.A2 = 999
      b.B2.A3[0] = 0
      b.B2.A3[1] = 1
      b.B2.A3[2] = 2
      b.B2.A3[3] = 3
      b.B2.A3[4] = 4
      b.B3 = 8888
      b.B4 = 88.8
      b.B5[0] = b.B2
      b.B5[1] = b.B2
      b.B5[2] = b.B2

      data, _ := cstruct.Marshal(&b)

      fmt.Println("len(b) =", unsafe.Sizeof(b))
      fmt.Println("cstruct data len = ", len(data))
      fmt.Println("cstruct data is:")
      for i := 0; i < len(data); i++ {
        fmt.Printf("%d ", data[i])
      }
    }
    ```

  - 以上控制台输出

    ```shell
    D:\golang\src\github.com\fananchong\ccstruct-go\example>main.exe
    len(b) = 76
    cstruct data len =  47
    cstruct data is:
    127 56 231 3 0 0 0 1 2 3 4 184 34 154 153 177 66 56 231 3 0 0 0 1 2 3 4 56 231 3 0 0 0 1 2 3 4 56 231 3 0 0 0 1 2 3 4
    D:\golang\src\github.com\fananchong\ccstruct-go\example>main_cpp.exe
    len(b) = 47
    cstruct data len = 47
    cstruct data is:
    127 56 231 3 0 0 0 1 2 3 4 184 34 154 153 177 66 56 231 3 0 0 0 1 2 3 4 56 231 3 0 0 0 1 2 3 4 56 231 3 0 0 0 1 2 3 4
    ```

  - 详细例子可以参考：
    - [example](example)


## 字节序

`小端`字节序。

## 基本类型

go类型    | 内存说明
------- | ----------------------------
bool    | 1 byte
int8    | 1 byte
uint8   | 1 byte
int16   | 2 byte
uint16  | 2 byte
int32   | 4 byte
uint32  | 4 byte
int64   | 8 byte
uint64  | 8 byte
float32 | 4 byte
float64 | 8 byte
string  | [2 byte] + [len(字符串) byte]
[]byte  | [2 byte] + [len(2进制数据) byte]


## 数组类型

- 支持基本类型的数组（string 、 []byte 除外，因为它们不定长）
- 支持定长struct的数组


## 指针类型

- 支持struct指针（默认格式）

go类型       | 内存说明                                |byte含义说明
------------ | ----------------------------------------|--------------
struct ptr   | [1 byte] + [len(struct内存占用) byte]   | 1byte值为0，表示指针为nil


- 支持struct指针（当 cstruct.OptionSliceIgnoreNil = true ）

go类型       | 内存说明                                
------------ | ----------------------------------------
struct ptr   | [len(struct内存占用) byte]              


## 复杂类型

- 支持struct嵌套
- 支持基本类型的Slice
- 支持struct指针的Slice
- 支持struct的Slice

go类型  | 内存说明
----- | -----------------------------
slice | [2 byte] + [len(元素内存占用) byte]

## 基准测试

![图](assets/1.jpg)

测试环境：`阿里云ECS服务器，配置： 1 vCPU 2 GB (I/O优化) ecs.t5-lc1m2.small 1Mbps`

基准测试代码：[cstrucgo_test.go](benchmarks/cstrucgo_test.go)


## 参考项目

- <https://github.com/golang/protobuf>
- <https://github.com/gogo/protobuf>
