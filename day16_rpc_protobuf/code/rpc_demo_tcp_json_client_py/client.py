

# 与127.0.0.1:9093 TCP连接
# 准备参数和返回值
# 调用


import socket
import json

request = {
    # "id": 0,
    "params": [{"x":100, "y":20}],  # 参数要对应上Args结构体
    "method": "ServiceA.Add"
}

# 新建一个tcp的client对象 重试5次
client = socket.create_connection(("127.0.0.1", 9093), 5)
# 发送数据 JSON格式的数据
client.sendall(json.dumps(request).encode())
# 接收返回的结果
rsp = client.recv(1024)
# JSON反序列化
rsp = json.loads(rsp.decode())
print(rsp)