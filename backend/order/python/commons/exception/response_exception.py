class ResponseException(Exception):
    def __init__(self, status, error_messages, message):
        self.status = status
        self.error_messages = error_messages
        self.message = message
        super().__init__(self.message)
    
    def __str__(self):
        return f"[{self.status}] {self.error_messages}: {self.message}"