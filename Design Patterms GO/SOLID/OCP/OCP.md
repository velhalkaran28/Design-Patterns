Open Closed Principle. 
1. Open for extension but closed for modification. 
2. This is well understood by using Specification Enterprise pattern. 

Consider we are operating a some online store selling widgets of some kind and we want the end user of website
to filter the items based on some properties like size or color. 

First we implemented a filter to filter products by Color. 
But what if now Manager asks to filter products by Size? - We will have to add yet another method for it. 
But what if manager comes back and asks to add new filter method - Filter by Size and Color - Again we shall
have to implement it (write a code for it)
We are always modifying the code for it which a violation of OCP. 

To solve this problem we use Specification Principle which is described in the code. 
We created a Specification interface which has a method IsSatisfied. 
The required specifications Color and Size are created as a structs which implements the interface Specification. 
Then we created a SpecFilter function which applies any Specification on products. 
With this, we achieved OCP Principle. Suppose any new specification is added, we can make a struct of it
Specification interface is Open for Extension but closed for Modfication.
and implement the method IsSatisfied, that's all!