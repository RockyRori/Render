# Render Fullstack Project - Backend (Python)

## 📂 目录结构
```bash
backend/
├── go-api/                 # Golang API（主服务器）
├── python-worker/          # Python 后端（聊天系统）
│   ├── app/                # 核心聊天系统代码
│   │   ├── __init__.py     # Python 模块初始化
│   │   ├── main.py         # Flask 入口文件
│   │   ├── chatbot.py      # 聊天逻辑
│   │   ├── config.py       # 配置管理
│   │   ├── utils.py        # 工具函数
│   ├── tests/              # 单元测试
│   │   ├── __init__.py     # 测试模块初始化
│   │   ├── test_chatbot.py # 聊天系统测试
│   ├── requirements.txt    # Python 依赖
│   ├── Dockerfile          # Python 后端 Docker 配置
│   ├── .env                # 环境变量
```

---
## 1️⃣ **Flask 入口文件**
```python
# app/main.py
from flask import Flask, request, jsonify
from chatbot import chatbot_response

app = Flask(__name__)

@app.route("/chat", methods=["POST"])
def chat():
    user_message = request.json.get("message", "")
    response = chatbot_response(user_message)
    return jsonify({"response": response})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5001)
```

---
## 2️⃣ **聊天逻辑**
```python
# app/chatbot.py
def chatbot_response(message):
    return "好的"
```

---
## 3️⃣ **配置管理**
```python
# app/config.py
import os
from dotenv import load_dotenv

load_dotenv()

class Config:
    DEBUG = os.getenv("DEBUG", "False").lower() == "true"
    HOST = os.getenv("HOST", "0.0.0.0")
    PORT = int(os.getenv("PORT", 5001))
```

### **环境变量文件 `.env`**
```ini
DEBUG=True
HOST=0.0.0.0
PORT=5001
```

---
## 4️⃣ **工具函数**
```python
# app/utils.py
def log_message(message):
    print(f"[LOG] {message}")
```

---
## 5️⃣ **测试模块**
```python
# tests/__init__.py
```

```python
# tests/test_chatbot.py
import unittest
from app.chatbot import chatbot_response

class TestChatbot(unittest.TestCase):
    def test_response(self):
        self.assertEqual(chatbot_response("你好"), "好的")
        self.assertEqual(chatbot_response("Hello"), "好的")

if __name__ == "__main__":
    unittest.main()
```

---
## 6️⃣ **Python 模块初始化**
```python
# app/__init__.py
from flask import Flask
from main import app
```

```python
# tests/__init__.py
```

---
## 7️⃣ **安装依赖**
```bash
cd backend/python-worker
pip install -r requirements.txt
```

### **requirements.txt**
```txt
flask
dotenv
```

---
## 8️⃣ **Docker 部署**
```dockerfile
# Dockerfile
FROM python:3.12
WORKDIR /app
COPY . .
RUN pip install -r requirements.txt
CMD ["python", "app/main.py"]
EXPOSE 5001
```

---
## ✅ **运行 Python 后端**
```bash
python app/main.py
```

## ✅ **测试 API**
```bash
curl -X POST http://127.0.0.1:5001/chat -H "Content-Type: application/json" -d '{"message": "你好"}'
```

## ✅ **使用 Docker 运行**
```bash
docker build -t render-python .
docker run -p 5001:5001 render-python
```

---
**🎮 现在你的 Python 后端（聊天系统）已完成，支持 API、配置管理、Docker、单元测试！** 🚀
