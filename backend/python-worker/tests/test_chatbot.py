import unittest
from app.chatbot import chatbot_response

class TestChatbot(unittest.TestCase):
    def test_response(self):
        self.assertEqual(chatbot_response("你好"), "好的")
        self.assertEqual(chatbot_response("Hello"), "好的")

if __name__ == "__main__":
    unittest.main()