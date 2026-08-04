## Some popular concepts that go do not support

Go also has a few other popular concepts not available.

In Go, we don't have some of these nice things. So we embrace the [Grug](https://grugbrain.dev/) mentality.

### Enums -> <https://en.wikipedia.org/wiki/Enumerated_type>

  An enumerated type (also called enumeration, enum, or factor in the R programming language, a condition-name in the COBOL programming language,
  a status variable in the JOVIAL programming language, an ordinal in the PL/I programming language, and a categorical variable in statistics)
  is a data type[a] consisting of a set of named values called elements, members, enumeral, or enumerators of the type.
  The enumerator names are usually identifiers that behave as constants in the language.
  An enumerated type can be seen as a degenerate tagged union of unit type.

  Go lacks enums. Example, in typescript, you can write

  ```ts
    // Example of an enum in TypeScript: defining and using an enum for directions

    enum Direction {
      North,
      South,
      East,
      West
    }

    // Usage:
    let travelDirection: Direction = Direction.North;

    if (travelDirection === Direction.North) {
      console.log("Heading north!");
    }
  ```

### Algebraic data types

  An algebraic data type is defined by two key constructions: a sum and a product. These are sometimes referred to as "OR" and "AND" types.

- #### Sum Types -> <https://en.wikipedia.org/wiki/Algebraic_data_type>

  A sum type is a choice between possibilities. The value of a sum type can match one of several defined variants.
  For example, a type representing the state of a traffic light could be either Red, Amber, or Green

  Go lacks sum types. Example, in typescript, you can write

  ```ts
    // sum types are basically union types in typescript.
    // To get a true sum type, you must introduce a tag. see ### Tagged Unions for example, that's a sum type
    type TrafficLightState = "Red" | "Amber" | "Green"
  ```

- #### Product Types -> <https://en.wikipedia.org/wiki/Algebraic_data_type>

  A product type combines types together. A value of a product type will contain a value for each of its component types.
  For example, a Point type might be defined to contain an x coordinate (an integer) and a y coordinate (also an integer)

    Go lacks product types. Example, in typescript, you can write

  ```ts
    // Example of a product type in TypeScript: defining a Point with x and y coordinates
    type Point = { // An object with multiple properties is the most common product type
      x: number;
      y: number;
    };

    const p: Point = { x: 10, y: 20 };

    // Tuples:
    // Tuples are arrays with a fixed length and specific types at each index
    type Point = [number, string]; // Must contain a number AND a string in that exact order

    // Intersection Types (&)
    // combining multiple types into one using the & operator
    type Character = { name: string };

    type Combatant = { hp: number };

    type Hero = Character & Combatant;
    // A Hero must have a name AND an hp property
    const player: Hero = { name: "Link", hp: 100 };
  ```

### Tagged Unions -> <https://en.wikipedia.org/wiki/Tagged_union>

  Also called a discriminated union, variant, variant record, choice type, disjoint union, sum type, or coproduct,
  is a data structure used to hold a value that could take on several different, but fixed, types.
  Only one of the types can be in use at any one time, and a tag field explicitly indicates which type is in use
  
  Go lacks enums. Example, in typescript, you can write

  ```ts
    // Example of a tagged union (discriminated union) in TypeScript:
    type Shape =
      | { kind: "circle"; radius: number }
      | { kind: "rectangle"; width: number; height: number };

    function area(shape: Shape): number {
      switch (shape.kind) {
        case "circle":
          return Math.PI * shape.radius * shape.radius;
        case "rectangle":
          return shape.width * shape.height;
      }
    }

    const c: Shape = { kind: "circle", radius: 3 };
    const r: Shape = { kind: "rectangle", width: 4, height: 5 };
    
    console.log(area(c)); // 28.27...
    console.log(area(r)); // 20
  ```
