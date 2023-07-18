from module1.class1 import Class1
from module2.class2 import Class2


if __name__ == '__main__':
    #Printing messages from dependencies
    Class1.use_dependency()
    Class2.use_dependency()