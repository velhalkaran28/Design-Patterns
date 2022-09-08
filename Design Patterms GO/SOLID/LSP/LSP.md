Liskov Substitution Principle. 
Functions that use pointers or references to base classes muse be able to use objects of derived classes without knowing it.


In the example we coded:
person is a interface having getName as a function. Human is implementing it. Also, we have 2 structs teacher and student which composes human. 

The function Info accepts interface person and prints its name. Now since human implements person and student & teacher composes human, the method function accepts all types. 

