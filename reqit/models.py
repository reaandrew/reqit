class Result:

  def __init__(self, response):
    self.response = response

  def accept(self, visitor):
    visitor.visit(self.response)
