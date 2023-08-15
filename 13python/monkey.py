class MyClass:
    def greet(self):
        return "Hello, world!"

# Monkey patching the greet method of MyClass
def new_greet(self):
    return "Hola, mundo!"


# Monkey patching the greet method of MyClass
def new_greet2(self):
    return "Madam, Hola"


MyClass.greet = new_greet2

# Creating an instance of MyClass and calling the modified greet method
obj = MyClass()
print(obj.greet())  # Output: Hola, mundo!
