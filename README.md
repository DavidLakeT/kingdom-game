### Kingdom Game

## Overview


- **What is this project?** This is an implementation of the "Abstract Factory" creational design pattern in Go.
- **What is Abstract Factory?** Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.
- **What is Go?** Go is a statically typed, compiled high-level programming language designed at Google.

## Applicability

 The Abstract Factory pattern provides an interface to create objects from each class of the product family.
 Utiliza el patrón Abstract Factory cuando tu código deba funcionar con varias familias de productos relacionados.

* The system should be configured with one of the multiple families of elements.
* * When you need consistency among elements.
* The family of related element objects is designed to be used together, and you need to enforce this constraint.
* You want to decide which product to call from a family at runtime.
* When you have a class with a group of factory methods that cloud its main responsibility.

## This Implementation

This example makes use of the "Abstract Factory" design pattern by simulating kingdoms of different types (in this case, elves and orcs), which have their own castles, armies and kings. The qualities of these elements are shared across all instances created (although the result of the functions may change).

This is achieved through the implementation of factories that, depending on the realm to which they correspond, will create different objects (Example: The factory of the elves' realm can create the castle, the army and the king of the elves).

Likewise, the customer can make use of "MakeFactory" within "KingdomType" to stipulate the kingdom to which the factory to be created will belong.