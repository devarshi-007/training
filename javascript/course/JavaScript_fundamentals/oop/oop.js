import { print } from "../shared/common.js";
// classes and modules

class Vehicle {
  constructor(wheels, type) {
    this.wheels = wheels;
    this.type = type;
  }
}

class Car extends Vehicle {
  constructor(id) {
    super(4, "heavy");
    this.id = id;
  }

  identify() {
    return `Car id: ${this.id}`;
  }
}

export default function main() {
  let car = new Car("abc");
  car.id = "def";
  print(car.id);
  print(car.identify());
  print(car.type);
  print(car);
}
