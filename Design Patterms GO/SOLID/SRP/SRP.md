Single Responsibility Principle - 
1. Seperation of Concerns. 
2. Every class should have only single responsibiity
3. An object which does everything is a GOD Object. 

Book struct is having method to add a book and display a book which are associated with it. 
Suppose if we also have to upload the book on the Web then we should not associate it with Book struct
instead create a new Struct called Upload and have that method associated with it. 
This is called as seperation of concern. 
